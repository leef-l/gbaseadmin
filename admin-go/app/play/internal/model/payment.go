package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Payment DTO 模型

// PaymentCreateInput 创建支付记录表输入
type PaymentCreateInput struct {
	OrderID snowflake.JsonInt64 `json:"orderID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	PaymentNo string `json:"paymentNo"`
	TradeNo string `json:"tradeNo"`
	PayType int `json:"payType"`
	PayAmount int64 `json:"payAmount"`
	PayStatus int `json:"payStatus"`
	PayAt *gtime.Time `json:"payAt"`
	RefundAt *gtime.Time `json:"refundAt"`
	RefundAmount int64 `json:"refundAmount"`
	CallbackContent string `json:"callbackContent"`
}

// PaymentUpdateInput 更新支付记录表输入
type PaymentUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	PaymentNo string `json:"paymentNo"`
	TradeNo string `json:"tradeNo"`
	PayType int `json:"payType"`
	PayAmount int64 `json:"payAmount"`
	PayStatus int `json:"payStatus"`
	PayAt *gtime.Time `json:"payAt"`
	RefundAt *gtime.Time `json:"refundAt"`
	RefundAmount int64 `json:"refundAmount"`
	CallbackContent string `json:"callbackContent"`
}

// PaymentDetailOutput 支付记录表详情输出
type PaymentDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	PaymentNo string `json:"paymentNo"`
	TradeNo string `json:"tradeNo"`
	PayType int `json:"payType"`
	PayAmount int64 `json:"payAmount"`
	PayStatus int `json:"payStatus"`
	PayAt *gtime.Time `json:"payAt"`
	RefundAt *gtime.Time `json:"refundAt"`
	RefundAmount int64 `json:"refundAmount"`
	CallbackContent string `json:"callbackContent"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// PaymentListOutput 支付记录表列表输出
type PaymentListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	PaymentNo string `json:"paymentNo"`
	TradeNo string `json:"tradeNo"`
	PayType int `json:"payType"`
	PayAmount int64 `json:"payAmount"`
	PayStatus int `json:"payStatus"`
	PayAt *gtime.Time `json:"payAt"`
	RefundAt *gtime.Time `json:"refundAt"`
	RefundAmount int64 `json:"refundAmount"`
	CallbackContent string `json:"callbackContent"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// PaymentListInput 支付记录表列表查询输入
type PaymentListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	PayType int `json:"payType"`
	PayStatus int `json:"payStatus"`
}

