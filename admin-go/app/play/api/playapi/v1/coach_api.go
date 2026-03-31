package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ========== 公开接口 ==========

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

type CoachListItem struct {
	CoachID    string  `json:"coachId" dc:"陪玩师ID"`
	Nickname   string  `json:"nickname" dc:"昵称"`
	Avatar     string  `json:"avatar" dc:"头像"`
	Gender     int     `json:"gender" dc:"性别"`
	IsOnline   int     `json:"isOnline" dc:"在线状态:0=离线,1=在线"`
	Score      float64 `json:"score" dc:"综合评分"`
	OrderCount int     `json:"orderCount" dc:"接单量"`
	Intro      string  `json:"intro" dc:"个人简介"`
	MinPrice   int64   `json:"minPrice" dc:"最低商品价格(分)"`
}

type CoachListRes struct {
	g.Meta `mime:"application/json"`
	Total  int             `json:"total" dc:"总数"`
	List   []CoachListItem `json:"list" dc:"陪玩师列表"`
}

type CoachDetailReq struct {
	g.Meta  `path:"/coach/detail" method:"get" tags:"C端陪玩师" summary:"陪玩师详情"`
	CoachID string `json:"coachId" v:"required#陪玩师ID不能为空" dc:"陪玩师ID"`
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

type CoachDetailRes struct {
	g.Meta        `mime:"application/json"`
	CoachID       string            `json:"coachId" dc:"陪玩师ID"`
	Nickname      string            `json:"nickname" dc:"昵称"`
	Avatar        string            `json:"avatar" dc:"头像"`
	Gender        int               `json:"gender" dc:"性别"`
	IsOnline      int               `json:"isOnline" dc:"在线状态"`
	Score         float64           `json:"score" dc:"综合评分"`
	OrderCount    int               `json:"orderCount" dc:"接单量"`
	Intro         string            `json:"intro" dc:"个人简介"`
	GoodsList     []CoachGoodsItem  `json:"goodsList" dc:"商品列表"`
	ReviewCount   int               `json:"reviewCount" dc:"评价总数"`
	RecentReviews []ReviewBriefItem `json:"recentReviews" dc:"最近3条评价"`
}

// ========== MemberAuth 接口 ==========

type CoachApplyReq struct {
	g.Meta           `path:"/coach/apply" method:"post" tags:"C端陪玩师" summary:"申请成为陪玩师"`
	RealName         string `json:"realName" v:"required|length:2,20#真实姓名不能为空|姓名长度2-20个字符" dc:"真实姓名"`
	IdCard           string `json:"idCard" v:"required|resident-id#身份证号不能为空|身份证号格式不正确" dc:"身份证号"`
	IdCardFrontImage string `json:"idCardFrontImage" v:"required#身份证正面照不能为空" dc:"身份证正面照URL"`
	IdCardBackImage  string `json:"idCardBackImage" v:"required#身份证反面照不能为空" dc:"身份证反面照URL"`
	SkillDesc        string `json:"skillDesc" v:"required|max-length:500#技能描述不能为空|技能描述最多500字" dc:"技能描述"`
}

type CoachApplyRes struct {
	g.Meta `mime:"application/json"`
}

type CoachApplyStatusReq struct {
	g.Meta `path:"/coach/apply_status" method:"get" tags:"C端陪玩师" summary:"查询申请状态"`
}

type CoachApplyStatusRes struct {
	g.Meta      `mime:"application/json"`
	HasApply    bool   `json:"hasApply" dc:"是否有申请记录"`
	AuditStatus int    `json:"auditStatus" dc:"审核状态:0=待审核,1=通过,2=拒绝"`
	AuditRemark string `json:"auditRemark" dc:"审核备注"`
}

// ========== CoachOnly 接口 ==========

type CoachOnlineReq struct {
	g.Meta   `path:"/coach/online" method:"put" tags:"C端陪玩师" summary:"设置在线状态"`
	IsOnline int `json:"isOnline" v:"required|in:0,1#在线状态不能为空|在线状态值不合法" dc:"在线状态:0=离线,1=在线"`
}

type CoachOnlineRes struct {
	g.Meta `mime:"application/json"`
}

type CoachMyGoodsReq struct {
	g.Meta   `path:"/coach/my_goods" method:"get" tags:"C端陪玩师" summary:"我的商品列表"`
	Status   *int `json:"status" dc:"状态筛选:0=下架,1=上架"`
	Page     int  `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int  `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type CoachMyGoodsRes struct {
	g.Meta `mime:"application/json"`
	Total  int              `json:"total" dc:"总数"`
	List   []CoachGoodsItem `json:"list" dc:"商品列表"`
}

type CoachGoodsCreateReq struct {
	g.Meta      `path:"/coach/goods/create" method:"post" tags:"C端陪玩师" summary:"发布商品"`
	CategoryID  string `json:"categoryId" v:"required#分类ID不能为空" dc:"分类ID"`
	Title       string `json:"title" v:"required|max-length:50#商品标题不能为空|标题最多50字" dc:"商品标题"`
	Description string `json:"description" v:"max-length:500#描述最多500字" dc:"商品描述"`
	Price       int64  `json:"price" v:"required|min:1#价格不能为空|价格必须大于0" dc:"单价(分)"`
	Unit        string `json:"unit" v:"required|in:局,小时,次#单位不能为空|单位值不合法" dc:"单位"`
	CoverImage  string `json:"coverImage" dc:"商品封面图URL"`
	Sort        int    `json:"sort" dc:"排序（升序）"`
}

type CoachGoodsCreateRes struct {
	g.Meta `mime:"application/json"`
}

type CoachGoodsUpdateReq struct {
	g.Meta      `path:"/coach/goods/update" method:"put" tags:"C端陪玩师" summary:"编辑商品"`
	GoodsID     string `json:"goodsId" v:"required#商品ID不能为空" dc:"商品ID"`
	CategoryID  string `json:"categoryId" dc:"分类ID"`
	Title       string `json:"title" v:"max-length:50#标题最多50字" dc:"商品标题"`
	Description string `json:"description" v:"max-length:500#描述最多500字" dc:"商品描述"`
	Price       *int64 `json:"price" v:"min:1#价格必须大于0" dc:"单价(分)"`
	Unit        string `json:"unit" v:"in:局,小时,次#单位值不合法" dc:"单位"`
	CoverImage  string `json:"coverImage" dc:"商品封面图URL"`
	Sort        *int   `json:"sort" dc:"排序"`
}

type CoachGoodsUpdateRes struct {
	g.Meta `mime:"application/json"`
}

type CoachGoodsStatusReq struct {
	g.Meta  `path:"/coach/goods/status" method:"put" tags:"C端陪玩师" summary:"上下架商品"`
	GoodsID string `json:"goodsId" v:"required#商品ID不能为空" dc:"商品ID"`
	Status  int    `json:"status" v:"required|in:0,1#状态不能为空|状态值不合法" dc:"状态:0=下架,1=上架"`
}

type CoachGoodsStatusRes struct {
	g.Meta `mime:"application/json"`
}

type CoachIncomeReq struct {
	g.Meta `path:"/coach/income" method:"get" tags:"C端陪玩师" summary:"收入统计"`
}

type CoachIncomeRes struct {
	g.Meta      `mime:"application/json"`
	TodayIncome int64 `json:"todayIncome" dc:"今日收入(分)"`
	WeekIncome  int64 `json:"weekIncome" dc:"本周收入(分)"`
	MonthIncome int64 `json:"monthIncome" dc:"本月收入(分)"`
	TotalIncome int64 `json:"totalIncome" dc:"累计总收入(分)"`
	TodayOrders int   `json:"todayOrders" dc:"今日接单数"`
	TotalOrders int   `json:"totalOrders" dc:"累计接单数"`
}

type CoachOrdersReq struct {
	g.Meta   `path:"/coach/orders" method:"get" tags:"C端陪玩师" summary:"我的接单列表"`
	Status   *int `json:"status" dc:"状态筛选"`
	Page     int  `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int  `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type CoachOrdersRes struct {
	g.Meta `mime:"application/json"`
	Total  int              `json:"total" dc:"总数"`
	List   []OrderListItem  `json:"list" dc:"订单列表"`
}
