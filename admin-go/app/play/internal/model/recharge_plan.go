package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// RechargePlan DTO 模型

// RechargePlanCreateInput 创建充值方案表输入
type RechargePlanCreateInput struct {
	Title string `json:"title"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// RechargePlanUpdateInput 更新充值方案表输入
type RechargePlanUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// RechargePlanDetailOutput 充值方案表详情输出
type RechargePlanDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// RechargePlanListOutput 充值方案表列表输出
type RechargePlanListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// RechargePlanListInput 充值方案表列表查询输入
type RechargePlanListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Status int `json:"status"`
}

