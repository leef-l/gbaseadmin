package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Review API

// ReviewCreateReq 创建评价表请求
type ReviewCreateReq struct {
	g.Meta `path:"/review/create" method:"post" tags:"评价表" summary:"创建评价表"`
	OrderID snowflake.JsonInt64 `json:"orderID" v:"required#订单ID不能为空" dc:"订单ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#评价会员ID不能为空" dc:"评价会员ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" v:"required#被评陪玩师ID不能为空" dc:"被评陪玩师ID"`
	Score int `json:"score"  dc:"评分（乘100，如 500=5.00分）"`
	ReviewContent string `json:"reviewContent"  dc:"评价内容"`
	ReviewImage string `json:"reviewImage"  dc:"评价图片（多张逗号分隔）"`
	ReplyContent string `json:"replyContent"  dc:"陪玩师回复内容"`
	ReplyAt *gtime.Time `json:"replyAt"  dc:"回复时间"`
	IsAnonymous int `json:"isAnonymous"  dc:"是否匿名"`
	Status int `json:"status"  dc:"状态"`
}

// ReviewCreateRes 创建评价表响应
type ReviewCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ReviewUpdateReq 更新评价表请求
type ReviewUpdateReq struct {
	g.Meta `path:"/review/update" method:"put" tags:"评价表" summary:"更新评价表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"评价表ID"`
	OrderID snowflake.JsonInt64 `json:"orderID" dc:"订单ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"评价会员ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" dc:"被评陪玩师ID"`
	Score int `json:"score" dc:"评分（乘100，如 500=5.00分）"`
	ReviewContent string `json:"reviewContent" dc:"评价内容"`
	ReviewImage string `json:"reviewImage" dc:"评价图片（多张逗号分隔）"`
	ReplyContent string `json:"replyContent" dc:"陪玩师回复内容"`
	ReplyAt *gtime.Time `json:"replyAt" dc:"回复时间"`
	IsAnonymous int `json:"isAnonymous" dc:"是否匿名"`
	Status int `json:"status" dc:"状态"`
}

// ReviewUpdateRes 更新评价表响应
type ReviewUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ReviewDeleteReq 删除评价表请求
type ReviewDeleteReq struct {
	g.Meta `path:"/review/delete" method:"delete" tags:"评价表" summary:"删除评价表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"评价表ID"`
}

// ReviewDeleteRes 删除评价表响应
type ReviewDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ReviewDetailReq 获取评价表详情请求
type ReviewDetailReq struct {
	g.Meta `path:"/review/detail" method:"get" tags:"评价表" summary:"获取评价表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"评价表ID"`
}

// ReviewDetailRes 获取评价表详情响应
type ReviewDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ReviewDetailOutput
}

// ReviewListReq 获取评价表列表请求
type ReviewListReq struct {
	g.Meta   `path:"/review/list" method:"get" tags:"评价表" summary:"获取评价表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	IsAnonymous int `json:"isAnonymous" dc:"是否匿名"`
	Status int `json:"status" dc:"状态"`
}

// ReviewListRes 获取评价表列表响应
type ReviewListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ReviewListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

