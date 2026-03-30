package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// RechargeOrder DTO 模型

// RechargeOrderCreateInput 创建å……å€¼è®¢å•è¡¨输入
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

// RechargeOrderUpdateInput 更新å……å€¼è®¢å•è¡¨输入
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

// RechargeOrderDetailOutput å……å€¼è®¢å•è¡¨详情输出
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

// RechargeOrderListOutput å……å€¼è®¢å•è¡¨列表输出
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

// RechargeOrderListInput å……å€¼è®¢å•è¡¨列表查询输入
type RechargeOrderListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	PayType int `json:"payType"`
	PayStatus int `json:"payStatus"`
}

