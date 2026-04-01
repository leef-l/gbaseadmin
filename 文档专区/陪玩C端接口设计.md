# 陪玩平台 C端接口设计文档

> 路由前缀：`/api/playapi`
> 应用位置：`app/play/` 内，与后台管理共用同一应用
> C端接口文件命名：`*_api.go`（区别于 codegen 生成的后台接口）

---

## 一、C端会员鉴权中间件设计

### 1.1 MemberAuth 中间件

文件位置：`app/play/internal/middleware/member_auth.go`

```go
package middleware

import (
    "strings"

    "github.com/gogf/gf/v2/net/ghttp"

    "gbaseadmin/utility/jwt"
    "gbaseadmin/utility/response"
)

// MemberAuth C端会员 JWT 鉴权中间件
// 区别于后台 Auth 中间件：
// 1. Claims 结构不同，包含 MemberID/IsCoach/CoachID/CurrentRole
// 2. 不校验后台权限（perms），仅校验会员状态
// 3. 支持双身份（member/coach）切换
func MemberAuth(r *ghttp.Request) {
    tokenStr := r.GetHeader("Authorization")
    if tokenStr == "" {
        response.Unauthorized(r)
        return
    }

    tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
    tokenStr = strings.TrimSpace(tokenStr)
    if tokenStr == "" {
        response.Unauthorized(r)
        return
    }

    claims, err := jwt.ParseMemberToken(tokenStr)
    if err != nil {
        response.Unauthorized(r, "Token无效或已过期")
        return
    }

    // 将会员信息写入 context
    r.SetCtxVar("jwt_member_id", claims.MemberID)
    r.SetCtxVar("jwt_is_coach", claims.IsCoach)
    r.SetCtxVar("jwt_coach_id", claims.CoachID)
    r.SetCtxVar("jwt_current_role", claims.CurrentRole)
    r.SetCtxVar("jwt_member_claims", claims)

    r.Middleware.Next()
}

// CoachOnly 陪玩师专属接口中间件（需配合 MemberAuth 使用）
// 校验当前身份必须是 coach 且 is_coach=1
func CoachOnly(r *ghttp.Request) {
    currentRole := r.GetCtxVar("jwt_current_role").String()
    isCoach := r.GetCtxVar("jwt_is_coach").Int()

    if currentRole != "coach" || isCoach != 1 {
        response.Forbidden(r, "需要切换到陪玩师身份")
        return
    }

    r.Middleware.Next()
}
```

### 1.2 MemberClaims 结构体

文件位置：`utility/jwt/jwt.go`（扩展）

```go
// MemberClaims C端会员 JWT 载荷
type MemberClaims struct {
    MemberID    int64  `json:"memberId"`
    Phone       string `json:"phone"`
    IsCoach     int    `json:"isCoach"`
    CoachID     int64  `json:"coachId"`
    CurrentRole string `json:"currentRole"` // member / coach
    gojwt.RegisteredClaims
}
```

---

## 二、认证接口（无需登录）

文件位置：`app/play/api/auth_api.go`

### 2.1 手机号+验证码登录

- **路径**：`POST /auth/login`
- **登录**：否
- **CoachOnly**：否
- **说明**：手机号+验证码登录，不存在自动注册

**请求参数**：

```go
type AuthLoginReq struct {
    g.Meta `path:"/auth/login" method:"post" tags:"C端认证" summary:"手机号验证码登录"`
    Phone  string `json:"phone" v:"required|phone#手机号不能为空|手机号格式不正确" dc:"手机号"`
    Code   string `json:"code" v:"required|length:4,6#验证码不能为空|验证码长度4-6位" dc:"短信验证码"`
}
```

**响应参数**：

```go
type AuthLoginRes struct {
    g.Meta       `mime:"application/json"`
    Token        string `json:"token" dc:"访问令牌"`
    RefreshToken string `json:"refreshToken" dc:"刷新令牌"`
    ExpiresIn    int64  `json:"expiresIn" dc:"过期时间(秒)"`
    IsNew        bool   `json:"isNew" dc:"是否新注册用户"`
}
```

**业务逻辑**：

```
1. 校验验证码
   - 从 Redis 获取 sms:login:{phone} 对应的验证码
   - 比对失败返回"验证码错误或已过期"
   - 比对成功后删除 Redis key
2. 查询会员是否存在
   - SELECT * FROM play_member WHERE phone=? AND deleted_at IS NULL
3. 若不存在，自动注册
   - 生成 Snowflake ID
   - 默认昵称="用户"+手机后4位，默认头像
   - 查询最低等级 member_level_id
   - INSERT INTO play_member
4. 若存在，检查 status
   - status=0 返回"账号已被禁用"
5. 签发 JWT（MemberClaims）
   - MemberID, Phone, IsCoach, CoachID, CurrentRole="member"
6. 生成 RefreshToken 并存入 Redis（refresh_token:{memberId}，TTL 30天）
7. 返回 Token + RefreshToken + ExpiresIn + IsNew
```

---

### 2.2 发送验证码

- **路径**：`POST /auth/send_code`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type AuthSendCodeReq struct {
    g.Meta `path:"/auth/send_code" method:"post" tags:"C端认证" summary:"发送验证码"`
    Phone  string `json:"phone" v:"required|phone#手机号不能为空|手机号格式不正确" dc:"手机号"`
    Scene  string `json:"scene" v:"required|in:login,bindPhone#场景不能为空|场景值不合法" dc:"场景:login=登录,bindPhone=绑定手机"`
}
```

**响应参数**：

```go
type AuthSendCodeRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 频率限制
   - Redis key: sms:limit:{phone}，1分钟内不可重复发送
   - 若存在返回"发送过于频繁，请稍后再试"
2. 生成6位随机验证码
3. 调用短信服务发送验证码
4. 存入 Redis
   - sms:{scene}:{phone} = code，TTL 5分钟
   - sms:limit:{phone} = 1，TTL 60秒
5. 返回成功
```

---

### 2.3 刷新Token

- **路径**：`POST /auth/refresh_token`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type AuthRefreshTokenReq struct {
    g.Meta       `path:"/auth/refresh_token" method:"post" tags:"C端认证" summary:"刷新Token"`
    RefreshToken string `json:"refreshToken" v:"required#刷新令牌不能为空" dc:"刷新令牌"`
}
```

**响应参数**：

```go
type AuthRefreshTokenRes struct {
    g.Meta       `mime:"application/json"`
    Token        string `json:"token" dc:"新访问令牌"`
    RefreshToken string `json:"refreshToken" dc:"新刷新令牌"`
    ExpiresIn    int64  `json:"expiresIn" dc:"过期时间(秒)"`
}
```

**业务逻辑**：

```
1. 解析 RefreshToken，提取 memberId
2. 从 Redis 获取 refresh_token:{memberId}，比对是否一致
   - 不一致返回"刷新令牌无效"
3. 查询会员信息（含 coach 信息）
   - status=0 返回"账号已被禁用"
4. 重新签发 JWT（MemberClaims）
5. 生成新 RefreshToken，更新 Redis（TTL 30天）
6. 返回新 Token + RefreshToken + ExpiresIn
```

---

### 2.4 微信登录

- **路径**：`POST /auth/wx_login`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type AuthWxLoginReq struct {
    g.Meta `path:"/auth/wx_login" method:"post" tags:"C端认证" summary:"微信登录"`
    Code   string `json:"code" v:"required#微信授权码不能为空" dc:"微信授权code"`
}
```

**响应参数**：

```go
type AuthWxLoginRes struct {
    g.Meta       `mime:"application/json"`
    Token        string `json:"token" dc:"访问令牌"`
    RefreshToken string `json:"refreshToken" dc:"刷新令牌"`
    ExpiresIn    int64  `json:"expiresIn" dc:"过期时间(秒)"`
    IsNew        bool   `json:"isNew" dc:"是否新注册用户"`
    NeedBindPhone bool  `json:"needBindPhone" dc:"是否需要绑定手机号"`
}
```

**业务逻辑**：

```
1. 用 code 调用微信 OAuth2 接口获取 openid + unionid + access_token
2. 查询 play_member WHERE wx_openid=?
3. 若存在
   - status=0 返回"账号已被禁用"
   - 签发 JWT，返回 Token
4. 若不存在
   - 调用微信用户信息接口获取昵称、头像
   - 生成 Snowflake ID
   - INSERT INTO play_member（phone 为空，wx_openid=openid）
   - needBindPhone=true
5. 签发 JWT + RefreshToken
6. 返回 Token + IsNew + NeedBindPhone
```

---

### 2.5 支付宝登录

- **路径**：`POST /auth/alipay_login`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type AuthAlipayLoginReq struct {
    g.Meta   `path:"/auth/alipay_login" method:"post" tags:"C端认证" summary:"支付宝登录"`
    AuthCode string `json:"authCode" v:"required#支付宝授权码不能为空" dc:"支付宝授权code"`
}
```

**响应参数**：

```go
type AuthAlipayLoginRes struct {
    g.Meta        `mime:"application/json"`
    Token         string `json:"token" dc:"访问令牌"`
    RefreshToken  string `json:"refreshToken" dc:"刷新令牌"`
    ExpiresIn     int64  `json:"expiresIn" dc:"过期时间(秒)"`
    IsNew         bool   `json:"isNew" dc:"是否新注册用户"`
    NeedBindPhone bool   `json:"needBindPhone" dc:"是否需要绑定手机号"`
}
```

**业务逻辑**：

```
1. 用 authCode 调用支付宝 OAuth 接口获取 user_id
2. 查询 play_member WHERE alipay_user_id=?
3. 若存在
   - status=0 返回"账号已被禁用"
   - 签发 JWT，返回 Token
4. 若不存在
   - 调用支付宝用户信息接口获取昵称、头像
   - 生成 Snowflake ID
   - INSERT INTO play_member（phone 为空，alipay_user_id=user_id）
   - needBindPhone=true
5. 签发 JWT + RefreshToken
6. 返回 Token + IsNew + NeedBindPhone
```

---

## 三、会员接口（需 MemberAuth）

文件位置：`app/play/api/member_api.go`

### 3.1 获取个人信息

