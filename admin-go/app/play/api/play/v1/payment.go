package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Payment API

// PaymentCreateReq 创建支付记录表请求
type PaymentCreateReq struct {
	g.Meta `path:"/payment/create" method:"post" tags:"支付记录表" summary:"创建支付记录表"`
	OrderID snowflake.JsonInt64 `json:"orderID" v:"required#订单ID不能为空" dc:"订单ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#会员ID不能为空" dc:"会员ID"`
	PaymentNo string `json:"paymentNo" v:"required#支付流水号（平台内部）不能为空" dc:"支付流水号（平台内部）"`
	TradeNo string `json:"tradeNo"  dc:"第三方交易号"`
	PayType int `json:"payType"  dc:"支付方式"`
	PayAmount int64 `json:"payAmount"  dc:"支付金额（分）"`
	PayStatus int `json:"payStatus"  dc:"支付状态"`
	PayAt *gtime.Time `json:"payAt"  dc:"支付成功时间"`
	RefundAt *gtime.Time `json:"refundAt"  dc:"退款时间"`
	RefundAmount int64 `json:"refundAmount"  dc:"退款金额（分）"`
	CallbackContent string `json:"callbackContent"  dc:"回调报文"`
}

// PaymentCreateRes 创建支付记录表响应
type PaymentCreateRes struct {
	g.Meta `mime:"application/json"`
}

// PaymentUpdateReq 更新支付记录表请求
type PaymentUpdateReq struct {
	g.Meta `path:"/payment/update" method:"put" tags:"支付记录表" summary:"更新支付记录表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"支付记录表ID"`
	OrderID snowflake.JsonInt64 `json:"orderID" dc:"订单ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"会员ID"`
	PaymentNo string `json:"paymentNo" dc:"支付流水号（平台内部）"`
	TradeNo string `json:"tradeNo" dc:"第三方交易号"`
	PayType int `json:"payType" dc:"支付方式"`
	PayAmount int64 `json:"payAmount" dc:"支付金额（分）"`
	PayStatus int `json:"payStatus" dc:"支付状态"`
	PayAt *gtime.Time `json:"payAt" dc:"支付成功时间"`
	RefundAt *gtime.Time `json:"refundAt" dc:"退款时间"`
	RefundAmount int64 `json:"refundAmount" dc:"退款金额（分）"`
	CallbackContent string `json:"callbackContent" dc:"回调报文"`
}

// PaymentUpdateRes 更新支付记录表响应
type PaymentUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// PaymentDeleteReq 删除支付记录表请求
type PaymentDeleteReq struct {
	g.Meta `path:"/payment/delete" method:"delete" tags:"支付记录表" summary:"删除支付记录表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"支付记录表ID"`
}

// PaymentDeleteRes 删除支付记录表响应
type PaymentDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// PaymentDetailReq 获取支付记录表详情请求
type PaymentDetailReq struct {
	g.Meta `path:"/payment/detail" method:"get" tags:"支付记录表" summary:"获取支付记录表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"支付记录表ID"`
}

// PaymentDetailRes 获取支付记录表详情响应
type PaymentDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.PaymentDetailOutput
}

// PaymentListReq 获取支付记录表列表请求
type PaymentListReq struct {
	g.Meta   `path:"/payment/list" method:"get" tags:"支付记录表" summary:"获取支付记录表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	PayType int `json:"payType" dc:"支付方式"`
	PayStatus int `json:"payStatus" dc:"支付状态"`
}

// PaymentListRes 获取支付记录表列表响应
type PaymentListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.PaymentListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

