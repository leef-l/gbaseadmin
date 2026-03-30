package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// Review API

// ReviewCreateReq 创建è¯„ä»·è¡¨请求
type ReviewCreateReq struct {
	g.Meta `path:"/review/create" method:"post" tags:"è¯„ä»·è¡¨" summary:"创建è¯„ä»·è¡¨"`
	OrderID snowflake.JsonInt64 `json:"orderID" v:"required#è®¢å•ID不能为空" dc:"è®¢å•ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#è¯„ä»·ä¼šå‘˜ID不能为空" dc:"è¯„ä»·ä¼šå‘˜ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" v:"required#è¢«è¯„é™ªçŽ©å¸ˆID不能为空" dc:"è¢«è¯„é™ªçŽ©å¸ˆID"`
	Score int `json:"score"  dc:"è¯„åˆ†ï¼ˆä¹˜100ï¼‰"`
	ReviewContent string `json:"reviewContent"  dc:"è¯„ä»·å†…å®¹"`
	ReviewImage string `json:"reviewImage"  dc:"è¯„ä»·å›¾ç‰‡ï¼ˆå¤šå¼ é€—å·åˆ†éš”ï¼‰"`
	ReplyContent string `json:"replyContent"  dc:"é™ªçŽ©å¸ˆå›žå¤å†…å®¹"`
	ReplyAt *gtime.Time `json:"replyAt"  dc:"å›žå¤æ—¶é—´"`
	IsAnonymous int `json:"isAnonymous"  dc:"æ˜¯å¦åŒ¿å"`
	Status int `json:"status"  dc:"çŠ¶æ€"`
}

// ReviewCreateRes 创建è¯„ä»·è¡¨响应
type ReviewCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ReviewUpdateReq 更新è¯„ä»·è¡¨请求
type ReviewUpdateReq struct {
	g.Meta `path:"/review/update" method:"put" tags:"è¯„ä»·è¡¨" summary:"更新è¯„ä»·è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"è¯„ä»·è¡¨ID"`
	OrderID snowflake.JsonInt64 `json:"orderID" dc:"è®¢å•ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"è¯„ä»·ä¼šå‘˜ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" dc:"è¢«è¯„é™ªçŽ©å¸ˆID"`
	Score int `json:"score" dc:"è¯„åˆ†ï¼ˆä¹˜100ï¼‰"`
	ReviewContent string `json:"reviewContent" dc:"è¯„ä»·å†…å®¹"`
	ReviewImage string `json:"reviewImage" dc:"è¯„ä»·å›¾ç‰‡ï¼ˆå¤šå¼ é€—å·åˆ†éš”ï¼‰"`
	ReplyContent string `json:"replyContent" dc:"é™ªçŽ©å¸ˆå›žå¤å†…å®¹"`
	ReplyAt *gtime.Time `json:"replyAt" dc:"å›žå¤æ—¶é—´"`
	IsAnonymous int `json:"isAnonymous" dc:"æ˜¯å¦åŒ¿å"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// ReviewUpdateRes 更新è¯„ä»·è¡¨响应
type ReviewUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ReviewDeleteReq 删除è¯„ä»·è¡¨请求
type ReviewDeleteReq struct {
	g.Meta `path:"/review/delete" method:"delete" tags:"è¯„ä»·è¡¨" summary:"删除è¯„ä»·è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"è¯„ä»·è¡¨ID"`
}

// ReviewDeleteRes 删除è¯„ä»·è¡¨响应
type ReviewDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ReviewDetailReq 获取è¯„ä»·è¡¨详情请求
type ReviewDetailReq struct {
	g.Meta `path:"/review/detail" method:"get" tags:"è¯„ä»·è¡¨" summary:"获取è¯„ä»·è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"è¯„ä»·è¡¨ID"`
}

// ReviewDetailRes 获取è¯„ä»·è¡¨详情响应
type ReviewDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ReviewDetailOutput
}

// ReviewListReq 获取è¯„ä»·è¡¨列表请求
type ReviewListReq struct {
	g.Meta   `path:"/review/list" method:"get" tags:"è¯„ä»·è¡¨" summary:"获取è¯„ä»·è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	IsAnonymous int `json:"isAnonymous" dc:"æ˜¯å¦åŒ¿å"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// ReviewListRes 获取è¯„ä»·è¡¨列表响应
type ReviewListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ReviewListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

