# 陪玩平台 WAP 端页面设计文档

> 技术栈：Taro 4.x + React 18 + NutUI-React + Zustand
> 项目位置：`wap-ui/`（与 admin-go、vue-vben-admin 平级）
> UI 风格：基于 `wap-ui-prototype/index.html` 紫色主题
> 支持平台：H5、微信小程序、抖音小程序

---

## 一、设计系统

### 1.1 主题色变量

```css
:root {
  --primary: #6C5CE7;
  --primary-light: #A29BFE;
  --primary-bg: #F0EEFF;
  --accent: #FF6B6B;
  --accent-light: #FFA8A8;
  --success: #00D2D3;
  --warning: #FECA57;
  --bg: #F8F9FE;
  --card: #FFFFFF;
  --text: #2D3436;
  --text-secondary: #636E72;
  --text-light: #B2BEC3;
  --border: #F1F2F6;
  --shadow: 0 2px 16px rgba(108, 92, 231, 0.08);
  --shadow-lg: 0 8px 32px rgba(108, 92, 231, 0.12);
  --radius: 16px;
  --radius-sm: 12px;
}
```

### 1.2 字体规范

- 标题：PingFang SC, 17px, font-weight: 700
- 正文：14px, font-weight: 400
- 辅助：12px, color: var(--text-secondary)
- 小字：10-11px, color: var(--text-light)

### 1.3 通用组件

| 组件 | 说明 | 使用场景 |
|------|------|----------|
| NavBar | 顶部导航栏，白色背景+返回箭头+标题 | 所有二级页面 |
| TabBar | 底部5Tab导航 | 5个主Tab页 |
| CoachCard | 陪玩师卡片 | 首页推荐、列表页 |
| GoodsCard | 商品卡片 | 首页热门、商品列表 |
| ActivityCard | 活动卡片 | 首页活动、活动列表 |
| OrderCard | 订单卡片 | 订单列表 |
| CouponCard | 优惠券卡片 | 优惠券列表 |
| PriceText | 价格显示（红色大字+单位） | 商品/订单 |
| StatusTag | 状态标签 | 订单/审核状态 |
| EmptyState | 空状态占位 | 列表为空时 |
| LoadMore | 上拉加载更多 | 所有列表页 |

---

## 二、项目目录结构

```
wap-ui/
├── src/
│   ├── app.ts                          # 应用入口
│   ├── app.config.ts                   # Taro 配置（pages/tabBar）
│   ├── app.scss                        # 全局样式 + 主题变量
│   ├── api/                            # API 请求层
│   │   ├── request.ts                  # Taro.request 封装 + Token 拦截
│   │   ├── auth.ts                     # 认证接口
│   │   ├── member.ts                   # 会员接口
│   │   ├── coach.ts                    # 陪玩师接口
│   │   ├── goods.ts                    # 商品接口
│   │   ├── order.ts                    # 订单接口
│   │   ├── payment.ts                  # 支付接口
│   │   ├── recharge.ts                 # 充值接口
│   │   ├── coupon.ts                   # 优惠券接口
│   │   ├── activity.ts                 # 活动接口
│   │   └── review.ts                   # 评价接口
│   ├── store/                          # Zustand 状态管理
│   │   ├── auth.ts                     # 登录态 + Token + 用户信息
│   │   └── cart.ts                     # 下单临时状态
│   ├── components/                     # 公共组件
│   │   ├── CoachCard/index.tsx
│   │   ├── GoodsCard/index.tsx
│   │   ├── ActivityCard/index.tsx
│   │   ├── OrderCard/index.tsx
│   │   ├── CouponCard/index.tsx
│   │   ├── PriceText/index.tsx
│   │   ├── StatusTag/index.tsx
│   │   ├── EmptyState/index.tsx
│   │   └── LoadMore/index.tsx
│   ├── pages/                          # 页面
│   │   ├── index/index.tsx             # 首页
│   │   ├── category/index.tsx          # 分类页
│   │   ├── search/index.tsx            # 搜索页
│   │   ├── coach/
│   │   │   ├── list.tsx                # 陪玩师列表
│   │   │   ├── detail.tsx              # 陪玩师详情
│   │   │   ├── apply.tsx               # 申请成为陪玩师
│   │   │   └── apply-status.tsx        # 申请状态
│   │   ├── goods/
│   │   │   └── detail.tsx              # 商品详情
│   │   ├── order/
│   │   │   ├── confirm.tsx             # 下单确认
│   │   │   ├── list.tsx                # 订单列表
│   │   │   ├── detail.tsx              # 订单详情
│   │   │   ├── pay.tsx                 # 支付页
│   │   │   └── review.tsx              # 评价页
│   │   ├── recharge/
│   │   │   └── index.tsx               # 充值页
│   │   ├── coupon/
│   │   │   ├── list.tsx                # 我的优惠券
│   │   │   └── center.tsx              # 领券中心
│   │   ├── activity/
│   │   │   ├── list.tsx                # 活动列表（Tab页）
│   │   │   └── detail.tsx              # 活动详情
│   │   ├── message/
│   │   │   └── index.tsx               # 消息列表（Tab页）
│   │   ├── mine/
│   │   │   ├── index.tsx               # 我的（Tab页）
│   │   │   ├── profile.tsx             # 编辑资料
│   │   │   ├── balance.tsx             # 余额明细
│   │   │   └── settings.tsx            # 设置
│   │   ├── login/
│   │   │   └── index.tsx               # 登录页
│   │   └── workspace/                  # 陪玩师工作台
│   │       ├── index.tsx               # 工作台首页
│   │       ├── orders.tsx              # 接单管理
│   │       ├── income.tsx              # 收入统计
│   │       └── goods.tsx               # 商品管理
│   └── utils/                          # 工具函数
│       ├── format.ts                   # 金额/日期格式化
│       ├── storage.ts                  # 本地存储封装
│       └── auth.ts                     # Token 管理
```

