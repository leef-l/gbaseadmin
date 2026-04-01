package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// CouponMember DTO 模型

// CouponMemberCreateInput 创建会员优惠券表输入
type CouponMemberCreateInput struct {
	CouponID snowflake.JsonInt64 `json:"couponID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	UseStatus int `json:"useStatus"`
	ClaimAt *gtime.Time `json:"claimAt"`
	UseAt *gtime.Time `json:"useAt"`
	ExpireAt *gtime.Time `json:"expireAt"`
}

// CouponMemberUpdateInput 更新会员优惠券表输入
type CouponMemberUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	CouponID snowflake.JsonInt64 `json:"couponID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	UseStatus int `json:"useStatus"`
	ClaimAt *gtime.Time `json:"claimAt"`
	UseAt *gtime.Time `json:"useAt"`
	ExpireAt *gtime.Time `json:"expireAt"`
}

// CouponMemberDetailOutput 会员优惠券表详情输出
type CouponMemberDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	CouponID snowflake.JsonInt64 `json:"couponID"`
	CouponTitle string `json:"couponTitle"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	MemberNickname string `json:"memberNickname"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	UseStatus int `json:"useStatus"`
	ClaimAt *gtime.Time `json:"claimAt"`
	UseAt *gtime.Time `json:"useAt"`
	ExpireAt *gtime.Time `json:"expireAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CouponMemberListOutput 会员优惠券表列表输出
type CouponMemberListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	CouponID snowflake.JsonInt64 `json:"couponID"`
	CouponTitle string `json:"couponTitle"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	MemberNickname string `json:"memberNickname"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	UseStatus int `json:"useStatus"`
	ClaimAt *gtime.Time `json:"claimAt"`
	UseAt *gtime.Time `json:"useAt"`
	ExpireAt *gtime.Time `json:"expireAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CouponMemberListInput 会员优惠券表列表查询输入
type CouponMemberListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	UseStatus int `json:"useStatus"`
}