- **路径**：`GET /member/info`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type MemberInfoReq struct {
    g.Meta `path:"/member/info" method:"get" tags:"C端会员" summary:"获取个人信息"`
}
```

**响应参数**：

```go
type MemberInfoRes struct {
    g.Meta        `mime:"application/json"`
    MemberID      string `json:"memberId" dc:"会员ID"`
    Phone         string `json:"phone" dc:"手机号（脱敏）"`
    Nickname      string `json:"nickname" dc:"昵称"`
    Avatar        string `json:"avatar" dc:"头像URL"`
    Gender        int    `json:"gender" dc:"性别:0=未知,1=男,2=女"`
    Balance       int64  `json:"balance" dc:"余额(分)"`
    LevelTitle    string `json:"levelTitle" dc:"等级名称"`
    LevelIcon     string `json:"levelIcon" dc:"等级图标"`
    Discount      int    `json:"discount" dc:"会员折扣"`
    Exp           int    `json:"exp" dc:"当前经验值"`
    IsCoach       int    `json:"isCoach" dc:"是否陪玩师:0=否,1=是"`
    CoachID       string `json:"coachId" dc:"陪玩师ID"`
    CurrentRole   string `json:"currentRole" dc:"当前身份:member/coach"`
    WxBound       bool   `json:"wxBound" dc:"是否绑定微信"`
    AlipayBound   bool   `json:"alipayBound" dc:"是否绑定支付宝"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_member + LEFT JOIN play_member_level
3. 手机号脱敏处理（保留前3后4）
4. 组装返回数据
```

---

### 3.2 编辑资料

- **路径**：`PUT /member/update`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type MemberUpdateReq struct {
    g.Meta   `path:"/member/update" method:"put" tags:"C端会员" summary:"编辑资料"`
    Nickname string `json:"nickname" v:"max-length:20#昵称最多20个字符" dc:"昵称"`
    Avatar   string `json:"avatar" dc:"头像URL"`
    Gender   *int   `json:"gender" v:"in:0,1,2#性别值不合法" dc:"性别:0=未知,1=男,2=女"`
}
```

**响应参数**：

```go
type MemberUpdateRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 构建更新 map（仅更新非空字段）
3. UPDATE play_member SET ... WHERE id=?
4. 若是陪玩师，同步更新 play_coach 的 nickname/avatar
5. 返回成功
```

---

### 3.3 切换身份

- **路径**：`POST /member/switch_role`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type MemberSwitchRoleReq struct {
    g.Meta `path:"/member/switch_role" method:"post" tags:"C端会员" summary:"切换身份"`
    Role   string `json:"role" v:"required|in:member,coach#身份不能为空|身份值必须为member或coach" dc:"目标身份:member/coach"`
}
```

**响应参数**：

```go
type MemberSwitchRoleRes struct {
    g.Meta `mime:"application/json"`
    Token  string `json:"token" dc:"新访问令牌"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_claims
2. 若切换到 coach
   - 校验 is_coach=1，否则返回"您还不是陪玩师"
   - 查询 play_coach 获取 coach_id
3. 重新签发 JWT，CurrentRole=目标身份
4. 返回新 Token
```

---

### 3.4 余额流水列表

- **路径**：`GET /member/balance_log`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type MemberBalanceLogReq struct {
    g.Meta   `path:"/member/balance_log" method:"get" tags:"C端会员" summary:"余额流水列表"`
    Type     string `json:"type" dc:"类型筛选:income=收入,expense=支出，空=全部"`
    Page     int    `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type MemberBalanceLogRes struct {
    g.Meta `mime:"application/json"`
    Total  int                   `json:"total" dc:"总数"`
    List   []MemberBalanceLogItem `json:"list" dc:"流水列表"`
}

type MemberBalanceLogItem struct {
    ID        string `json:"id" dc:"流水ID"`
    Type      string `json:"type" dc:"类型:recharge=充值,pay=支付,refund=退款,income=收入,withdraw=提现"`
    Amount    int64  `json:"amount" dc:"金额(分)，正数=收入，负数=支出"`
    Balance   int64  `json:"balance" dc:"变动后余额(分)"`
    Title     string `json:"title" dc:"标题"`
    Remark    string `json:"remark" dc:"备注"`
    CreatedAt string `json:"createdAt" dc:"创建时间"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_balance_log WHERE member_id=?
   - 若 type=income，筛选 amount>0
   - 若 type=expense，筛选 amount<0
3. ORDER BY created_at DESC，分页返回
```

---

## 四、陪玩师接口

文件位置：`app/play/api/coach_api.go`

### 4.1 陪玩师列表

- **路径**：`GET /coach/list`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type CoachListReq struct {
    g.Meta     `path:"/coach/list" method:"get" tags:"C端陪玩师" summary:"陪玩师列表"`
    CategoryID string `json:"categoryId" dc:"分类ID筛选"`
    Keyword    string `json:"keyword" dc:"关键词搜索（昵称/技能描述）"`
    Gender     *int   `json:"gender" dc:"性别筛选:1=男,2=女"`
    OnlineOnly *int   `json:"onlineOnly" dc:"仅看在线:1=是"`
    SortBy     string `json:"sortBy" v:"in:score,orderCount,price_asc,price_desc#排序值不合法" dc:"排序:score=评分,orderCount=接单量,price_asc=价格升序,price_desc=价格降序"`
    Page       int    `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize   int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type CoachListRes struct {
    g.Meta `mime:"application/json"`
    Total  int             `json:"total" dc:"总数"`
    List   []CoachListItem `json:"list" dc:"陪玩师列表"`
}

type CoachListItem struct {
    CoachID      string   `json:"coachId" dc:"陪玩师ID"`
    Nickname     string   `json:"nickname" dc:"昵称"`
    Avatar       string   `json:"avatar" dc:"头像"`
    Gender       int      `json:"gender" dc:"性别"`
    IsOnline     int      `json:"isOnline" dc:"在线状态:0=离线,1=在线"`
    Score        float64  `json:"score" dc:"综合评分"`
    OrderCount   int      `json:"orderCount" dc:"接单量"`
    SkillDesc    string   `json:"skillDesc" dc:"技能描述"`
    MinPrice     int64    `json:"minPrice" dc:"最低商品价格(分)"`
    Categories   []string `json:"categories" dc:"擅长分类名称列表"`
}
```

**业务逻辑**：

```
1. 查询 play_coach WHERE status=1（已审核通过）
2. 若 categoryId 不为空
   - JOIN play_goods WHERE category_id=? 筛选有该分类商品的陪玩师
3. 若 keyword 不为空
   - WHERE nickname LIKE '%keyword%' OR skill_desc LIKE '%keyword%'
4. 若 gender 不为空，WHERE gender=?
5. 若 onlineOnly=1，WHERE is_online=1
6. 排序处理
   - score: ORDER BY score DESC
   - orderCount: ORDER BY order_count DESC
   - price_asc: ORDER BY min_price ASC
   - price_desc: ORDER BY min_price DESC
   - 默认: ORDER BY is_online DESC, score DESC
7. 分页返回
```

---

### 4.2 陪玩师详情

- **路径**：`GET /coach/detail`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type CoachDetailReq struct {
    g.Meta  `path:"/coach/detail" method:"get" tags:"C端陪玩师" summary:"陪玩师详情"`
    CoachID string `json:"coachId" v:"required#陪玩师ID不能为空" dc:"陪玩师ID"`
}
```

**响应参数**：

```go
type CoachDetailRes struct {
    g.Meta       `mime:"application/json"`
    CoachID      string            `json:"coachId" dc:"陪玩师ID"`
    Nickname     string            `json:"nickname" dc:"昵称"`
    Avatar       string            `json:"avatar" dc:"头像"`
    Gender       int               `json:"gender" dc:"性别"`
    IsOnline     int               `json:"isOnline" dc:"在线状态"`
    Score        float64           `json:"score" dc:"综合评分"`
    OrderCount   int               `json:"orderCount" dc:"接单量"`
    SkillDesc    string            `json:"skillDesc" dc:"技能描述"`
    GoodsList    []CoachGoodsItem  `json:"goodsList" dc:"商品列表"`
    ReviewCount  int               `json:"reviewCount" dc:"评价总数"`
    RecentReviews []ReviewBriefItem `json:"recentReviews" dc:"最近3条评价"`
}

type CoachGoodsItem struct {
    GoodsID      string `json:"goodsId" dc:"商品ID"`
    Title        string `json:"title" dc:"商品标题"`
    CategoryName string `json:"categoryName" dc:"分类名称"`
    Price        int64  `json:"price" dc:"单价(分)"`
    Unit         string `json:"unit" dc:"单位(局/小时/次)"`
    Description  string `json:"description" dc:"商品描述"`
    Status       int    `json:"status" dc:"状态:0=下架,1=上架"`
}

type ReviewBriefItem struct {
    Nickname    string  `json:"nickname" dc:"评价者昵称"`
    Avatar      string  `json:"avatar" dc:"评价者头像"`
    Score       float64 `json:"score" dc:"评分"`
    Content     string  `json:"content" dc:"评价内容"`
    IsAnonymous int     `json:"isAnonymous" dc:"是否匿名"`
    CreatedAt   string  `json:"createdAt" dc:"评价时间"`
}
```

**业务逻辑**：

```
1. 查询 play_coach WHERE id=? AND status=1
   - 不存在返回"陪玩师不存在"
2. 查询该陪玩师的商品列表
   - SELECT * FROM play_goods WHERE coach_id=? AND status=1 ORDER BY sort ASC
3. 查询评价统计
   - SELECT COUNT(*), AVG(score) FROM play_review WHERE coach_id=?
4. 查询最近3条评价
   - SELECT * FROM play_review WHERE coach_id=? ORDER BY created_at DESC LIMIT 3
   - 匿名评价隐藏昵称和头像
5. 组装返回数据
```

---

### 4.3 申请成为陪玩师

- **路径**：`POST /coach/apply`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type CoachApplyReq struct {
    g.Meta    `path:"/coach/apply" method:"post" tags:"C端陪玩师" summary:"申请成为陪玩师"`
    RealName  string `json:"realName" v:"required|length:2,20#真实姓名不能为空|姓名长度2-20个字符" dc:"真实姓名"`
    IDCard    string `json:"idCard" v:"required|resident-id#身份证号不能为空|身份证号格式不正确" dc:"身份证号"`
    IDCardFront string `json:"idCardFront" v:"required#证件照正面不能为空" dc:"证件照正面URL"`
    IDCardBack  string `json:"idCardBack" v:"required#证件照反面不能为空" dc:"证件照反面URL"`
    SkillDesc string `json:"skillDesc" v:"required|max-length:500#技能描述不能为空|技能描述最多500字" dc:"技能描述"`
}
```

**响应参数**：

```go
type CoachApplyRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 检查是否已是陪玩师
   - SELECT * FROM play_coach WHERE member_id=? AND deleted_at IS NULL
   - 若 status=1 返回"您已经是陪玩师"
   - 若 status=0（待审核）返回"您的申请正在审核中"
3. 生成 Snowflake ID
4. INSERT INTO play_coach
   - member_id, real_name, id_card, id_card_front, id_card_back, skill_desc
   - status=0（待审核）
   - 从 play_member 同步 nickname, avatar, gender
5. 返回成功
```

---

### 4.4 查询申请状态

- **路径**：`GET /coach/apply_status`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type CoachApplyStatusReq struct {
    g.Meta `path:"/coach/apply_status" method:"get" tags:"C端陪玩师" summary:"查询申请状态"`
}
```

**响应参数**：

```go
type CoachApplyStatusRes struct {
    g.Meta       `mime:"application/json"`
    HasApplied   bool   `json:"hasApplied" dc:"是否已申请"`
    Status       int    `json:"status" dc:"状态:0=待审核,1=已通过,2=已拒绝"`
    RejectReason string `json:"rejectReason" dc:"拒绝原因"`
    AppliedAt    string `json:"appliedAt" dc:"申请时间"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. SELECT * FROM play_coach WHERE member_id=? AND deleted_at IS NULL ORDER BY created_at DESC LIMIT 1
3. 若无记录，hasApplied=false
4. 若有记录，返回 status/reject_reason/created_at
```

---

### 4.5 设置在线状态

- **路径**：`PUT /coach/online`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type CoachOnlineReq struct {
    g.Meta   `path:"/coach/online" method:"put" tags:"C端陪玩师" summary:"设置在线状态"`
    IsOnline int `json:"isOnline" v:"required|in:0,1#在线状态不能为空|在线状态值不合法" dc:"在线状态:0=离线,1=在线"`
}
```

**响应参数**：

```go
type CoachOnlineRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. UPDATE play_coach SET is_online=? WHERE id=?
3. 返回成功
```

---

### 4.6 我的商品列表

- **路径**：`GET /coach/my_goods`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type CoachMyGoodsReq struct {
    g.Meta   `path:"/coach/my_goods" method:"get" tags:"C端陪玩师" summary:"我的商品列表"`
    Status   *int `json:"status" dc:"状态筛选:0=下架,1=上架"`
    Page     int  `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int  `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type CoachMyGoodsRes struct {
    g.Meta `mime:"application/json"`
    Total  int              `json:"total" dc:"总数"`
    List   []CoachGoodsItem `json:"list" dc:"商品列表"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. SELECT * FROM play_goods WHERE coach_id=?
   - 若 status 不为空，WHERE status=?
3. ORDER BY sort ASC, created_at DESC，分页返回
```

---

### 4.7 发布商品

- **路径**：`POST /coach/goods/create`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type CoachGoodsCreateReq struct {
    g.Meta      `path:"/coach/goods/create" method:"post" tags:"C端陪玩师" summary:"发布商品"`
    CategoryID  string `json:"categoryId" v:"required#分类ID不能为空" dc:"分类ID"`
    Title       string `json:"title" v:"required|max-length:50#商品标题不能为空|标题最多50字" dc:"商品标题"`
    Description string `json:"description" v:"max-length:500#描述最多500字" dc:"商品描述"`
    Price       int64  `json:"price" v:"required|min:1#价格不能为空|价格必须大于0" dc:"单价(分)"`
    Unit        string `json:"unit" v:"required|in:局,小时,次#单位不能为空|单位值不合法" dc:"单位"`
    Images      string `json:"images" dc:"商品图片(逗号分隔URL)"`
    Sort        int    `json:"sort" dc:"排序（升序）"`
}
```

**响应参数**：

```go
type CoachGoodsCreateRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. 校验分类是否存在
   - SELECT * FROM play_category WHERE id=? AND status=1 AND deleted_at IS NULL
   - 不存在返回"分类不存在或已禁用"
3. 生成 Snowflake ID
4. INSERT INTO play_goods
   - coach_id, category_id, title, description, price, unit, images, sort
   - status=1（默认上架）
5. 更新陪玩师最低价格
   - UPDATE play_coach SET min_price = (SELECT MIN(price) FROM play_goods WHERE coach_id=? AND status=1) WHERE id=?
6. 返回成功
```

---

### 4.8 编辑商品

- **路径**：`PUT /coach/goods/update`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type CoachGoodsUpdateReq struct {
    g.Meta      `path:"/coach/goods/update" method:"put" tags:"C端陪玩师" summary:"编辑商品"`
    GoodsID     string `json:"goodsId" v:"required#商品ID不能为空" dc:"商品ID"`
    CategoryID  string `json:"categoryId" dc:"分类ID"`
    Title       string `json:"title" v:"max-length:50#标题最多50字" dc:"商品标题"`
    Description string `json:"description" v:"max-length:500#描述最多500字" dc:"商品描述"`
    Price       *int64 `json:"price" v:"min:1#价格必须大于0" dc:"单价(分)"`
    Unit        string `json:"unit" v:"in:局,小时,次#单位值不合法" dc:"单位"`
    Images      string `json:"images" dc:"商品图片(逗号分隔URL)"`
    Sort        *int   `json:"sort" dc:"排序"`
}
```

**响应参数**：

```go
type CoachGoodsUpdateRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. 查询商品 WHERE id=? AND coach_id=?（确保只能编辑自己的商品）
   - 不存在返回"商品不存在"
3. 若 categoryId 不为空，校验分类是否存在
4. 构建更新 map（仅更新非空字段）
5. UPDATE play_goods SET ... WHERE id=?
6. 若价格变更，更新陪玩师最低价格
7. 返回成功
```

---

### 4.9 上下架商品

- **路径**：`PUT /coach/goods/status`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type CoachGoodsStatusReq struct {
    g.Meta  `path:"/coach/goods/status" method:"put" tags:"C端陪玩师" summary:"上下架商品"`
    GoodsID string `json:"goodsId" v:"required#商品ID不能为空" dc:"商品ID"`
    Status  int    `json:"status" v:"required|in:0,1#状态不能为空|状态值不合法" dc:"状态:0=下架,1=上架"`
}
```

**响应参数**：

```go
type CoachGoodsStatusRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. 查询商品 WHERE id=? AND coach_id=?
   - 不存在返回"商品不存在"
3. UPDATE play_goods SET status=? WHERE id=?
4. 更新陪玩师最低价格
5. 返回成功
```

---

### 4.10 收入统计

- **路径**：`GET /coach/income`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type CoachIncomeReq struct {
    g.Meta `path:"/coach/income" method:"get" tags:"C端陪玩师" summary:"收入统计"`
}
```

**响应参数**：

```go
type CoachIncomeRes struct {
    g.Meta       `mime:"application/json"`
    TodayIncome  int64 `json:"todayIncome" dc:"今日收入(分)"`
    WeekIncome   int64 `json:"weekIncome" dc:"本周收入(分)"`
    MonthIncome  int64 `json:"monthIncome" dc:"本月收入(分)"`
    TotalIncome  int64 `json:"totalIncome" dc:"累计总收入(分)"`
    TodayOrders  int   `json:"todayOrders" dc:"今日接单数"`
    TotalOrders  int   `json:"totalOrders" dc:"累计接单数"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. 查询 play_order WHERE coach_id=? AND status=3（已完成）
   - 今日收入: SUM(coach_amount) WHERE DATE(finished_at)=CURDATE()
   - 本周收入: SUM(coach_amount) WHERE YEARWEEK(finished_at)=YEARWEEK(NOW())
   - 本月收入: SUM(coach_amount) WHERE YEAR(finished_at)=YEAR(NOW()) AND MONTH(finished_at)=MONTH(NOW())
   - 累计总收入: SUM(coach_amount)
3. 查询接单数
   - 今日接单: COUNT(*) WHERE DATE(created_at)=CURDATE() AND status IN (2,3)
   - 累计接单: COUNT(*) WHERE status IN (2,3)
4. 返回统计数据
```

---

### 4.11 我的接单列表

- **路径**：`GET /coach/orders`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type CoachOrdersReq struct {
    g.Meta   `path:"/coach/orders" method:"get" tags:"C端陪玩师" summary:"我的接单列表"`
    Status   *int `json:"status" dc:"状态筛选:1=待接单,2=进行中,3=已完成"`
    Page     int  `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int  `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type CoachOrdersRes struct {
    g.Meta `mime:"application/json"`
    Total  int              `json:"total" dc:"总数"`
    List   []OrderListItem  `json:"list" dc:"订单列表"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. 查询 play_order WHERE coach_id=?
   - 若 status 不为空，WHERE status=?
3. LEFT JOIN play_member 获取下单用户信息
4. ORDER BY created_at DESC，分页返回
```

---

## 五、商品接口

文件位置：`app/play/api/goods_api.go`

### 5.1 商品列表

- **路径**：`GET /goods/list`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type GoodsListReq struct {
    g.Meta     `path:"/goods/list" method:"get" tags:"C端商品" summary:"商品列表"`
    CategoryID string `json:"categoryId" dc:"分类ID筛选"`
    Keyword    string `json:"keyword" dc:"关键词搜索（标题/描述）"`
    SortBy     string `json:"sortBy" v:"in:price_asc,price_desc,sales,newest#排序值不合法" dc:"排序:price_asc=价格升序,price_desc=价格降序,sales=销量,newest=最新"`
    Page       int    `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize   int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type GoodsListRes struct {
    g.Meta `mime:"application/json"`
    Total  int              `json:"total" dc:"总数"`
    List   []GoodsListItem  `json:"list" dc:"商品列表"`
}

type GoodsListItem struct {
    GoodsID      string  `json:"goodsId" dc:"商品ID"`
    Title        string  `json:"title" dc:"商品标题"`
    CategoryName string  `json:"categoryName" dc:"分类名称"`
    Price        int64   `json:"price" dc:"单价(分)"`
    Unit         string  `json:"unit" dc:"单位"`
    Images       string  `json:"images" dc:"商品图片"`
    CoachID      string  `json:"coachId" dc:"陪玩师ID"`
    CoachName    string  `json:"coachName" dc:"陪玩师昵称"`
    CoachAvatar  string  `json:"coachAvatar" dc:"陪玩师头像"`
    CoachScore   float64 `json:"coachScore" dc:"陪玩师评分"`
    CoachOnline  int     `json:"coachOnline" dc:"陪玩师在线状态"`
    SalesCount   int     `json:"salesCount" dc:"销量"`
}
```

**业务逻辑**：

```
1. 查询 play_goods WHERE status=1
   - JOIN play_coach WHERE play_coach.status=1
   - JOIN play_category
2. 若 categoryId 不为空，WHERE category_id=?
3. 若 keyword 不为空
   - WHERE title LIKE '%keyword%' OR description LIKE '%keyword%'
4. 排序处理
   - price_asc: ORDER BY price ASC
   - price_desc: ORDER BY price DESC
   - sales: ORDER BY sales_count DESC
   - newest: ORDER BY play_goods.created_at DESC
   - 默认: ORDER BY sort ASC, sales_count DESC
5. 分页返回
```

---

### 5.2 商品详情

- **路径**：`GET /goods/detail`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type GoodsDetailReq struct {
    g.Meta  `path:"/goods/detail" method:"get" tags:"C端商品" summary:"商品详情"`
    GoodsID string `json:"goodsId" v:"required#商品ID不能为空" dc:"商品ID"`
}
```

**响应参数**：

```go
type GoodsDetailRes struct {
    g.Meta       `mime:"application/json"`
    GoodsID      string  `json:"goodsId" dc:"商品ID"`
    Title        string  `json:"title" dc:"商品标题"`
    Description  string  `json:"description" dc:"商品描述"`
    CategoryID   string  `json:"categoryId" dc:"分类ID"`
    CategoryName string  `json:"categoryName" dc:"分类名称"`
    Price        int64   `json:"price" dc:"单价(分)"`
    Unit         string  `json:"unit" dc:"单位"`
    Images       string  `json:"images" dc:"商品图片"`
    SalesCount   int     `json:"salesCount" dc:"销量"`
    CoachID      string  `json:"coachId" dc:"陪玩师ID"`
    CoachName    string  `json:"coachName" dc:"陪玩师昵称"`
    CoachAvatar  string  `json:"coachAvatar" dc:"陪玩师头像"`
    CoachScore   float64 `json:"coachScore" dc:"陪玩师评分"`
    CoachOnline  int     `json:"coachOnline" dc:"陪玩师在线状态"`
}
```

**业务逻辑**：

```
1. 查询 play_goods WHERE id=? AND status=1
   - 不存在返回"商品不存在或已下架"
2. JOIN play_coach 获取陪玩师信息
3. JOIN play_category 获取分类名称
4. 组装返回数据
```

---

### 5.3 分类列表

- **路径**：`GET /category/list`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type CategoryListReq struct {
    g.Meta `path:"/category/list" method:"get" tags:"C端商品" summary:"分类列表（树形）"`
}
```

**响应参数**：

```go
type CategoryListRes struct {
    g.Meta `mime:"application/json"`
    List   []CategoryTreeItem `json:"list" dc:"分类树形列表"`
}

type CategoryTreeItem struct {
    CategoryID string             `json:"categoryId" dc:"分类ID"`
    Name       string             `json:"name" dc:"分类名称"`
    Icon       string             `json:"icon" dc:"分类图标"`
    Sort       int                `json:"sort" dc:"排序"`
    Children   []CategoryTreeItem `json:"children" dc:"子分类"`
}
```

**业务逻辑**：

```
1. 查询 play_category WHERE status=1 AND deleted_at IS NULL
2. ORDER BY sort ASC
3. 构建树形结构（parent_id=0 为顶级）
4. 返回树形列表
```

---

## 六、订单接口（需 MemberAuth）

文件位置：`app/play/api/order_api.go`

### 6.1 创建订单

- **路径**：`POST /order/create`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type OrderCreateReq struct {
    g.Meta         `path:"/order/create" method:"post" tags:"C端订单" summary:"创建订单"`
    GoodsID        string `json:"goodsId" v:"required#商品ID不能为空" dc:"商品ID"`
    Quantity       int    `json:"quantity" v:"required|min:1|max:99#数量不能为空|数量最少1|数量最多99" dc:"购买数量"`
    CouponMemberID string `json:"couponMemberId" dc:"使用的优惠券ID（play_coupon_member.id）"`
    Remark         string `json:"remark" v:"max-length:200#备注最多200字" dc:"订单备注"`
}
```

**响应参数**：

```go
type OrderCreateRes struct {
    g.Meta      `mime:"application/json"`
    OrderID     string `json:"orderId" dc:"订单ID"`
    OrderNo     string `json:"orderNo" dc:"订单编号"`
    PayAmount   int64  `json:"payAmount" dc:"应付金额(分)"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查商品信息
   - SELECT * FROM play_goods WHERE id=? AND status=1
   - 不存在返回"商品不存在或已下架"
3. 查陪玩师信息
   - SELECT * FROM play_coach WHERE id=goods.coach_id AND status=1
   - 不可用返回"陪玩师不可用"
   - 不能购买自己的商品
4. 查会员等级折扣
   - SELECT discount FROM play_member_level WHERE id=member.member_level_id AND status=1
5. 计算金额
   - totalAmount = goods.price * quantity
   - discountAmount = totalAmount * (100 - discount) / 100
6. 处理优惠券
   - 若 couponMemberId 不为空
   - 校验优惠券归属、状态、有效期、使用门槛
   - 计算优惠券抵扣金额 couponAmount
7. 计算应付金额
   - payAmount = totalAmount - discountAmount - couponAmount
   - payAmount 最小为 0
8. 开启事务
   - 生成 Snowflake ID + 订单编号（时间戳+随机数）
   - INSERT INTO play_order（status=0 待支付）
   - 若使用优惠券，UPDATE play_coupon_member SET status=1（已锁定）
9. 设置30分钟超时自动取消（延迟队列/定时任务）
10. 返回 orderID + orderNo + payAmount
```

---

### 6.2 我的订单列表

- **路径**：`GET /order/list`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type OrderListReq struct {
    g.Meta   `path:"/order/list" method:"get" tags:"C端订单" summary:"我的订单列表"`
    Status   *int `json:"status" dc:"状态筛选:0=待支付,1=已支付,2=进行中,3=已完成,4=已取消,5=退款中,6=已退款"`
    Page     int  `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int  `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type OrderListRes struct {
    g.Meta `mime:"application/json"`
    Total  int             `json:"total" dc:"总数"`
    List   []OrderListItem `json:"list" dc:"订单列表"`
}

type OrderListItem struct {
    OrderID      string `json:"orderId" dc:"订单ID"`
    OrderNo      string `json:"orderNo" dc:"订单编号"`
    GoodsTitle   string `json:"goodsTitle" dc:"商品标题"`
    GoodsImage   string `json:"goodsImage" dc:"商品图片"`
    CoachID      string `json:"coachId" dc:"陪玩师ID"`
    CoachName    string `json:"coachName" dc:"陪玩师昵称"`
    CoachAvatar  string `json:"coachAvatar" dc:"陪玩师头像"`
    Quantity     int    `json:"quantity" dc:"数量"`
    TotalAmount  int64  `json:"totalAmount" dc:"订单总额(分)"`
    PayAmount    int64  `json:"payAmount" dc:"实付金额(分)"`
    Status       int    `json:"status" dc:"订单状态"`
    StatusText   string `json:"statusText" dc:"状态文本"`
    HasReviewed  bool   `json:"hasReviewed" dc:"是否已评价"`
    CreatedAt    string `json:"createdAt" dc:"下单时间"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_order WHERE member_id=?
   - 若 status 不为空，WHERE status=?
3. LEFT JOIN play_goods, play_coach 获取商品和陪玩师信息
4. 查询是否已评价（LEFT JOIN play_review）
5. ORDER BY created_at DESC，分页返回
```

---

### 6.3 订单详情

- **路径**：`GET /order/detail`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type OrderDetailReq struct {
    g.Meta  `path:"/order/detail" method:"get" tags:"C端订单" summary:"订单详情"`
    OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
}
```

**响应参数**：

```go
type OrderDetailRes struct {
    g.Meta         `mime:"application/json"`
    OrderID        string `json:"orderId" dc:"订单ID"`
    OrderNo        string `json:"orderNo" dc:"订单编号"`
    Status         int    `json:"status" dc:"订单状态"`
    StatusText     string `json:"statusText" dc:"状态文本"`
    GoodsID        string `json:"goodsId" dc:"商品ID"`
    GoodsTitle     string `json:"goodsTitle" dc:"商品标题"`
    GoodsImage     string `json:"goodsImage" dc:"商品图片"`
    Price          int64  `json:"price" dc:"商品单价(分)"`
    Unit           string `json:"unit" dc:"单位"`
    Quantity       int    `json:"quantity" dc:"数量"`
    TotalAmount    int64  `json:"totalAmount" dc:"订单总额(分)"`
    DiscountAmount int64  `json:"discountAmount" dc:"会员折扣金额(分)"`
    CouponAmount   int64  `json:"couponAmount" dc:"优惠券抵扣金额(分)"`
    PayAmount      int64  `json:"payAmount" dc:"实付金额(分)"`
    PayType        string `json:"payType" dc:"支付方式:balance/wechat/alipay"`
    CoachID        string `json:"coachId" dc:"陪玩师ID"`
    CoachName      string `json:"coachName" dc:"陪玩师昵称"`
    CoachAvatar    string `json:"coachAvatar" dc:"陪玩师头像"`
    Remark         string `json:"remark" dc:"订单备注"`
    HasReviewed    bool   `json:"hasReviewed" dc:"是否已评价"`
    RefundReason   string `json:"refundReason" dc:"退款原因"`
    CreatedAt      string `json:"createdAt" dc:"下单时间"`
    PaidAt         string `json:"paidAt" dc:"支付时间"`
    AcceptedAt     string `json:"acceptedAt" dc:"接单时间"`
    FinishedAt     string `json:"finishedAt" dc:"完成时间"`
    CancelledAt    string `json:"cancelledAt" dc:"取消时间"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_order WHERE id=? AND (member_id=? OR coach_id=?)
   - 不存在返回"订单不存在"
3. JOIN play_goods, play_coach 获取关联信息
4. 查询是否已评价
5. 组装返回数据
```

---

### 6.4 取消订单

- **路径**：`POST /order/cancel`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type OrderCancelReq struct {
    g.Meta  `path:"/order/cancel" method:"post" tags:"C端订单" summary:"取消订单"`
    OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
}
```

**响应参数**：

```go
type OrderCancelRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_order WHERE id=? AND member_id=?
   - 不存在返回"订单不存在"
3. 校验状态
   - status=0（待支付）：直接取消
   - status=1（已支付，未接单）：取消并触发全额退款
   - 其他状态返回"当前状态不可取消"
4. 开启事务
   - UPDATE play_order SET status=4, cancelled_at=NOW()
   - 若 status=1，触发退款流程（退回余额或原路退回）
   - 若使用了优惠券，释放优惠券（UPDATE play_coupon_member SET status=0）
5. 返回成功
```

---

### 6.5 申请退款

- **路径**：`POST /order/refund`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type OrderRefundReq struct {
    g.Meta  `path:"/order/refund" method:"post" tags:"C端订单" summary:"申请退款"`
    OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
    Reason  string `json:"reason" v:"required|max-length:200#退款原因不能为空|退款原因最多200字" dc:"退款原因"`
}
```

**响应参数**：

```go
type OrderRefundRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_order WHERE id=? AND member_id=?
3. 校验状态
   - status=2（进行中）：可申请退款
   - 其他状态返回"当前状态不可申请退款"
4. UPDATE play_order SET status=5, refund_reason=?
5. 通知陪玩师（站内消息/推送）
6. 返回成功（等待后台审核或陪玩师同意）
```

---

### 6.6 陪玩师接单

- **路径**：`POST /order/accept`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type OrderAcceptReq struct {
    g.Meta  `path:"/order/accept" method:"post" tags:"C端订单" summary:"陪玩师接单"`
    OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
}
```

**响应参数**：

```go
type OrderAcceptRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. 查询 play_order WHERE id=? AND coach_id=?
   - 不存在返回"订单不存在"
3. 校验状态
   - status=1（已支付）：可接单
   - 其他状态返回"当前状态不可接单"
4. UPDATE play_order SET status=2, accepted_at=NOW()
5. 通知用户（站内消息/推送）
6. 返回成功
```

---

### 6.7 陪玩师完成服务

- **路径**：`POST /order/finish`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type OrderFinishReq struct {
    g.Meta  `path:"/order/finish" method:"post" tags:"C端订单" summary:"陪玩师完成服务"`
    OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
}
```

**响应参数**：

```go
type OrderFinishRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. 查询 play_order WHERE id=? AND coach_id=?
3. 校验状态
   - status=2（进行中）：可完成
   - 其他状态返回"当前状态不可完成"
4. 开启事务
   - UPDATE play_order SET status=3, finished_at=NOW()
   - 计算陪玩师收入（扣除平台抽成）
     coach_amount = pay_amount * (100 - platform_rate) / 100
   - 增加陪玩师余额
     UPDATE play_member SET balance=balance+coach_amount WHERE id=coach.member_id
   - 写入余额流水 play_balance_log
   - 更新陪玩师接单量
     UPDATE play_coach SET order_count=order_count+1
   - 更新商品销量
     UPDATE play_goods SET sales_count=sales_count+quantity
5. 通知用户（站内消息/推送）
6. 返回成功
```

---

## 七、支付接口

文件位置：`app/play/api/payment_api.go`

### 7.1 发起支付

- **路径**：`POST /payment/pay`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type PaymentPayReq struct {
    g.Meta  `path:"/payment/pay" method:"post" tags:"C端支付" summary:"发起支付"`
    OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
    PayType string `json:"payType" v:"required|in:balance,wechat,alipay#支付方式不能为空|支付方式不合法" dc:"支付方式:balance=余额,wechat=微信,alipay=支付宝"`
}
```

**响应参数**：

```go
type PaymentPayRes struct {
    g.Meta    `mime:"application/json"`
    PayResult string `json:"payResult" dc:"支付结果:success=余额支付成功,pending=等待第三方支付"`
    PayParams string `json:"payParams" dc:"第三方支付参数(JSON字符串，前端调起支付SDK用)"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_order WHERE id=? AND member_id=?
   - 不存在返回"订单不存在"
   - status!=0 返回"订单状态不可支付"
3. 校验订单是否超时（created_at + 30分钟 < NOW()）
   - 超时则自动取消并返回"订单已超时"
4. 根据 payType 处理
   A. balance（余额支付）：
      - 查询会员余额 >= payAmount，否则返回"余额不足"
      - 开启事务
        - UPDATE play_member SET balance=balance-payAmount
        - 写入余额流水 play_balance_log（type=pay）
        - UPDATE play_order SET status=1, pay_type='balance', paid_at=NOW()
        - 若使用优惠券，UPDATE play_coupon_member SET status=2（已使用）
      - payResult="success"
   B. wechat（微信支付）：
      - 调用微信统一下单接口
      - 生成预支付参数
      - UPDATE play_order SET pay_type='wechat'
      - payResult="pending", payParams=预支付参数JSON
   C. alipay（支付宝支付）：
      - 调用支付宝下单接口
      - 生成支付参数
      - UPDATE play_order SET pay_type='alipay'
      - payResult="pending", payParams=支付参数JSON
5. 返回 payResult + payParams
```

---

### 7.2 微信支付回调

- **路径**：`POST /payment/wx_callback`
- **登录**：否（回调接口）
- **CoachOnly**：否

**请求参数**：

```go
type PaymentWxCallbackReq struct {
    g.Meta `path:"/payment/wx_callback" method:"post" tags:"C端支付" summary:"微信支付回调"`
    // 微信回调参数由框架自动解析 XML body
}
```

**响应参数**：

```go
type PaymentWxCallbackRes struct {
    g.Meta `mime:"application/xml"`
    // 返回微信要求的 XML 格式
}
```

**业务逻辑**：

```
1. 验证微信签名
   - 签名不通过直接返回失败
2. 解析回调参数，获取 out_trade_no（订单编号）
3. 查询 play_order WHERE order_no=?
   - 不存在或 status!=0，返回成功（幂等处理）
4. 开启事务
   - UPDATE play_order SET status=1, paid_at=NOW(), trade_no=微信交易号
   - 若使用优惠券，UPDATE play_coupon_member SET status=2（已使用）
5. 通知陪玩师有新订单（站内消息/推送）
6. 返回成功 XML
```

---

### 7.3 支付宝支付回调

- **路径**：`POST /payment/alipay_callback`
- **登录**：否（回调接口）
- **CoachOnly**：否

**请求参数**：

```go
type PaymentAlipayCallbackReq struct {
    g.Meta `path:"/payment/alipay_callback" method:"post" tags:"C端支付" summary:"支付宝支付回调"`
    // 支付宝回调参数由框架自动解析 form 表单
}
```

**响应参数**：

```go
type PaymentAlipayCallbackRes struct {
    g.Meta `mime:"text/plain"`
    // 返回 "success" 字符串
}
```

**业务逻辑**：

```
1. 验证支付宝签名
   - 签名不通过返回 "failure"
2. 解析回调参数，获取 out_trade_no（订单编号）
3. 校验 trade_status 是否为 TRADE_SUCCESS
4. 查询 play_order WHERE order_no=?
   - 不存在或 status!=0，返回 "success"（幂等处理）
5. 开启事务
   - UPDATE play_order SET status=1, paid_at=NOW(), trade_no=支付宝交易号
   - 若使用优惠券，UPDATE play_coupon_member SET status=2（已使用）
6. 通知陪玩师有新订单
7. 返回 "success"
```

---

## 八、充值接口

文件位置：`app/play/api/recharge_api.go`

### 8.1 充值方案列表

- **路径**：`GET /recharge/plans`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type RechargePlansReq struct {
    g.Meta `path:"/recharge/plans" method:"get" tags:"C端充值" summary:"充值方案列表"`
}
```

**响应参数**：

```go
type RechargePlansRes struct {
    g.Meta `mime:"application/json"`
    List   []RechargePlanItem `json:"list" dc:"充值方案列表"`
}

type RechargePlanItem struct {
    PlanID     string `json:"planId" dc:"方案ID"`
    Amount     int64  `json:"amount" dc:"充值金额(分)"`
    GiveAmount int64  `json:"giveAmount" dc:"赠送金额(分)"`
    Title      string `json:"title" dc:"方案标题"`
    Tag        string `json:"tag" dc:"标签(如:推荐/热门)"`
    Sort       int    `json:"sort" dc:"排序"`
}
```

**业务逻辑**：

```
1. 查询 play_recharge_plan WHERE status=1 AND deleted_at IS NULL
2. ORDER BY sort ASC
3. 返回列表
```

---

### 8.2 创建充值订单

- **路径**：`POST /recharge/create`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type RechargeCreateReq struct {
    g.Meta  `path:"/recharge/create" method:"post" tags:"C端充值" summary:"创建充值订单"`
    PlanID  string `json:"planId" v:"required#充值方案ID不能为空" dc:"充值方案ID"`
    PayType string `json:"payType" v:"required|in:wechat,alipay#支付方式不能为空|支付方式不合法" dc:"支付方式:wechat=微信,alipay=支付宝"`
}
```

**响应参数**：

```go
type RechargeCreateRes struct {
    g.Meta    `mime:"application/json"`
    OrderID   string `json:"orderId" dc:"充值订单ID"`
    PayParams string `json:"payParams" dc:"第三方支付参数(JSON字符串)"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询充值方案 WHERE id=? AND status=1
   - 不存在返回"充值方案不存在"
3. 生成 Snowflake ID + 充值订单编号
4. INSERT INTO play_recharge_order
   - member_id, plan_id, amount, give_amount, pay_type, status=0
5. 调用第三方支付下单接口
   - wechat: 微信统一下单
   - alipay: 支付宝下单
6. 返回 orderID + payParams
```

---

### 8.3 充值微信回调

- **路径**：`POST /recharge/wx_callback`
- **登录**：否（回调接口）
- **CoachOnly**：否

**请求参数**：

```go
type RechargeWxCallbackReq struct {
    g.Meta `path:"/recharge/wx_callback" method:"post" tags:"C端充值" summary:"充值微信回调"`
}
```

**响应参数**：

```go
type RechargeWxCallbackRes struct {
    g.Meta `mime:"application/xml"`
}
```

**业务逻辑**：

```
1. 验证微信签名
2. 解析回调参数，获取 out_trade_no（充值订单编号）
3. 查询 play_recharge_order WHERE order_no=?
   - 不存在或 status!=0，返回成功（幂等）
4. 开启事务
   - UPDATE play_recharge_order SET status=1, paid_at=NOW(), trade_no=微信交易号
   - 增加会员余额（充值金额 + 赠送金额）
     UPDATE play_member SET balance=balance+(amount+give_amount) WHERE id=member_id
   - 写入余额流水 play_balance_log（type=recharge）
   - 增加会员经验值
     UPDATE play_member SET exp=exp+经验值
   - 检查是否触发升级
5. 返回成功 XML
```

---

### 8.4 充值支付宝回调

- **路径**：`POST /recharge/alipay_callback`
- **登录**：否（回调接口）
- **CoachOnly**：否

**请求参数**：

```go
type RechargeAlipayCallbackReq struct {
    g.Meta `path:"/recharge/alipay_callback" method:"post" tags:"C端充值" summary:"充值支付宝回调"`
}
```

**响应参数**：

```go
type RechargeAlipayCallbackRes struct {
    g.Meta `mime:"text/plain"`
}
```

**业务逻辑**：

```
1. 验证支付宝签名
2. 解析回调参数，获取 out_trade_no
3. 校验 trade_status=TRADE_SUCCESS
4. 查询 play_recharge_order WHERE order_no=?
   - 不存在或 status!=0，返回 "success"（幂等）
5. 开启事务
   - UPDATE play_recharge_order SET status=1, paid_at=NOW(), trade_no=支付宝交易号
   - 增加会员余额（充值金额 + 赠送金额）
   - 写入余额流水
   - 增加经验值，检查升级
6. 返回 "success"
```

---

## 九、优惠券接口

文件位置：`app/play/api/coupon_api.go`

### 9.1 可领取优惠券列表

- **路径**：`GET /coupon/list`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type CouponListReq struct {
    g.Meta `path:"/coupon/list" method:"get" tags:"C端优惠券" summary:"可领取优惠券列表"`
    Page     int `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type CouponListRes struct {
    g.Meta `mime:"application/json"`
    Total  int              `json:"total" dc:"总数"`
    List   []CouponListItem `json:"list" dc:"优惠券列表"`
}

type CouponListItem struct {
    CouponID    string `json:"couponId" dc:"优惠券ID"`
    Title       string `json:"title" dc:"优惠券标题"`
    Type        int    `json:"type" dc:"类型:1=满减,2=折扣,3=无门槛"`
    Amount      int64  `json:"amount" dc:"优惠金额(分)/折扣百分比"`
    MinAmount   int64  `json:"minAmount" dc:"最低使用金额(分)，0=无门槛"`
    StartTime   string `json:"startTime" dc:"有效期开始"`
    EndTime     string `json:"endTime" dc:"有效期结束"`
    TotalCount  int    `json:"totalCount" dc:"发放总量"`
    ClaimedCount int   `json:"claimedCount" dc:"已领取数量"`
    HasClaimed  bool   `json:"hasClaimed" dc:"当前用户是否已领取（未登录为false）"`
}
```

**业务逻辑**：

```
1. 查询 play_coupon WHERE status=1 AND end_time>NOW() AND deleted_at IS NULL
   - 且 claimed_count < total_count（未领完）
2. 若用户已登录（从 header 尝试解析 token，不强制）
   - 查询 play_coupon_member 判断是否已领取，标记 hasClaimed
3. ORDER BY sort ASC, created_at DESC，分页返回
```

---

### 9.2 领取优惠券

- **路径**：`POST /coupon/claim`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type CouponClaimReq struct {
    g.Meta   `path:"/coupon/claim" method:"post" tags:"C端优惠券" summary:"领取优惠券"`
    CouponID string `json:"couponId" v:"required#优惠券ID不能为空" dc:"优惠券ID"`
}
```

**响应参数**：

```go
type CouponClaimRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询优惠券 WHERE id=? AND status=1 AND end_time>NOW()
   - 不存在返回"优惠券不存在或已过期"
3. 校验是否已领完
   - claimed_count >= total_count 返回"优惠券已领完"
4. 校验是否已领取
   - SELECT COUNT(*) FROM play_coupon_member WHERE coupon_id=? AND member_id=?
   - 已领取返回"您已领取过该优惠券"
5. 开启事务
   - 生成 Snowflake ID
   - INSERT INTO play_coupon_member（status=0 未使用）
   - UPDATE play_coupon SET claimed_count=claimed_count+1
6. 返回成功
```

---

### 9.3 我的优惠券列表

- **路径**：`GET /coupon/my_list`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type CouponMyListReq struct {
    g.Meta   `path:"/coupon/my_list" method:"get" tags:"C端优惠券" summary:"我的优惠券列表"`
    Status   *int `json:"status" dc:"状态筛选:0=未使用,1=已锁定,2=已使用,3=已过期"`
    Page     int  `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int  `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type CouponMyListRes struct {
    g.Meta `mime:"application/json"`
    Total  int                `json:"total" dc:"总数"`
    List   []CouponMyListItem `json:"list" dc:"我的优惠券列表"`
}

type CouponMyListItem struct {
    CouponMemberID string `json:"couponMemberId" dc:"用户优惠券ID"`
    CouponID       string `json:"couponId" dc:"优惠券ID"`
    Title          string `json:"title" dc:"优惠券标题"`
    Type           int    `json:"type" dc:"类型:1=满减,2=折扣,3=无门槛"`
    Amount         int64  `json:"amount" dc:"优惠金额(分)/折扣百分比"`
    MinAmount      int64  `json:"minAmount" dc:"最低使用金额(分)"`
    StartTime      string `json:"startTime" dc:"有效期开始"`
    EndTime        string `json:"endTime" dc:"有效期结束"`
    Status         int    `json:"status" dc:"状态:0=未使用,1=已锁定,2=已使用,3=已过期"`
    ClaimedAt      string `json:"claimedAt" dc:"领取时间"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_coupon_member WHERE member_id=?
   - JOIN play_coupon 获取优惠券详情
   - 若 status 不为空，WHERE status=?
   - 自动标记已过期：status=0 且 end_time<NOW() 的标记为 status=3
3. ORDER BY status ASC, claimed_at DESC，分页返回
```

---

### 9.4 下单可用优惠券列表

- **路径**：`GET /coupon/usable`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type CouponUsableReq struct {
    g.Meta      `path:"/coupon/usable" method:"get" tags:"C端优惠券" summary:"下单可用优惠券列表"`
    OrderAmount int64 `json:"orderAmount" v:"required|min:1#订单金额不能为空|订单金额必须大于0" dc:"订单金额(分)"`
}
```

**响应参数**：

```go
type CouponUsableRes struct {
    g.Meta `mime:"application/json"`
    List   []CouponMyListItem `json:"list" dc:"可用优惠券列表"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_coupon_member WHERE member_id=? AND status=0
   - JOIN play_coupon
   - WHERE end_time>NOW()（未过期）
   - WHERE min_amount<=orderAmount（满足使用门槛）
3. ORDER BY amount DESC（优惠金额最大的排前面）
4. 返回列表
```

---

### 9.5 新人券列表

- **路径**：`GET /coupon/new_member`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type CouponNewMemberReq struct {
    g.Meta `path:"/coupon/new_member" method:"get" tags:"C端优惠券" summary:"新人券列表"`
}
```

**响应参数**：

```go
type CouponNewMemberRes struct {
    g.Meta `mime:"application/json"`
    List   []CouponListItem `json:"list" dc:"新人券列表"`
}
```

**业务逻辑**：

```
1. 查询 play_coupon WHERE status=1 AND is_new_member=1 AND end_time>NOW()
2. ORDER BY sort ASC
3. 返回列表
```

---

## 十、活动接口

文件位置：`app/play/api/activity_api.go`

### 10.1 活动列表

- **路径**：`GET /activity/list`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type ActivityListReq struct {
    g.Meta   `path:"/activity/list" method:"get" tags:"C端活动" summary:"活动列表"`
    Status   *int `json:"status" dc:"状态筛选:1=进行中,2=即将开始,3=已结束"`
    Page     int  `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int  `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type ActivityListRes struct {
    g.Meta `mime:"application/json"`
    Total  int                `json:"total" dc:"总数"`
    List   []ActivityListItem `json:"list" dc:"活动列表"`
}

type ActivityListItem struct {
    ActivityID  string `json:"activityId" dc:"活动ID"`
    Title       string `json:"title" dc:"活动标题"`
    Cover       string `json:"cover" dc:"封面图"`
    Description string `json:"description" dc:"活动简介"`
    Type        int    `json:"type" dc:"类型:1=新人活动,2=充值活动,3=消费活动,4=邀请活动"`
    StartTime   string `json:"startTime" dc:"开始时间"`
    EndTime     string `json:"endTime" dc:"结束时间"`
    Status      int    `json:"status" dc:"状态:1=进行中,2=即将开始,3=已结束"`
    JoinCount   int    `json:"joinCount" dc:"参与人数"`
}
```

**业务逻辑**：

```
1. 查询 play_activity WHERE status=1（已发布）AND deleted_at IS NULL
2. 计算活动状态
   - start_time>NOW() → 即将开始
   - start_time<=NOW() AND end_time>=NOW() → 进行中
   - end_time<NOW() → 已结束
3. 若 status 筛选不为空，按计算后的状态筛选
4. ORDER BY start_time DESC，分页返回
```

---

### 10.2 活动详情

- **路径**：`GET /activity/detail`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type ActivityDetailReq struct {
    g.Meta     `path:"/activity/detail" method:"get" tags:"C端活动" summary:"活动详情"`
    ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
}
```

**响应参数**：

```go
type ActivityDetailRes struct {
    g.Meta      `mime:"application/json"`
    ActivityID  string               `json:"activityId" dc:"活动ID"`
    Title       string               `json:"title" dc:"活动标题"`
    Cover       string               `json:"cover" dc:"封面图"`
    Content     string               `json:"content" dc:"活动详情(富文本)"`
    Type        int                  `json:"type" dc:"类型"`
    StartTime   string               `json:"startTime" dc:"开始时间"`
    EndTime     string               `json:"endTime" dc:"结束时间"`
    Status      int                  `json:"status" dc:"状态"`
    JoinCount   int                  `json:"joinCount" dc:"参与人数"`
    HasJoined   bool                 `json:"hasJoined" dc:"当前用户是否已参与"`
    MyProgress  *ActivityProgressInfo `json:"myProgress" dc:"我的进度（已登录且已参与时返回）"`
    Steps       []ActivityStepItem   `json:"steps" dc:"活动步骤列表"`
    Rewards     []ActivityRewardItem `json:"rewards" dc:"奖励列表"`
}

type ActivityStepItem struct {
    StepID      string `json:"stepId" dc:"步骤ID"`
    StepNo      int    `json:"stepNo" dc:"步骤序号"`
    Title       string `json:"title" dc:"步骤标题"`
    Description string `json:"description" dc:"步骤描述"`
    TargetValue int    `json:"targetValue" dc:"目标值"`
}

type ActivityRewardItem struct {
    RewardID    string `json:"rewardId" dc:"奖励ID"`
    Title       string `json:"title" dc:"奖励标题"`
    Type        int    `json:"type" dc:"奖励类型:1=优惠券,2=余额,3=经验值"`
    Value       int64  `json:"value" dc:"奖励值(优惠券ID/金额分/经验值)"`
    Description string `json:"description" dc:"奖励描述"`
}

type ActivityProgressInfo struct {
    CurrentStep  int    `json:"currentStep" dc:"当前步骤序号"`
    IsCompleted  bool   `json:"isCompleted" dc:"是否已完成活动"`
    IsRewarded   bool   `json:"isRewarded" dc:"是否已领取奖励"`
    JoinedAt     string `json:"joinedAt" dc:"参与时间"`
}
```

**业务逻辑**：

```
1. 查询 play_activity WHERE id=? AND status=1
   - 不存在返回"活动不存在"
2. 查询活动步骤
   - SELECT * FROM play_activity_step WHERE activity_id=? ORDER BY step_no ASC
3. 查询活动奖励
   - SELECT * FROM play_activity_reward WHERE activity_id=? ORDER BY sort ASC
4. 若用户已登录（尝试解析 token）
   - 查询 play_activity_member WHERE activity_id=? AND member_id=?
   - 若已参与，返回进度信息
5. 组装返回数据
```

---

### 10.3 报名参与活动

- **路径**：`POST /activity/join`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type ActivityJoinReq struct {
    g.Meta     `path:"/activity/join" method:"post" tags:"C端活动" summary:"报名参与活动"`
    ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
}
```

**响应参数**：

```go
type ActivityJoinRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询活动 WHERE id=? AND status=1
   - 不存在返回"活动不存在"
   - 校验活动时间：start_time<=NOW()<=end_time，否则返回"活动未开始或已结束"
3. 校验是否已参与
   - SELECT COUNT(*) FROM play_activity_member WHERE activity_id=? AND member_id=?
   - 已参与返回"您已参与该活动"
4. 开启事务
   - 生成 Snowflake ID
   - INSERT INTO play_activity_member（current_step=0, is_completed=0, is_rewarded=0）
   - UPDATE play_activity SET join_count=join_count+1
5. 返回成功
```

---

### 10.4 完成活动步骤

- **路径**：`POST /activity/complete_step`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type ActivityCompleteStepReq struct {
    g.Meta     `path:"/activity/complete_step" method:"post" tags:"C端活动" summary:"完成活动步骤"`
    ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
    StepID     string `json:"stepId" v:"required#步骤ID不能为空" dc:"步骤ID"`
}
```

**响应参数**：

```go
type ActivityCompleteStepRes struct {
    g.Meta      `mime:"application/json"`
    CurrentStep int  `json:"currentStep" dc:"当前步骤序号"`
    IsCompleted bool `json:"isCompleted" dc:"是否已完成全部步骤"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询参与记录 WHERE activity_id=? AND member_id=?
   - 不存在返回"您未参与该活动"
   - is_completed=1 返回"活动已完成"
3. 查询步骤 WHERE id=? AND activity_id=?
   - 不存在返回"步骤不存在"
4. 校验步骤顺序
   - step.step_no 应等于 current_step+1，否则返回"请按顺序完成步骤"
5. 开启事务
   - UPDATE play_activity_member SET current_step=step.step_no
   - 查询总步骤数，若 current_step=总步骤数，SET is_completed=1
6. 返回 currentStep + isCompleted
```

---

### 10.5 完成活动

- **路径**：`POST /activity/finish`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type ActivityFinishReq struct {
    g.Meta     `path:"/activity/finish" method:"post" tags:"C端活动" summary:"完成活动"`
    ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
}
```

**响应参数**：

```go
type ActivityFinishRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询参与记录 WHERE activity_id=? AND member_id=?
   - 不存在返回"您未参与该活动"
3. 校验所有步骤是否已完成
   - current_step < 总步骤数 返回"请先完成所有步骤"
4. UPDATE play_activity_member SET is_completed=1, completed_at=NOW()
5. 返回成功
```

---

### 10.6 领取奖励

- **路径**：`POST /activity/claim_reward`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type ActivityClaimRewardReq struct {
    g.Meta     `path:"/activity/claim_reward" method:"post" tags:"C端活动" summary:"领取奖励"`
    ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
    RewardID   string `json:"rewardId" v:"required#奖励ID不能为空" dc:"奖励ID"`
}
```

**响应参数**：

```go
type ActivityClaimRewardRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询参与记录 WHERE activity_id=? AND member_id=?
   - 不存在返回"您未参与该活动"
   - is_completed!=1 返回"请先完成活动"
3. 查询奖励 WHERE id=? AND activity_id=?
   - 不存在返回"奖励不存在"
4. 校验是否已领取该奖励
   - SELECT COUNT(*) FROM play_activity_reward_log WHERE reward_id=? AND member_id=?
   - 已领取返回"您已领取过该奖励"
5. 开启事务，根据奖励类型发放
   A. type=1（优惠券）：
      - INSERT INTO play_coupon_member
   B. type=2（余额）：
      - UPDATE play_member SET balance=balance+value
      - 写入余额流水
   C. type=3（经验值）：
      - UPDATE play_member SET exp=exp+value
      - 检查是否触发升级
6. 记录领取日志 INSERT INTO play_activity_reward_log
7. 更新参与记录 SET is_rewarded=1
8. 返回成功
```

---

### 10.7 我参与的活动列表

- **路径**：`GET /activity/my_list`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type ActivityMyListReq struct {
    g.Meta   `path:"/activity/my_list" method:"get" tags:"C端活动" summary:"我参与的活动列表"`
    Page     int `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type ActivityMyListRes struct {
    g.Meta `mime:"application/json"`
    Total  int                    `json:"total" dc:"总数"`
    List   []ActivityMyListItem   `json:"list" dc:"我参与的活动列表"`
}

type ActivityMyListItem struct {
    ActivityID  string `json:"activityId" dc:"活动ID"`
    Title       string `json:"title" dc:"活动标题"`
    Cover       string `json:"cover" dc:"封面图"`
    Type        int    `json:"type" dc:"类型"`
    StartTime   string `json:"startTime" dc:"开始时间"`
    EndTime     string `json:"endTime" dc:"结束时间"`
    CurrentStep int    `json:"currentStep" dc:"当前步骤"`
    TotalSteps  int    `json:"totalSteps" dc:"总步骤数"`
    IsCompleted bool   `json:"isCompleted" dc:"是否已完成"`
    IsRewarded  bool   `json:"isRewarded" dc:"是否已领奖"`
    JoinedAt    string `json:"joinedAt" dc:"参与时间"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询 play_activity_member WHERE member_id=?
   - JOIN play_activity 获取活动信息
   - 子查询获取每个活动的总步骤数
3. ORDER BY joined_at DESC，分页返回
```

---

## 十一、评价接口

文件位置：`app/play/api/review_api.go`

### 11.1 提交评价

- **路径**：`POST /review/create`
- **登录**：是（MemberAuth）
- **CoachOnly**：否

**请求参数**：

```go
type ReviewCreateReq struct {
    g.Meta      `path:"/review/create" method:"post" tags:"C端评价" summary:"提交评价"`
    OrderID     string  `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
    Score       float64 `json:"score" v:"required|between:1,5#评分不能为空|评分须在1-5之间" dc:"评分(1-5)"`
    Content     string  `json:"content" v:"required|max-length:500#评价内容不能为空|评价内容最多500字" dc:"评价内容"`
    Images      string  `json:"images" dc:"评价图片(逗号分隔URL，最多9张)"`
    IsAnonymous int     `json:"isAnonymous" v:"in:0,1#匿名值不合法" dc:"是否匿名:0=否,1=是" d:"0"`
}
```

**响应参数**：

```go
type ReviewCreateRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_member_id
2. 查询订单 WHERE id=? AND member_id=?
   - 不存在返回"订单不存在"
   - status!=3 返回"订单未完成，不可评价"
3. 校验是否已评价
   - SELECT COUNT(*) FROM play_review WHERE order_id=?
   - 已评价返回"该订单已评价"
4. 校验图片数量
   - images 按逗号分割，超过9张返回"最多上传9张图片"
5. 开启事务
   - 生成 Snowflake ID
   - INSERT INTO play_review
     member_id, order_id, coach_id=order.coach_id, score, content, images, is_anonymous
   - 更新陪玩师评分
     UPDATE play_coach SET
       score = (SELECT AVG(score) FROM play_review WHERE coach_id=?),
       review_count = review_count + 1
     WHERE id=?
   - 增加会员经验值（评价奖励）
     UPDATE play_member SET exp=exp+10
6. 返回成功
```

---

### 11.2 某陪玩师评价列表

- **路径**：`GET /review/list`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type ReviewListReq struct {
    g.Meta   `path:"/review/list" method:"get" tags:"C端评价" summary:"陪玩师评价列表"`
    CoachID  string `json:"coachId" v:"required#陪玩师ID不能为空" dc:"陪玩师ID"`
    Page     int    `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type ReviewListRes struct {
    g.Meta    `mime:"application/json"`
    Total     int              `json:"total" dc:"总数"`
    AvgScore  float64          `json:"avgScore" dc:"平均评分"`
    List      []ReviewListItem `json:"list" dc:"评价列表"`
}

type ReviewListItem struct {
    ReviewID    string  `json:"reviewId" dc:"评价ID"`
    Nickname    string  `json:"nickname" dc:"评价者昵称（匿名显示'匿名用户'）"`
    Avatar      string  `json:"avatar" dc:"评价者头像（匿名显示默认头像）"`
    Score       float64 `json:"score" dc:"评分"`
    Content     string  `json:"content" dc:"评价内容"`
    Images      string  `json:"images" dc:"评价图片"`
    IsAnonymous int     `json:"isAnonymous" dc:"是否匿名"`
    Reply       string  `json:"reply" dc:"陪玩师回复"`
    RepliedAt   string  `json:"repliedAt" dc:"回复时间"`
    CreatedAt   string  `json:"createdAt" dc:"评价时间"`
}
```

**业务逻辑**：

```
1. 查询 play_review WHERE coach_id=?
   - LEFT JOIN play_member 获取评价者信息
2. 匿名评价处理
   - is_anonymous=1 时，nickname="匿名用户"，avatar=默认头像
3. 计算平均评分
   - SELECT AVG(score) FROM play_review WHERE coach_id=?
4. ORDER BY created_at DESC，分页返回
```

---

### 11.3 陪玩师回复评价

- **路径**：`POST /review/reply`
- **登录**：是（MemberAuth + CoachOnly）
- **CoachOnly**：是

**请求参数**：

```go
type ReviewReplyReq struct {
    g.Meta   `path:"/review/reply" method:"post" tags:"C端评价" summary:"陪玩师回复评价"`
    ReviewID string `json:"reviewId" v:"required#评价ID不能为空" dc:"评价ID"`
    Reply    string `json:"reply" v:"required|max-length:200#回复内容不能为空|回复内容最多200字" dc:"回复内容"`
}
```

**响应参数**：

```go
type ReviewReplyRes struct {
    g.Meta `mime:"application/json"`
}
```

**业务逻辑**：

```
1. 从 ctx 获取 jwt_coach_id
2. 查询评价 WHERE id=? AND coach_id=?
   - 不存在返回"评价不存在"
   - reply 不为空返回"已回复过该评价"
3. UPDATE play_review SET reply=?, replied_at=NOW() WHERE id=?
4. 返回成功
```

---

## 十二、搜索接口

文件位置：`app/play/api/search_api.go`

### 12.1 综合搜索

- **路径**：`GET /search`
- **登录**：否
- **CoachOnly**：否

**请求参数**：

```go
type SearchReq struct {
    g.Meta   `path:"/search" method:"get" tags:"C端搜索" summary:"综合搜索"`
    Keyword  string `json:"keyword" v:"required|min-length:1|max-length:50#关键词不能为空|关键词至少1个字符|关键词最多50个字符" dc:"搜索关键词"`
    Type     string `json:"type" v:"in:all,coach,goods#搜索类型不合法" dc:"搜索类型:all=全部,coach=陪玩师,goods=商品" d:"all"`
    Page     int    `json:"page" v:"min:1" dc:"页码" d:"1"`
    PageSize int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}
```

**响应参数**：

```go
type SearchRes struct {
    g.Meta     `mime:"application/json"`
    CoachTotal int              `json:"coachTotal" dc:"陪玩师匹配总数"`
    CoachList  []CoachListItem  `json:"coachList" dc:"陪玩师结果列表"`
    GoodsTotal int              `json:"goodsTotal" dc:"商品匹配总数"`
    GoodsList  []GoodsListItem  `json:"goodsList" dc:"商品结果列表"`
}
```

**业务逻辑**：

```
1. 根据 type 决定搜索范围
2. 搜索陪玩师（type=all 或 type=coach）
   - SELECT * FROM play_coach WHERE status=1
     AND (nickname LIKE '%keyword%' OR skill_desc LIKE '%keyword%')
   - ORDER BY score DESC
   - type=coach 时分页，type=all 时最多返回5条
3. 搜索商品（type=all 或 type=goods）
   - SELECT * FROM play_goods WHERE status=1
     AND (title LIKE '%keyword%' OR description LIKE '%keyword%')
   - JOIN play_coach WHERE play_coach.status=1
   - ORDER BY sales_count DESC
   - type=goods 时分页，type=all 时最多返回10条
4. 返回搜索结果
```

---

## 附录：C端路由注册总览

文件位置：`app/play/internal/router/playapi_router.go`

```go
package router

import (
    "github.com/gogf/gf/v2/net/ghttp"

    "gbaseadmin/app/play/api"
    "gbaseadmin/app/play/internal/middleware"
)

// RegisterPlayAPIRoutes 注册 C端 API 路由
func RegisterPlayAPIRoutes(group *ghttp.RouterGroup) {
    group.Group("/api/playapi", func(group *ghttp.RouterGroup) {

        // ========== 无需登录 ==========

        // 认证
        group.Group("/auth", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.AuthLogin,        // POST /auth/login
                api.AuthSendCode,     // POST /auth/send_code
                api.AuthRefreshToken, // POST /auth/refresh_token
                api.AuthWxLogin,      // POST /auth/wx_login
                api.AuthAlipayLogin,  // POST /auth/alipay_login
            )
        })

        // 陪玩师（公开）
        group.Group("/coach", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.CoachList,   // GET /coach/list
                api.CoachDetail, // GET /coach/detail
            )
        })

        // 商品（公开）
        group.Group("/goods", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.GoodsList,   // GET /goods/list
                api.GoodsDetail, // GET /goods/detail
            )
        })

        // 分类（公开）
        group.Group("/category", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.CategoryList, // GET /category/list
            )
        })

        // 充值方案（公开）
        group.Group("/recharge", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.RechargePlans, // GET /recharge/plans
            )
        })

        // 优惠券（公开）
        group.Group("/coupon", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.CouponList,      // GET /coupon/list
                api.CouponNewMember, // GET /coupon/new_member
            )
        })

        // 活动（公开）
        group.Group("/activity", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.ActivityList,   // GET /activity/list
                api.ActivityDetail, // GET /activity/detail
            )
        })

        // 评价（公开）
        group.Group("/review", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.ReviewList, // GET /review/list
            )
        })

        // 搜索（公开）
        group.Bind(
            api.Search, // GET /search
        )

        // 支付回调（无需登录）
        group.Group("/payment", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.PaymentWxCallback,    // POST /payment/wx_callback
                api.PaymentAlipayCallback, // POST /payment/alipay_callback
            )
        })

        // 充值回调（无需登录）
        group.Group("/recharge", func(group *ghttp.RouterGroup) {
            group.Bind(
                api.RechargeWxCallback,    // POST /recharge/wx_callback
                api.RechargeAlipayCallback, // POST /recharge/alipay_callback
            )
        })

        // ========== 需要登录（MemberAuth） ==========

        group.Group("/", func(group *ghttp.RouterGroup) {
            group.Middleware(middleware.MemberAuth)

            // 会员
            group.Group("/member", func(group *ghttp.RouterGroup) {
                group.Bind(
                    api.MemberInfo,       // GET /member/info
                    api.MemberUpdate,     // PUT /member/update
                    api.MemberSwitchRole, // POST /member/switch_role
                    api.MemberBalanceLog, // GET /member/balance_log
                )
            })

            // 陪玩师申请
            group.Group("/coach", func(group *ghttp.RouterGroup) {
                group.Bind(
                    api.CoachApply,       // POST /coach/apply
                    api.CoachApplyStatus, // GET /coach/apply_status
                )
            })

            // 订单
            group.Group("/order", func(group *ghttp.RouterGroup) {
                group.Bind(
                    api.OrderCreate, // POST /order/create
                    api.OrderList,   // GET /order/list
                    api.OrderDetail, // GET /order/detail
                    api.OrderCancel, // POST /order/cancel
                    api.OrderRefund, // POST /order/refund
                )
            })

            // 支付
            group.Group("/payment", func(group *ghttp.RouterGroup) {
                group.Bind(
                    api.PaymentPay, // POST /payment/pay
                )
            })

            // 充值
            group.Group("/recharge", func(group *ghttp.RouterGroup) {
                group.Bind(
                    api.RechargeCreate, // POST /recharge/create
                )
            })

            // 优惠券（需登录）
            group.Group("/coupon", func(group *ghttp.RouterGroup) {
                group.Bind(
                    api.CouponClaim,  // POST /coupon/claim
                    api.CouponMyList, // GET /coupon/my_list
                    api.CouponUsable, // GET /coupon/usable
                )
            })

            // 活动（需登录）
            group.Group("/activity", func(group *ghttp.RouterGroup) {
                group.Bind(
                    api.ActivityJoin,         // POST /activity/join
                    api.ActivityCompleteStep, // POST /activity/complete_step
                    api.ActivityFinish,       // POST /activity/finish
                    api.ActivityClaimReward,  // POST /activity/claim_reward
                    api.ActivityMyList,       // GET /activity/my_list
                )
            })

            // 评价（需登录）
            group.Group("/review", func(group *ghttp.RouterGroup) {
                group.Bind(
                    api.ReviewCreate, // POST /review/create
                )
            })

            // ========== 陪玩师专属（MemberAuth + CoachOnly） ==========

            group.Group("/", func(group *ghttp.RouterGroup) {
                group.Middleware(middleware.CoachOnly)

                group.Group("/coach", func(group *ghttp.RouterGroup) {
                    group.Bind(
                        api.CoachOnline,      // PUT /coach/online
                        api.CoachMyGoods,     // GET /coach/my_goods
                        api.CoachGoodsCreate, // POST /coach/goods/create
                        api.CoachGoodsUpdate, // PUT /coach/goods/update
                        api.CoachGoodsStatus, // PUT /coach/goods/status
                        api.CoachIncome,      // GET /coach/income
                        api.CoachOrders,      // GET /coach/orders
                    )
                })

                group.Group("/order", func(group *ghttp.RouterGroup) {
                    group.Bind(
                        api.OrderAccept, // POST /order/accept
                        api.OrderFinish, // POST /order/finish
                    )
                })

                group.Group("/review", func(group *ghttp.RouterGroup) {
                    group.Bind(
                        api.ReviewReply, // POST /review/reply
                    )
                })
            })
        })
    })
}
```
