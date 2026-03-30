package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// RechargeOrder DTO 模型

// RechargeOrderCreateInput 创建充值订单表输入
type RechargeOrderCreateInput struct {
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	RechargePlanID snowflake.JsonInt64 `json:"rechargePlanID"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	PayType int `json:"payType"`
	TradeNo string `json:"tradeNo"`
	PayStatus int `json:"payStatus"`
	PayAt *gtime.Time `json:"payAt"`
}

// RechargeOrderUpdateInput 更新充值订单表输入
type RechargeOrderUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	RechargePlanID snowflake.JsonInt64 `json:"rechargePlanID"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	PayType int `json:"payType"`
	TradeNo string `json:"tradeNo"`
	PayStatus int `json:"payStatus"`
	PayAt *gtime.Time `json:"payAt"`
}

// RechargeOrderDetailOutput 充值订单表详情输出
type RechargeOrderDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	RechargePlanID snowflake.JsonInt64 `json:"rechargePlanID"`
	RechargePlanTitle string `json:"rechargePlanTitle"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	PayType int `json:"payType"`
	TradeNo string `json:"tradeNo"`
	PayStatus int `json:"payStatus"`
	PayAt *gtime.Time `json:"payAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// RechargeOrderListOutput 充值订单表列表输出
type RechargeOrderListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	RechargePlanID snowflake.JsonInt64 `json:"rechargePlanID"`
	RechargePlanTitle string `json:"rechargePlanTitle"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	PayType int `json:"payType"`
	TradeNo string `json:"tradeNo"`
	PayStatus int `json:"payStatus"`
	PayAt *gtime.Time `json:"payAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// RechargeOrderListInput 充值订单表列表查询输入
type RechargeOrderListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	PayType int `json:"payType"`
	PayStatus int `json:"payStatus"`
}