---

## 三、TabBar 配置

| Tab | 图标 | 页面路径 | 说明 |
|-----|------|----------|------|
| 首页 | home | pages/index/index | 主页面 |
| 分类 | category | pages/category/index | 商品分类 |
| 活动 | gift | pages/activity/list | 活动列表 |
| 消息 | message | pages/message/index | 消息中心 |
| 我的 | user | pages/mine/index | 个人中心 |

TabBar 样式：白色背景，顶部 1px 分割线，选中态紫色(--primary)，未选中灰色(--text-light)。

---

## 四、页面详细设计（共 27 个页面）

### 4.1 首页（pages/index/index）

**布局从上到下**：

**① 状态栏 + 搜索头部**
- 紫色渐变背景（linear-gradient 135deg, #6C5CE7 → #A29BFE）
- 左侧：定位城市图标+城市名
- 右侧：通知铃铛 + 扫码图标（白色SVG）
- 搜索栏：rgba(255,255,255,0.2) 背景，backdrop-filter: blur(10px)，圆角24px
- placeholder: "搜索陪玩师、游戏、服务..."
- 底部圆角裁切 border-radius: 0 0 24px 24px
- 点击搜索栏跳转 `/pages/search/index`

**② Banner 轮播**
- padding: 16px，高度 160px，圆角 16px
- 渐变背景色（每张不同），左侧文字+右侧圆形图标区
- 底部圆点指示器（白色活跃态宽18px，非活跃6px）
- 自动轮播 3 秒
- 数据来源：活动列表前3条 或 后台配置

