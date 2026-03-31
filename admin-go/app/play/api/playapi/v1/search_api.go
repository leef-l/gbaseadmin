package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 综合搜索（公开） ====================

type SearchReq struct {
	g.Meta   `path:"/search" method:"get" tags:"C端搜索" summary:"综合搜索"`
	Keyword  string `json:"keyword" v:"required|min-length:1|max-length:50#关键词不能为空|关键词至少1个字符|关键词最多50个字符" dc:"搜索关键词"`
	Type     string `json:"type" v:"in:all,coach,goods#搜索类型不合法" dc:"搜索类型:all=全部,coach=陪玩师,goods=商品" d:"all"`
	Page     int    `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type SearchRes struct {
	g.Meta     `mime:"application/json"`
	CoachTotal int               `json:"coachTotal" dc:"陪玩师匹配总数"`
	CoachList  []SearchCoachItem `json:"coachList" dc:"陪玩师结果列表"`
	GoodsTotal int               `json:"goodsTotal" dc:"商品匹配总数"`
	GoodsList  []SearchGoodsItem `json:"goodsList" dc:"商品结果列表"`
}

type SearchCoachItem struct {
	CoachID    string  `json:"coachId" dc:"陪玩师ID"`
	Nickname   string  `json:"nickname" dc:"昵称"`
	Avatar     string  `json:"avatar" dc:"头像"`
	IsOnline   int     `json:"isOnline" dc:"在线状态"`
	Score      float64 `json:"score" dc:"综合评分"`
	OrderCount int     `json:"orderCount" dc:"接单量"`
}

type SearchGoodsItem struct {
	GoodsID      string `json:"goodsId" dc:"商品ID"`
	Title        string `json:"title" dc:"商品标题"`
	CoverImage   string `json:"coverImage" dc:"封面图"`
	Price        int64  `json:"price" dc:"单价(分)"`
	Unit         string `json:"unit" dc:"单位"`
	CoachID      string `json:"coachId" dc:"陪玩师ID"`
	CoachName    string `json:"coachName" dc:"陪玩师昵称"`
}
