package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Coupon DTO 模型

// CouponCreateInput 创建优惠券模板表输入
type CouponCreateInput struct {
	Title string `json:"title"`
	Type int `json:"type"`
	IsNewMember int `json:"isNewMember"`
	FaceValue int64 `json:"faceValue"`
	MinAmount int64 `json:"minAmount"`
	TotalNum int `json:"totalNum"`
	UsedNum int `json:"usedNum"`
	ClaimNum int `json:"claimNum"`
	PerLimit int `json:"perLimit"`
	ValidStartAt *gtime.Time `json:"validStartAt"`
	ValidEndAt *gtime.Time `json:"validEndAt"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// CouponUpdateInput 更新优惠券模板表输入
type CouponUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Type int `json:"type"`
	IsNewMember int `json:"isNewMember"`
	FaceValue int64 `json:"faceValue"`
	MinAmount int64 `json:"minAmount"`
	TotalNum int `json:"totalNum"`
	UsedNum int `json:"usedNum"`
	ClaimNum int `json:"claimNum"`
	PerLimit int `json:"perLimit"`
	ValidStartAt *gtime.Time `json:"validStartAt"`
	ValidEndAt *gtime.Time `json:"validEndAt"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// CouponDetailOutput 优惠券模板表详情输出
type CouponDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Type int `json:"type"`
	IsNewMember int `json:"isNewMember"`
	FaceValue int64 `json:"faceValue"`
	MinAmount int64 `json:"minAmount"`
	TotalNum int `json:"totalNum"`
	UsedNum int `json:"usedNum"`
	ClaimNum int `json:"claimNum"`
	PerLimit int `json:"perLimit"`
	ValidStartAt *gtime.Time `json:"validStartAt"`
	ValidEndAt *gtime.Time `json:"validEndAt"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CouponListOutput 优惠券模板表列表输出
type CouponListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Type int `json:"type"`
	IsNewMember int `json:"isNewMember"`
	FaceValue int64 `json:"faceValue"`
	MinAmount int64 `json:"minAmount"`
	TotalNum int `json:"totalNum"`
	UsedNum int `json:"usedNum"`
	ClaimNum int `json:"claimNum"`
	PerLimit int `json:"perLimit"`
	ValidStartAt *gtime.Time `json:"validStartAt"`
	ValidEndAt *gtime.Time `json:"validEndAt"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CouponListInput 优惠券模板表列表查询输入
type CouponListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Type int `json:"type"`
	IsNewMember int `json:"isNewMember"`
	Status int `json:"status"`
}

