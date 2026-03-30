package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ========== MemberAuth 接口 ==========

type ReviewCreateReq struct {
	g.Meta      `path:"/review/create" method:"post" tags:"C端评价" summary:"发表评价"`
	OrderID     string  `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
	Score       float64 `json:"score" v:"required|min:1|max:5#评分不能为空|评分最低1分|评分最高5分" dc:"评分(1-5)"`
	Content     string  `json:"content" v:"required|max-length:500#评价内容不能为空|评价内容最多500字" dc:"评价内容"`
	Images      string  `json:"images" dc:"评价图片(逗号分隔URL)"`
	IsAnonymous int     `json:"isAnonymous" v:"in:0,1#匿名值不合法" dc:"是否匿名:0=否,1=是" d:"0"`
}

type ReviewCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ========== 公开接口 ==========

type ReviewListReq struct {
	g.Meta   `path:"/review/list" method:"get" tags:"C端评价" summary:"评价列表"`
	CoachID  string `json:"coachId" v:"required#陪玩师ID不能为空" dc:"陪玩师ID"`
	Page     int    `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type ReviewListItem struct {
	ReviewID    string  `json:"reviewId" dc:"评价ID"`
	Nickname    string  `json:"nickname" dc:"评价者昵称"`
	Avatar      string  `json:"avatar" dc:"评价者头像"`
	Score       float64 `json:"score" dc:"评分"`
	Content     string  `json:"content" dc:"评价内容"`
	Images      string  `json:"images" dc:"评价图片"`
	IsAnonymous int     `json:"isAnonymous" dc:"是否匿名"`
	ReplyContent string `json:"replyContent" dc:"陪玩师回复"`
	CreatedAt   string  `json:"createdAt" dc:"评价时间"`
}

type ReviewListRes struct {
	g.Meta `mime:"application/json"`
	Total  int              `json:"total" dc:"总数"`
	List   []ReviewListItem `json:"list" dc:"评价列表"`
}

// ========== CoachOnly 接口 ==========

type ReviewReplyReq struct {
	g.Meta   `path:"/review/reply" method:"post" tags:"C端评价" summary:"陪玩师回复评价"`
	ReviewID string `json:"reviewId" v:"required#评价ID不能为空" dc:"评价ID"`
	Reply    string `json:"reply" v:"required|max-length:200#回复内容不能为空|回复内容最多200字" dc:"回复内容"`
}

type ReviewReplyRes struct {
	g.Meta `mime:"application/json"`
}