**③ 快捷导航（5宫格）**
- 5列 grid，gap: 8px，padding: 20px 16px 8px
- 每项：渐变色圆角图标(48x48, radius:14px) + 11px文字
- 游戏陪玩(紫#6C5CE7) | 语音聊天(红#FF6B6B) | 看电影(青#00D2D3) | 唱歌(黄#FECA57) | 更多(粉#fd79a8)
- 点击跳转对应分类：`/pages/coach/list?categoryId=xxx`

**④ 热门活动（横向滚动）**
- section-header：标题"热门活动"(17px 加粗) + "查看全部 >"(12px 灰色)
- 横向滚动，-webkit-overflow-scrolling: touch，隐藏滚动条
- ActivityCard（min-width: 260px）：
  - 顶部 120px 渐变色区域 + 左上角类型角标(半透明黑底) + 活动标语
  - 底部白色区域：活动名称(14px加粗) + 时间·参与人数(11px灰) + 奖励标签
- 接口：`GET /api/playapi/activity/list?pageSize=5`

**⑤ 推荐陪玩师（2列网格）**
- section-header：标题"推荐陪玩师" + "查看全部 >"
- 2列 grid，gap: 12px
- CoachCard：
  - 头像区 140px 高，灰色渐变背景，居中大头像图标
  - 右上角：在线绿点(8px, #00b894, box-shadow: 0 0 0 2px #fff)
  - 左下角：等级标签(rgba(0,0,0,0.4) backdrop-filter:blur, 10px字)
  - 信息区 padding: 10px 12px：
    - 昵称(14px加粗)
    - 标签行：紫色小标签(10px, padding:2px 6px, radius:6px, --primary-bg底色)
    - 底部：价格(16px加粗红色 + "元/局"11px灰) | 评分(黄色星+分数)
- 接口：`GET /api/playapi/coach/list?pageSize=4&sort=score`

**⑥ 热门商品（纵向列表）**
- section-header：标题"热门服务" + "查看全部 >"
- 纵向卡片列表，gap: 14px
- GoodsCard：
  - 封面图 140px 高，底部渐变遮罩(transparent→rgba(0,0,0,0.6))
  - 封面上叠加标题(白色16px加粗, text-shadow)
  - 信息区：左侧(陪玩师头像+昵称+描述) | 右侧(价格+下单按钮)
  - 下单按钮：紫色渐变圆角按钮(12px, padding:6px 16px)
- 接口：`GET /api/playapi/goods/list?pageSize=10&sort=sales`
- 触底加载更多

**⑦ 底部留白** 80px

**交互**：下拉刷新 + 触底加载（热门商品分页）

---

### 4.2 分类页（pages/category/index）

**左右分栏布局**：

**① 顶部搜索栏**
- 白色背景，搜索框，点击跳转搜索页

**② 左侧分类导航（宽90px）**
- 灰色背景(#F5F6FA)，纵向滚动，高度撑满
- 每项：分类图标(20px) + 名称(12px)，padding: 14px 8px
- 选中态：白色背景 + 左侧3px紫色竖条 + 文字变紫色加粗
- 数据来源：`GET /api/playapi/category/list`（树形，取一级分类）

**③ 右侧商品列表**
- 白色背景，padding: 12px
- 顶部：当前分类名(15px加粗) + 商品数量(12px灰)
- 2列网格商品卡片，gap: 10px
- 每张卡片：封面图(正方形) + 商品名(13px, 2行截断) + 价格(红色14px加粗)
- 下拉刷新 + 触底加载
- 接口：`GET /api/playapi/goods/list?categoryId=xxx&pageNum=1&pageSize=10`

---

### 4.3 搜索页（pages/search/index）

**① 搜索头部**
- 返回按钮 + 搜索输入框（自动聚焦，圆角24px）+ "搜索"文字按钮
- placeholder: "搜索陪玩师、游戏、服务..."

**② 未输入时**
- 搜索历史：标题 + 清空按钮，标签流式布局（灰色圆角标签）
- 热门搜索：标签流式布局（紫色边框标签）
- 历史存储在本地 Storage

**③ 输入后 — 搜索结果**
- Tab 切换栏：陪玩师 | 商品（下划线指示器，紫色）
- 陪玩师结果：CoachCard 2列网格
- 商品结果：GoodsCard 纵向列表
- 空结果：EmptyState "未找到相关内容"
- 接口：`GET /api/playapi/search?keyword=xxx&type=coach|goods`

---

### 4.4 陪玩师列表页（pages/coach/list）

**① NavBar**："陪玩师"

**② 筛选栏（吸顶 sticky）**
- 横向滚动分类标签：全部 | 游戏陪玩 | 语音聊天 | 看电影 | 唱歌 | ...
- 选中态：紫色背景白色文字，未选中：灰色背景
- 排序按钮行：综合 | 评分最高 | 接单最多 | 价格最低

**③ 陪玩师列表**
- 2列网格，CoachCard 组件
- 下拉刷新 + 触底加载
- 接口：`GET /api/playapi/coach/list?categoryId=xxx&sort=xxx&pageNum=1&pageSize=10`

---

### 4.5 陪玩师详情页（pages/coach/detail）

**① 顶部大图区域**
- 封面图全宽，高度 280px
- 底部渐变遮罩
- 左上角：返回按钮（白色圆形半透明背景）
- 右上角：分享按钮
- 底部叠加：在线状态点 + 昵称(20px白色加粗) + 等级标签

**② 信息卡片（白色圆角，margin-top: -20px，叠在大图上）**
- 评分：大字评分(24px黄色) + 星星 + 评分人数
- 数据行：接单数 | 好评率 | 响应时间
- 个人简介：intro 文字（可展开/收起）
- 标签：技能标签列表（紫色小标签）

**③ Tab 切换区域（吸顶）**
- 服务项目 | 用户评价

**④ 服务项目 Tab**
- 商品列表，纵向排列
- 每项：封面小图(80x80) + 商品名+描述 + 价格 + "立即预约"按钮
- 接口：`GET /api/playapi/goods/list?coachId=xxx`

**⑤ 用户评价 Tab**
- 评价统计：平均分 + 评分分布条
- 评价列表：
  - 用户头像+昵称(匿名则显示"匿名用户") + 评分星星 + 时间
  - 评价内容 + 图片（可点击放大）
  - 陪玩师回复（灰色背景区块）
- 接口：`GET /api/playapi/review/list?coachId=xxx&pageNum=1`

**⑥ 底部固定操作栏**
- 左侧：收藏按钮 | 客服按钮
- 右侧：紫色渐变大按钮 "立即预约"（跳转商品选择或直接下单）

---

### 4.6 商品详情页（pages/goods/detail）

**① NavBar**："服务详情"

**② 商品封面**
- 轮播图（如果有多图）或单图，高度 240px

**③ 商品信息卡片**
- 价格：大红色(24px加粗) ¥xx /局
- 商品名称(17px加粗)
- 销量：已售 xxx 单(12px灰)
- 计量单位说明

**④ 陪玩师信息栏**
- 左侧：头像(40px圆形) + 昵称 + 等级标签 + 评分
- 右侧：在线状态 + "查看主页 >"
- 点击跳转陪玩师详情

**⑤ 商品详情**
- desc_content 富文本渲染
- 图文混排

**⑥ 用户评价预览**
- 显示前3条评价
- "查看全部评价 >" 跳转评价列表

**⑦ 底部固定操作栏**
- 左侧：客服 | 收藏
- 右侧：紫色渐变大按钮 "立即下单"
- 点击跳转下单确认页

---

### 4.7 下单确认页（pages/order/confirm）

**① NavBar**："确认订单"

**② 商品信息**
- 商品封面小图(80x80) + 商品名 + 陪玩师昵称 + 单价
- 数量选择器（NutUI InputNumber，最小1）

**③ 优惠券选择**
- 点击展开优惠券选择弹窗
- 显示可用优惠券数量，已选优惠券面额
- 无可用券显示"暂无可用优惠券"
- 接口：`GET /api/playapi/coupon/usable?amount=xxx`

**④ 订单备注**
- 输入框，placeholder: "给陪玩师留言（选填）"

**⑤ 金额明细**
- 商品金额：¥xx.xx
- 会员折扣：-¥xx.xx（有折扣时显示，绿色）
- 优惠券：-¥xx.xx（选了券时显示，绿色）
- 分割线
- 实付金额：¥xx.xx（红色大字）

**⑥ 底部固定**
- 左侧：合计 ¥xx.xx（红色）
- 右侧：紫色渐变大按钮 "提交订单"
- 点击调用：`POST /api/playapi/order/create`
- 成功后跳转支付页

---

### 4.8 支付页（pages/order/pay）

**① NavBar**："订单支付"

**② 支付金额**
- 居中大字显示：¥xx.xx（红色 32px）
- 订单编号(12px灰)

**③ 支付方式选择**
- 单选列表，白色卡片：
  - 微信支付（绿色图标）
  - 支付宝支付（蓝色图标）
  - 余额支付（紫色图标）— 显示当前余额，余额不足时置灰+提示"余额不足，去充值"

**④ 底部固定**
- 紫色渐变大按钮 "确认支付 ¥xx.xx"
- 点击调用：`POST /api/playapi/payment/pay`
- 余额支付直接完成，微信/支付宝唤起第三方支付

**⑤ 支付结果**
- 成功：跳转订单详情页，显示成功动画
- 失败：提示重试

**⑥ 倒计时提示**
- 顶部显示"请在 xx:xx 内完成支付"（30分钟倒计时）

---

### 4.9 订单列表页（pages/order/list）

**① NavBar**："我的订单"

**② 状态 Tab 栏（吸顶）**
- 全部 | 待支付 | 已支付 | 进行中 | 已完成 | 退款
- 选中态：紫色文字 + 底部紫色下划线(2px)
- 对应 order_status：全部(-1) | 0 | 1 | 2 | 3 | 5,6

**③ 订单列表**
- OrderCard 组件：
  - 顶部行：订单编号(12px灰) + 状态标签（不同颜色）
    - 待支付：橙色 | 已支付：蓝色 | 进行中：紫色 | 已完成：绿色 | 已取消：灰色 | 退款中：红色 | 已退款：灰色
  - 中间：商品封面小图(60x60) + 商品名 + 陪玩师昵称 + 数量x单价
  - 底部行：实付金额(红色加粗) + 操作按钮
    - 待支付：去支付(紫色实心) + 取消订单(灰色边框)
    - 已支付：取消订单(灰色边框) + 催接单
    - 进行中：申请退款(灰色边框)
    - 已完成（未评价）：去评价(紫色实心)
    - 已完成（已评价）：再来一单(紫色边框)
- 空状态：EmptyState "暂无订单"
- 下拉刷新 + 触底加载
- 接口：`GET /api/playapi/order/list?status=xxx&pageNum=1&pageSize=10`

---

### 4.10 订单详情页（pages/order/detail）

**① NavBar**："订单详情"

**② 订单状态区域**
- 顶部彩色背景区域（根据状态不同颜色）
- 大字状态文字 + 状态说明
- 待支付：显示倒计时 "剩余 xx:xx 自动取消"

**③ 商品信息**
- 商品封面(80x80) + 商品名 + 陪玩师昵称 + 单价 x 数量

**④ 陪玩师信息**
- 头像 + 昵称 + 在线状态 + "联系陪玩师"按钮

**⑤ 金额明细**
- 商品金额 / 会员折扣 / 优惠券 / 实付金额

**⑥ 订单信息**
- 订单编号（可复制）
- 下单时间 / 支付时间 / 支付方式
- 服务开始时间 / 完成时间（有则显示）
- 订单备注

**⑦ 底部操作栏**（根据状态显示不同按钮）
- 待支付：去支付 + 取消订单
- 已支付：取消订单
- 进行中：申请退款
- 已完成：去评价 / 再来一单

---

### 4.11 评价页（pages/order/review）

**① NavBar**："评价服务"

**② 陪玩师信息**
- 头像 + 昵称 + 商品名

**③ 评分**
- 5颗星评分（NutUI Rate 组件），默认5星
- 评分文字提示：1差评 2一般 3还行 4满意 5非常满意

**④ 评价内容**
- 多行文本输入框，placeholder: "分享你的体验，帮助其他用户~"
- 最多 500 字，显示字数统计

**⑤ 上传图片**
- 最多 9 张图片，NutUI Uploader 组件
- 支持拍照 + 相册选择

**⑥ 匿名评价**
- Switch 开关 "匿名评价"

**⑦ 底部固定**
- 紫色渐变大按钮 "提交评价"
- 接口：`POST /api/playapi/review/create`
- 成功后返回订单详情页

---

### 4.12 充值页（pages/recharge/index）

**① NavBar**："充值中心"

**② 当前余额**
- 紫色渐变背景卡片
- 大字余额(28px白色加粗) + "当前余额"标签
- 余额明细入口 "查看明细 >"

**③ 充值方案列表**
- 网格布局（2列），gap: 12px
- 每个方案卡片：
  - 充值金额(20px加粗)
  - 赠送金额（红色标签 "送¥xx"，有赠送时显示）
  - 选中态：紫色边框 + 紫色背景浅色
- 接口：`GET /api/playapi/recharge/plans`

**④ 支付方式**
- 微信支付 / 支付宝支付（单选）

**⑤ 底部固定**
- 紫色渐变大按钮 "立即充值 ¥xx"
- 接口：`POST /api/playapi/recharge/create`

---

### 4.13 优惠券列表页（pages/coupon/list）

**① NavBar**："我的优惠券"，右上角 "领券中心"

**② 状态 Tab**
- 未使用 | 已使用 | 已过期
- 对应 use_status: 0 | 1 | 2

**③ 优惠券列表**
- CouponCard 组件：
  - 左侧：面额区域（红色大字 ¥xx 或 xx折）
  - 右侧：优惠券名称 + 使用条件("满xx可用") + 有效期
  - 右下角：状态按钮
    - 未使用："去使用"(紫色) → 跳转首页
    - 已使用：灰色"已使用"
    - 已过期：灰色"已过期"
- 卡片样式：左侧圆形缺口装饰（经典优惠券造型）
- 接口：`GET /api/playapi/coupon/my_list?useStatus=xxx`

---

### 4.14 领券中心（pages/coupon/center）

**① NavBar**："领券中心"

**② 优惠券列表**
- CouponCard 组件（同上，但按钮不同）：
  - 可领取："立即领取"(紫色实心按钮)
  - 已领取："已领取"(灰色)
  - 已领完："已领完"(灰色)
- 新人专区（顶部横幅，仅新用户可见）
- 接口：`GET /api/playapi/coupon/list`
- 领取接口：`POST /api/playapi/coupon/claim`

---

### 4.15 活动列表页（pages/activity/list）— TabBar 页

**① 顶部**
- 标题"活动中心"(大字) + 搜索图标

**② 活动分类 Tab**
- 全部 | 充值活动 | 下单活动 | 新人活动 | 步骤活动
- 对应 type: 全部(0) | 1 | 2 | 3 | 4

**③ 活动列表**
- ActivityCard 组件（纵向排列，比首页横向卡片更大）：
  - 顶部：渐变色大图区域(180px)，活动标语 + 类型角标
  - 底部：活动名称 + 时间范围 + 参与人数 + 奖励标签列表
  - 状态：进行中(绿色) / 即将开始(橙色) / 已结束(灰色)
- 下拉刷新 + 触底加载
- 接口：`GET /api/playapi/activity/list?type=xxx&pageNum=1`

---

### 4.16 活动详情页（pages/activity/detail）

**① 顶部大图**
- 活动封面图全宽，高度 200px
- 渐变遮罩 + 活动名称 + 时间范围

**② 活动信息卡片**
- 活动状态 + 参与人数/上限
- 活动描述（desc_content 富文本）

**③ 奖励列表**
- 标题"活动奖励"
- 奖励卡片列表：
  - 图标(余额/优惠券/经验/等级) + 奖励名称 + 奖励值
  - 如：余额图标 + "赠送余额" + "¥30.00"

**④ 步骤引导（type=4 图文步骤活动时显示）**
- 标题"活动步骤"
- 步骤列表（纵向时间线样式）：
  - 步骤序号圆圈（已完成：紫色实心✓ / 当前：紫色边框 / 未完成：灰色）
  - 步骤标题 + 描述文字
  - 步骤示例图片（可点击放大）
  - 当前步骤：显示"完成此步骤"按钮
- 接口：`POST /api/playapi/activity/complete_step`

**⑤ 底部固定操作栏**
- 未参与："立即参与"(紫色) → `POST /api/playapi/activity/join`
- 进行中（步骤活动）："继续完成"(紫色)
- 已完成未领奖："领取奖励"(红色) → `POST /api/playapi/activity/claim_reward`
- 已领奖："已完成"(灰色)
- 已结束："活动已结束"(灰色禁用)

---

### 4.17 消息列表页（pages/message/index）— TabBar 页

**① 顶部**
- 标题"消息"(大字) + 右上角"全部已读"

**② 消息列表**
- 消息卡片：
  - 左侧：消息类型图标(40px圆形，不同颜色背景)
    - 订单消息：紫色购物袋图标
    - 系统通知：蓝色铃铛图标
    - 活动消息：红色礼物图标
  - 中间：消息标题(14px) + 消息摘要(12px灰, 单行截断) + 时间(10px灰)
  - 右侧：未读红点(8px)
- 空状态：EmptyState "暂无消息"
- 点击跳转对应详情（订单详情/活动详情等）

> 注：消息系统为占位设计，后续可接入 WebSocket/IM

---

### 4.18 我的页面（pages/mine/index）— TabBar 页

**① 用户信息区域（紫色渐变背景，圆角底部）**
- 未登录：显示默认头像 + "点击登录"
- 已登录：
  - 头像(64px圆形, 白色边框3px) + 昵称(18px白色加粗) + 会员等级标签
  - 手机号(14px白色半透明)
  - 右侧：设置图标 → 跳转设置页
- 身份切换入口（is_coach=1 时显示）：
  - 当前身份标签 "会员模式" / "陪玩师模式"
  - 点击切换，调用 `POST /api/playapi/member/switch_role`

**② 资产栏（白色卡片，margin-top: -20px 叠在紫色区域上）**
- 3列等分：
  - 余额：¥xx.xx → 跳转余额明细
  - 优惠券：x 张 → 跳转我的优惠券
  - 经验值：xxx → 显示等级进度

**③ 订单快捷入口**
- 标题"我的订单" + "查看全部 >"
- 4个图标入口（横向等分）：
  - 待支付(钱包图标) | 进行中(时钟图标) | 已完成(勾选图标) | 退款(退款图标)
  - 有数量时显示红色角标数字

**④ 功能列表（白色卡片）**
- Cell 列表（NutUI Cell 组件）：
  - 充值中心 → /pages/recharge/index
  - 领券中心 → /pages/coupon/center
  - 我的优惠券 → /pages/coupon/list
  - 余额明细 → /pages/mine/balance
  - 我的评价 → 评价列表
  - 分割线
  - 申请成为陪玩师 → /pages/coach/apply（is_coach=0 时显示）
  - 陪玩师工作台 → /pages/workspace/index（is_coach=1 且 currentRole=coach 时显示）
  - 分割线
  - 设置 → /pages/mine/settings
  - 关于我们

---

### 4.19 编辑资料页（pages/mine/profile）

**① NavBar**："编辑资料"

**② 表单**
- 头像：点击更换，NutUI Uploader（圆形裁剪）
- 昵称：输入框
- 性别：Radio 选择（男/女/保密）

**③ 底部固定**
- 紫色渐变大按钮 "保存"
- 接口：`PUT /api/playapi/member/update`

---

### 4.20 余额明细页（pages/mine/balance）

**① NavBar**："余额明细"

**② 当前余额**
- 顶部紫色卡片，大字余额 + "去充值"按钮

**③ 流水列表**
- 筛选 Tab：全部 | 充值 | 消费 | 退款 | 活动赠送 | 提现
- 对应 biz_type: 全部(0) | 1 | 2 | 3 | 4 | 5
- 每条流水：
  - 左侧：类型图标 + 描述(remark) + 时间(12px灰)
  - 右侧：金额（正数绿色+号，负数红色-号）
- 下拉刷新 + 触底加载
- 接口：`GET /api/playapi/member/balance_log?bizType=xxx&pageNum=1`

---

### 4.21 设置页（pages/mine/settings）

**① NavBar**："设置"

**② 设置列表**
- Cell 列表：
  - 修改密码 → 弹窗输入旧密码+新密码
  - 清除缓存 → 确认弹窗
  - 用户协议 → 跳转协议页
  - 隐私政策 → 跳转隐私页
  - 关于我们 → 版本号等信息

**③ 底部**
- 红色文字按钮 "退出登录"
- 确认弹窗 → 清除 Token + 跳转登录页

---

### 4.22 登录页（pages/login/index）

**① 顶部 Logo 区域**
- 紫色渐变背景（占屏幕上半部分 40%）
- 居中：App Logo + "PlayBuddy" 文字 + "你的专属陪玩平台" slogan

**② 登录表单（白色卡片，圆角顶部，叠在紫色区域上）**
- 手机号输入框（带 +86 前缀图标）
- 验证码输入框 + "获取验证码"按钮（60秒倒计时）
- 紫色渐变大按钮 "登录 / 注册"
- 接口：
  - 发送验证码：`POST /api/playapi/auth/send_code`
  - 登录：`POST /api/playapi/auth/login`

**③ 第三方登录**
- 分割线 "其他登录方式"
- 图标按钮：微信(绿色) | 支付宝(蓝色)
- 接口：`POST /api/playapi/auth/wx_login` / `POST /api/playapi/auth/alipay_login`

**④ 底部**
- 协议勾选："登录即同意《用户协议》和《隐私政策》"

---

### 4.23 陪玩师申请页（pages/coach/apply）

**① NavBar**："申请成为陪玩师"

**② 申请表单**
- 真实姓名：输入框（必填）
- 身份证号：输入框（必填，校验格式）
- 身份证正面照：图片上传（必填）
- 身份证反面照：图片上传（必填）
- 技能描述：多行文本输入框，placeholder: "介绍你擅长的游戏和服务..."

**③ 申请须知**
- 折叠面板，展示申请条件和注意事项

**④ 底部固定**
- 紫色渐变大按钮 "提交申请"
- 接口：`POST /api/playapi/coach/apply`
- 成功后跳转申请状态页

---

### 4.24 申请状态页（pages/coach/apply-status）

**① NavBar**："申请状态"

**② 状态展示（居中）**
- 审核中(audit_status=0)：
  - 橙色时钟大图标
  - "审核中" 大字
  - "预计 1-3 个工作日内完成审核" 说明文字
- 已通过(audit_status=1)：
  - 绿色勾选大图标
  - "审核通过" 大字
  - "恭喜你成为陪玩师！" + "前往工作台"按钮
- 已拒绝(audit_status=2)：
  - 红色叉号大图标
  - "审核未通过" 大字
  - 拒绝原因(audit_remark)
  - "重新申请"按钮

- 接口：`GET /api/playapi/coach/apply_status`

---

### 4.25 陪玩师工作台首页（pages/workspace/index）

**① NavBar**："工作台" + 右上角在线/离线切换开关

**② 今日数据概览（紫色渐变卡片）**
- 4格数据：今日接单 | 今日收入 | 总评分 | 好评率
- 数字大字(20px白色加粗) + 标签(12px白色半透明)

**③ 快捷操作（4宫格）**
- 接单管理 | 商品管理 | 收入统计 | 我的评价
- 每项：图标(40px) + 文字，点击跳转对应页面

**④ 待处理订单（最近3条）**
- 标题"待处理" + "查看全部 >"
- 简化版 OrderCard：商品名 + 用户昵称 + 金额 + "接单"按钮
- 接口：`GET /api/playapi/coach/orders?status=1&pageSize=3`

**⑤ 收入趋势（简化图表）**
- 最近7天收入柱状图（简单 CSS 实现）

---

### 4.26 接单管理页（pages/workspace/orders）

**① NavBar**："接单管理"

**② 状态 Tab**
- 待接单 | 进行中 | 已完成
- 对应 order_status: 1 | 2 | 3

**③ 订单列表**
- OrderCard（陪玩师视角）：
  - 用户头像+昵称 + 下单时间
  - 商品名 + 数量 + 金额
  - 操作按钮：
    - 待接单："接单"(紫色实心) → `POST /api/playapi/order/accept`
    - 进行中："完成服务"(绿色实心) → `POST /api/playapi/order/finish`
    - 已完成：显示评分和收入
- 下拉刷新 + 触底加载
- 接口：`GET /api/playapi/coach/orders?status=xxx&pageNum=1`

---

### 4.27 收入统计页（pages/workspace/income）

**① NavBar**："收入统计"

**② 收入概览卡片（紫色渐变）**
- 可提现余额(大字) + "去提现"按钮
- 累计收入 | 本月收入 | 今日收入

**③ 时间筛选**
- Tab：近7天 | 近30天 | 全部

**④ 收入趋势图**
- 折线图（简单 CSS/Canvas 实现）
- X轴：日期，Y轴：收入金额

**⑤ 收入明细列表**
- 每条：订单编号 + 商品名 + 实付金额 + 我的收入(绿色) + 时间
- 展开详情：平台抽成 xx% (¥xx) + 店铺抽成 xx% (¥xx) + 我的收入 ¥xx
- 下拉刷新 + 触底加载

---

### 4.28 商品管理页（pages/workspace/goods）

**① NavBar**："商品管理" + 右上角"发布商品"按钮

**② 商品列表**
- 每项：封面小图(60x60) + 商品名 + 价格 + 销量 + 状态标签(上架/下架)
- 操作：编辑 | 上架/下架切换
- 接口：`GET /api/playapi/coach/my_goods`

**③ 发布/编辑商品弹窗（半屏弹窗）**
- 商品名称：输入框
- 分类选择：Picker 选择器
- 封面图：图片上传
- 单价：输入框（元，提交时转为分）
- 计量单位：输入框（如"局"、"小时"）
- 商品描述：多行文本
- 提交按钮
- 接口：`POST /api/playapi/coach/goods/create` 或 `PUT /api/playapi/coach/goods/update`

---

## 五、页面路由汇总

| 序号 | 页面 | 路径 | 类型 | 需登录 |
|------|------|------|------|--------|
| 1 | 首页 | /pages/index/index | TabBar | 否 |
| 2 | 分类 | /pages/category/index | TabBar | 否 |
| 3 | 搜索 | /pages/search/index | 普通 | 否 |
| 4 | 陪玩师列表 | /pages/coach/list | 普通 | 否 |
| 5 | 陪玩师详情 | /pages/coach/detail | 普通 | 否 |
| 6 | 商品详情 | /pages/goods/detail | 普通 | 否 |
| 7 | 下单确认 | /pages/order/confirm | 普通 | 是 |
| 8 | 支付 | /pages/order/pay | 普通 | 是 |
| 9 | 订单列表 | /pages/order/list | 普通 | 是 |
| 10 | 订单详情 | /pages/order/detail | 普通 | 是 |
| 11 | 评价 | /pages/order/review | 普通 | 是 |
| 12 | 充值 | /pages/recharge/index | 普通 | 是 |
| 13 | 我的优惠券 | /pages/coupon/list | 普通 | 是 |
| 14 | 领券中心 | /pages/coupon/center | 普通 | 否 |
| 15 | 活动列表 | /pages/activity/list | TabBar | 否 |
| 16 | 活动详情 | /pages/activity/detail | 普通 | 否 |
| 17 | 消息 | /pages/message/index | TabBar | 是 |
| 18 | 我的 | /pages/mine/index | TabBar | 是 |
| 19 | 编辑资料 | /pages/mine/profile | 普通 | 是 |
| 20 | 余额明细 | /pages/mine/balance | 普通 | 是 |
| 21 | 设置 | /pages/mine/settings | 普通 | 是 |
| 22 | 登录 | /pages/login/index | 普通 | 否 |
| 23 | 陪玩师申请 | /pages/coach/apply | 普通 | 是 |
| 24 | 申请状态 | /pages/coach/apply-status | 普通 | 是 |
| 25 | 工作台 | /pages/workspace/index | 普通 | 是(Coach) |
| 26 | 接单管理 | /pages/workspace/orders | 普通 | 是(Coach) |
| 27 | 收入统计 | /pages/workspace/income | 普通 | 是(Coach) |
| 28 | 商品管理 | /pages/workspace/goods | 普通 | 是(Coach) |

---

## 六、接口调用汇总

| 页面 | 接口 | 方法 |
|------|------|------|
| 首页 | /api/playapi/activity/list | GET |
| 首页 | /api/playapi/coach/list | GET |
| 首页 | /api/playapi/goods/list | GET |
| 分类 | /api/playapi/category/list | GET |
| 分类 | /api/playapi/goods/list | GET |
| 搜索 | /api/playapi/search | GET |
| 陪玩师列表 | /api/playapi/coach/list | GET |
| 陪玩师详情 | /api/playapi/coach/detail | GET |
| 陪玩师详情 | /api/playapi/goods/list | GET |
| 陪玩师详情 | /api/playapi/review/list | GET |
| 商品详情 | /api/playapi/goods/detail | GET |
| 下单确认 | /api/playapi/coupon/usable | GET |
| 下单确认 | /api/playapi/order/create | POST |
| 支付 | /api/playapi/payment/pay | POST |
| 订单列表 | /api/playapi/order/list | GET |
| 订单详情 | /api/playapi/order/detail | GET |
| 取消订单 | /api/playapi/order/cancel | POST |
| 申请退款 | /api/playapi/order/refund | POST |
| 评价 | /api/playapi/review/create | POST |
| 充值 | /api/playapi/recharge/plans | GET |
| 充值 | /api/playapi/recharge/create | POST |
| 优惠券 | /api/playapi/coupon/my_list | GET |
| 领券中心 | /api/playapi/coupon/list | GET |
| 领券 | /api/playapi/coupon/claim | POST |
| 活动列表 | /api/playapi/activity/list | GET |
| 活动详情 | /api/playapi/activity/detail | GET |
| 参与活动 | /api/playapi/activity/join | POST |
| 完成步骤 | /api/playapi/activity/complete_step | POST |
| 领取奖励 | /api/playapi/activity/claim_reward | POST |
| 个人信息 | /api/playapi/member/info | GET |
| 编辑资料 | /api/playapi/member/update | PUT |
| 切换身份 | /api/playapi/member/switch_role | POST |
| 余额明细 | /api/playapi/member/balance_log | GET |
| 登录 | /api/playapi/auth/login | POST |
| 发验证码 | /api/playapi/auth/send_code | POST |
| 微信登录 | /api/playapi/auth/wx_login | POST |
| 陪玩师申请 | /api/playapi/coach/apply | POST |
| 申请状态 | /api/playapi/coach/apply_status | GET |
| 在线状态 | /api/playapi/coach/online | PUT |
| 接单 | /api/playapi/order/accept | POST |
| 完成服务 | /api/playapi/order/finish | POST |
| 我的接单 | /api/playapi/coach/orders | GET |
| 我的商品 | /api/playapi/coach/my_goods | GET |
| 发布商品 | /api/playapi/coach/goods/create | POST |
| 编辑商品 | /api/playapi/coach/goods/update | PUT |
| 上下架 | /api/playapi/coach/goods/status | PUT |
| 收入统计 | /api/playapi/coach/income | GET |

---

## 七、交互规范

### 7.1 加载状态
- 页面首次加载：骨架屏（Skeleton）
- 列表加载更多：底部 LoadMore 组件（加载中/没有更多了）
- 按钮操作：按钮 loading 态，防止重复点击

### 7.2 下拉刷新
- 所有列表页支持下拉刷新
- 刷新动画：紫色 loading 圆圈

### 7.3 空状态
- 列表为空时显示 EmptyState 组件
- 不同场景不同图标和文案

### 7.4 错误处理
- 网络错误：Toast 提示 "网络异常，请重试"
- 401 未登录：自动跳转登录页，登录后返回原页面
- 业务错误：Toast 提示具体错误信息

### 7.5 登录拦截
- 需登录页面：在 onLoad 中检查 Token，无 Token 跳转登录页
- 登录成功后通过 redirectUrl 参数返回原页面

### 7.6 金额显示
- 后端返回分，前端统一除以 100 显示元
- 使用 PriceText 组件统一格式化
- 保留 2 位小数
