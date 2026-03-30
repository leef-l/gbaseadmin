package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// Payment API

// PaymentCreateReq 创建æ”¯ä»˜è®°å½•è¡¨请求
type PaymentCreateReq struct {
	g.Meta `path:"/payment/create" method:"post" tags:"æ”¯ä»˜è®°å½•è¡¨" summary:"创建æ”¯ä»˜è®°å½•è¡¨"`
	OrderID snowflake.JsonInt64 `json:"orderID" v:"required#è®¢å•ID不能为空" dc:"è®¢å•ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#ä¼šå‘˜ID不能为空" dc:"ä¼šå‘˜ID"`
	PaymentNo string `json:"paymentNo" v:"required#æ”¯ä»˜æµæ°´å·不能为空" dc:"æ”¯ä»˜æµæ°´å·"`
	TradeNo string `json:"tradeNo"  dc:"ç¬¬ä¸‰æ–¹äº¤æ˜“å·"`
	PayType int `json:"payType"  dc:"æ”¯ä»˜æ–¹å¼"`
	PayAmount int64 `json:"payAmount"  dc:"æ”¯ä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	PayStatus int `json:"payStatus"  dc:"æ”¯ä»˜çŠ¶æ€"`
	PayAt *gtime.Time `json:"payAt"  dc:"æ”¯ä»˜æˆåŠŸæ—¶é—´"`
	RefundAt *gtime.Time `json:"refundAt"  dc:"é€€æ¬¾æ—¶é—´"`
	RefundAmount int64 `json:"refundAmount"  dc:"é€€æ¬¾é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CallbackContent string `json:"callbackContent"  dc:"å›žè°ƒæŠ¥æ–‡"`
}

// PaymentCreateRes 创建æ”¯ä»˜è®°å½•è¡¨响应
type PaymentCreateRes struct {
	g.Meta `mime:"application/json"`
}

// PaymentUpdateReq 更新æ”¯ä»˜è®°å½•è¡¨请求
type PaymentUpdateReq struct {
	g.Meta `path:"/payment/update" method:"put" tags:"æ”¯ä»˜è®°å½•è¡¨" summary:"更新æ”¯ä»˜è®°å½•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ”¯ä»˜è®°å½•è¡¨ID"`
	OrderID snowflake.JsonInt64 `json:"orderID" dc:"è®¢å•ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"ä¼šå‘˜ID"`
	PaymentNo string `json:"paymentNo" dc:"æ”¯ä»˜æµæ°´å·"`
	TradeNo string `json:"tradeNo" dc:"ç¬¬ä¸‰æ–¹äº¤æ˜“å·"`
	PayType int `json:"payType" dc:"æ”¯ä»˜æ–¹å¼"`
	PayAmount int64 `json:"payAmount" dc:"æ”¯ä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	PayStatus int `json:"payStatus" dc:"æ”¯ä»˜çŠ¶æ€"`
	PayAt *gtime.Time `json:"payAt" dc:"æ”¯ä»˜æˆåŠŸæ—¶é—´"`
	RefundAt *gtime.Time `json:"refundAt" dc:"é€€æ¬¾æ—¶é—´"`
	RefundAmount int64 `json:"refundAmount" dc:"é€€æ¬¾é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CallbackContent string `json:"callbackContent" dc:"å›žè°ƒæŠ¥æ–‡"`
}

// PaymentUpdateRes 更新æ”¯ä»˜è®°å½•è¡¨响应
type PaymentUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// PaymentDeleteReq 删除æ”¯ä»˜è®°å½•è¡¨请求
type PaymentDeleteReq struct {
	g.Meta `path:"/payment/delete" method:"delete" tags:"æ”¯ä»˜è®°å½•è¡¨" summary:"删除æ”¯ä»˜è®°å½•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ”¯ä»˜è®°å½•è¡¨ID"`
}

// PaymentDeleteRes 删除æ”¯ä»˜è®°å½•è¡¨响应
type PaymentDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// PaymentDetailReq 获取æ”¯ä»˜è®°å½•è¡¨详情请求
type PaymentDetailReq struct {
	g.Meta `path:"/payment/detail" method:"get" tags:"æ”¯ä»˜è®°å½•è¡¨" summary:"获取æ”¯ä»˜è®°å½•è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ”¯ä»˜è®°å½•è¡¨ID"`
}

// PaymentDetailRes 获取æ”¯ä»˜è®°å½•è¡¨详情响应
type PaymentDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.PaymentDetailOutput
}

// PaymentListReq 获取æ”¯ä»˜è®°å½•è¡¨列表请求
type PaymentListReq struct {
	g.Meta   `path:"/payment/list" method:"get" tags:"æ”¯ä»˜è®°å½•è¡¨" summary:"获取æ”¯ä»˜è®°å½•è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	PayType int `json:"payType" dc:"æ”¯ä»˜æ–¹å¼"`
	PayStatus int `json:"payStatus" dc:"æ”¯ä»˜çŠ¶æ€"`
}

// PaymentListRes 获取æ”¯ä»˜è®°å½•è¡¨列表响应
type PaymentListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.PaymentListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

