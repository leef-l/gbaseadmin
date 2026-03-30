package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Order DTO 模型

// OrderCreateInput 创建è®¢å•è¡¨输入
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

// OrderUpdateInput 更新è®¢å•è¡¨输入
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

// OrderDetailOutput è®¢å•è¡¨详情输出
type OrderDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	ShopTitle string `json:"shopTitle"`
	GoodsID snowflake.JsonInt64 `json:"goodsID"`
	GoodsTitle string `json:"goodsTitle"`
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

// OrderListOutput è®¢å•è¡¨列表输出
type OrderListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderNo string `json:"orderNo"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	ShopTitle string `json:"shopTitle"`
	GoodsID snowflake.JsonInt64 `json:"goodsID"`
	GoodsTitle string `json:"goodsTitle"`
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

// OrderListInput è®¢å•è¡¨列表查询输入
type OrderListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	PayType int `json:"payType"`
	OrderStatus int `json:"orderStatus"`
}

