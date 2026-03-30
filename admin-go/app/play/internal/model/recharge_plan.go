package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// RechargePlan DTO 模型

// RechargePlanCreateInput 创建å……å€¼æ–¹æ¡ˆè¡¨输入
type RechargePlanCreateInput struct {
	Title string `json:"title"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// RechargePlanUpdateInput 更新å……å€¼æ–¹æ¡ˆè¡¨输入
type RechargePlanUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Amount int64 `json:"amount"`
	GiftAmount int64 `json:"giftAmount"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// RechargePlanDetailOutput å……å€¼æ–¹æ¡ˆè¡¨详情输出
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

// RechargePlanListOutput å……å€¼æ–¹æ¡ˆè¡¨列表输出
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

// RechargePlanListInput å……å€¼æ–¹æ¡ˆè¡¨列表查询输入
type RechargePlanListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Status int `json:"status"`
}

