package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Payment DTO 模型

// PaymentCreateInput 创建æ”¯ä»˜è®°å½•è¡¨输入
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

// PaymentUpdateInput 更新æ”¯ä»˜è®°å½•è¡¨输入
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

// PaymentDetailOutput æ”¯ä»˜è®°å½•è¡¨详情输出
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

// PaymentListOutput æ”¯ä»˜è®°å½•è¡¨列表输出
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

// PaymentListInput æ”¯ä»˜è®°å½•è¡¨列表查询输入
type PaymentListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	PayType int `json:"payType"`
	PayStatus int `json:"payStatus"`
}

