package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// Order API

// OrderCreateReq 创建è®¢å•è¡¨请求
type OrderCreateReq struct {
	g.Meta `path:"/order/create" method:"post" tags:"è®¢å•è¡¨" summary:"创建è®¢å•è¡¨"`
	OrderNo string `json:"orderNo" v:"required#è®¢å•ç¼–å·不能为空" dc:"è®¢å•ç¼–å·"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#ä¸‹å•ä¼šå‘˜ID不能为空" dc:"ä¸‹å•ä¼šå‘˜ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" v:"required#é™ªçŽ©å¸ˆID不能为空" dc:"é™ªçŽ©å¸ˆID"`
	ShopID snowflake.JsonInt64 `json:"shopID"  dc:"åº—é“ºID"`
	GoodsID snowflake.JsonInt64 `json:"goodsID" v:"required#å•†å“ID不能为空" dc:"å•†å“ID"`
	GoodsTitle string `json:"goodsTitle" v:"required#å•†å“åç§°ï¼ˆå†—ä½™ï¼‰不能为空" dc:"å•†å“åç§°ï¼ˆå†—ä½™ï¼‰"`
	GoodsPrice int64 `json:"goodsPrice" v:"required#å•†å“å•ä»·ï¼ˆåˆ†ï¼‰不能为空" dc:"å•†å“å•ä»·ï¼ˆåˆ†ï¼‰"`
	Quantity int `json:"quantity"  dc:"æ•°é‡"`
	TotalAmount int64 `json:"totalAmount"  dc:"è®¢å•æ€»é¢ï¼ˆåˆ†ï¼‰"`
	DiscountAmount int64 `json:"discountAmount"  dc:"ä¼šå‘˜æŠ˜æ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CouponAmount int64 `json:"couponAmount"  dc:"ä¼˜æƒ åˆ¸æŠµæ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	PayAmount int64 `json:"payAmount"  dc:"å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CouponMemberID snowflake.JsonInt64 `json:"couponMemberID"  dc:"ä½¿ç”¨çš„ä¼˜æƒ åˆ¸é¢†å–è®°å½•ID"`
	PayType int `json:"payType"  dc:"æ”¯ä»˜æ–¹å¼"`
	OrderStatus int `json:"orderStatus"  dc:"è®¢å•çŠ¶æ€"`
	PayAt *gtime.Time `json:"payAt"  dc:"æ”¯ä»˜æ—¶é—´"`
	StartAt *gtime.Time `json:"startAt"  dc:"æœåŠ¡å¼€å§‹æ—¶é—´"`
	FinishAt *gtime.Time `json:"finishAt"  dc:"æœåŠ¡å®Œæˆæ—¶é—´"`
	CancelAt *gtime.Time `json:"cancelAt"  dc:"å–æ¶ˆæ—¶é—´"`
	CancelReason string `json:"cancelReason"  dc:"å–æ¶ˆåŽŸå›"`
	Remark string `json:"remark"  dc:"è®¢å•å¤‡æ³¨"`
}

// OrderCreateRes 创建è®¢å•è¡¨响应
type OrderCreateRes struct {
	g.Meta `mime:"application/json"`
}

// OrderUpdateReq 更新è®¢å•è¡¨请求
type OrderUpdateReq struct {
	g.Meta `path:"/order/update" method:"put" tags:"è®¢å•è¡¨" summary:"更新è®¢å•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"è®¢å•è¡¨ID"`
	OrderNo string `json:"orderNo" dc:"è®¢å•ç¼–å·"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"ä¸‹å•ä¼šå‘˜ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" dc:"é™ªçŽ©å¸ˆID"`
	ShopID snowflake.JsonInt64 `json:"shopID" dc:"åº—é“ºID"`
	GoodsID snowflake.JsonInt64 `json:"goodsID" dc:"å•†å“ID"`
	GoodsTitle string `json:"goodsTitle" dc:"å•†å“åç§°ï¼ˆå†—ä½™ï¼‰"`
	GoodsPrice int64 `json:"goodsPrice" dc:"å•†å“å•ä»·ï¼ˆåˆ†ï¼‰"`
	Quantity int `json:"quantity" dc:"æ•°é‡"`
	TotalAmount int64 `json:"totalAmount" dc:"è®¢å•æ€»é¢ï¼ˆåˆ†ï¼‰"`
	DiscountAmount int64 `json:"discountAmount" dc:"ä¼šå‘˜æŠ˜æ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CouponAmount int64 `json:"couponAmount" dc:"ä¼˜æƒ åˆ¸æŠµæ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	PayAmount int64 `json:"payAmount" dc:"å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CouponMemberID snowflake.JsonInt64 `json:"couponMemberID" dc:"ä½¿ç”¨çš„ä¼˜æƒ åˆ¸é¢†å–è®°å½•ID"`
	PayType int `json:"payType" dc:"æ”¯ä»˜æ–¹å¼"`
	OrderStatus int `json:"orderStatus" dc:"è®¢å•çŠ¶æ€"`
	PayAt *gtime.Time `json:"payAt" dc:"æ”¯ä»˜æ—¶é—´"`
	StartAt *gtime.Time `json:"startAt" dc:"æœåŠ¡å¼€å§‹æ—¶é—´"`
	FinishAt *gtime.Time `json:"finishAt" dc:"æœåŠ¡å®Œæˆæ—¶é—´"`
	CancelAt *gtime.Time `json:"cancelAt" dc:"å–æ¶ˆæ—¶é—´"`
	CancelReason string `json:"cancelReason" dc:"å–æ¶ˆåŽŸå›"`
	Remark string `json:"remark" dc:"è®¢å•å¤‡æ³¨"`
}

// OrderUpdateRes 更新è®¢å•è¡¨响应
type OrderUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// OrderDeleteReq 删除è®¢å•è¡¨请求
type OrderDeleteReq struct {
	g.Meta `path:"/order/delete" method:"delete" tags:"è®¢å•è¡¨" summary:"删除è®¢å•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"è®¢å•è¡¨ID"`
}

// OrderDeleteRes 删除è®¢å•è¡¨响应
type OrderDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// OrderDetailReq 获取è®¢å•è¡¨详情请求
type OrderDetailReq struct {
	g.Meta `path:"/order/detail" method:"get" tags:"è®¢å•è¡¨" summary:"获取è®¢å•è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"è®¢å•è¡¨ID"`
}

// OrderDetailRes 获取è®¢å•è¡¨详情响应
type OrderDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.OrderDetailOutput
}

// OrderListReq 获取è®¢å•è¡¨列表请求
type OrderListReq struct {
	g.Meta   `path:"/order/list" method:"get" tags:"è®¢å•è¡¨" summary:"获取è®¢å•è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	PayType int `json:"payType" dc:"æ”¯ä»˜æ–¹å¼"`
	OrderStatus int `json:"orderStatus" dc:"è®¢å•çŠ¶æ€"`
}

// OrderListRes 获取è®¢å•è¡¨列表响应
type OrderListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.OrderListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

