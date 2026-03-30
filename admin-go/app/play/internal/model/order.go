package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Order DTO 模型

// OrderCreateInput 创建订单表输入
type OrderCreateInput struct {
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	GoodsID snowflake.JsonInt64 `json:"goodsID"`
	GoodsTitle string `json:"goodsTitle"`
	GoodsPrice int64 `json:"goodsPrice"`
	Quantity int `json:"quantity"`
	TotalAmount int64 `json:"totalAmount"`
	DiscountAmount int64 `json:"discountAmount"`
	CouponAmount int64 `json:"couponAmount"`
	PayAmount int64 `json:"payAmount"`
	CouponMemberID snowflake.JsonInt64 `json:"couponMemberID"`
	PayType int `json:"payType"`
	OrderStatus int `json:"orderStatus"`
	PayAt *gtime.Time `json:"payAt"`
	StartAt *gtime.Time `json:"startAt"`
	FinishAt *gtime.Time `json:"finishAt"`
	CancelAt *gtime.Time `json:"cancelAt"`
	CancelReason string `json:"cancelReason"`
	Remark string `json:"remark"`
}

// OrderUpdateInput 更新订单表输入
type OrderUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	GoodsID snowflake.JsonInt64 `json:"goodsID"`
	GoodsTitle string `json:"goodsTitle"`
	GoodsPrice int64 `json:"goodsPrice"`
	Quantity int `json:"quantity"`
	TotalAmount int64 `json:"totalAmount"`
	DiscountAmount int64 `json:"discountAmount"`
	CouponAmount int64 `json:"couponAmount"`
	PayAmount int64 `json:"payAmount"`
	CouponMemberID snowflake.JsonInt64 `json:"couponMemberID"`
	PayType int `json:"payType"`
	OrderStatus int `json:"orderStatus"`
	PayAt *gtime.Time `json:"payAt"`
	StartAt *gtime.Time `json:"startAt"`
	FinishAt *gtime.Time `json:"finishAt"`
	CancelAt *gtime.Time `json:"cancelAt"`
	CancelReason string `json:"cancelReason"`
	Remark string `json:"remark"`
}

// OrderDetailOutput 订单表详情输出
type OrderDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	ShopTitle string `json:"shopTitle"`
	GoodsID snowflake.JsonInt64 `json:"goodsID"`
	GoodsTitle string `json:"goodsTitle"`
	GoodsPrice int64 `json:"goodsPrice"`
	Quantity int `json:"quantity"`
	TotalAmount int64 `json:"totalAmount"`
	DiscountAmount int64 `json:"discountAmount"`
	CouponAmount int64 `json:"couponAmount"`
	PayAmount int64 `json:"payAmount"`
	CouponMemberID snowflake.JsonInt64 `json:"couponMemberID"`
	PayType int `json:"payType"`
	OrderStatus int `json:"orderStatus"`
	PayAt *gtime.Time `json:"payAt"`
	StartAt *gtime.Time `json:"startAt"`
	FinishAt *gtime.Time `json:"finishAt"`
	CancelAt *gtime.Time `json:"cancelAt"`
	CancelReason string `json:"cancelReason"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// OrderListOutput 订单表列表输出
type OrderListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	ShopTitle string `json:"shopTitle"`
	GoodsID snowflake.JsonInt64 `json:"goodsID"`
	GoodsTitle string `json:"goodsTitle"`
	GoodsPrice int64 `json:"goodsPrice"`
	Quantity int `json:"quantity"`
	TotalAmount int64 `json:"totalAmount"`
	DiscountAmount int64 `json:"discountAmount"`
	CouponAmount int64 `json:"couponAmount"`
	PayAmount int64 `json:"payAmount"`
	CouponMemberID snowflake.JsonInt64 `json:"couponMemberID"`
	PayType int `json:"payType"`
	OrderStatus int `json:"orderStatus"`
	PayAt *gtime.Time `json:"payAt"`
	StartAt *gtime.Time `json:"startAt"`
	FinishAt *gtime.Time `json:"finishAt"`
	CancelAt *gtime.Time `json:"cancelAt"`
	CancelReason string `json:"cancelReason"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// OrderListInput 订单表列表查询输入
type OrderListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	PayType int `json:"payType"`
	OrderStatus int `json:"orderStatus"`
}


