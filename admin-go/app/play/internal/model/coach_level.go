package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// CoachLevel DTO 模型

// CoachLevelCreateInput 创建陪玩师等级表输入
type CoachLevelCreateInput struct {
	Title string `json:"title"`
	Level int `json:"level"`
	Icon string `json:"icon"`
	MinOrders int `json:"minOrders"`
	MinScore int `json:"minScore"`
	CommissionRate int `json:"commissionRate"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// CoachLevelUpdateInput 更新陪玩师等级表输入
type CoachLevelUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Level int `json:"level"`
	Icon string `json:"icon"`
	MinOrders int `json:"minOrders"`
	MinScore int `json:"minScore"`
	CommissionRate int `json:"commissionRate"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// CoachLevelDetailOutput 陪玩师等级表详情输出
type CoachLevelDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Level int `json:"level"`
	Icon string `json:"icon"`
	MinOrders int `json:"minOrders"`
	MinScore int `json:"minScore"`
	CommissionRate int `json:"commissionRate"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CoachLevelListOutput 陪玩师等级表列表输出
type CoachLevelListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Level int `json:"level"`
	Icon string `json:"icon"`
	MinOrders int `json:"minOrders"`
	MinScore int `json:"minScore"`
	CommissionRate int `json:"commissionRate"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CoachLevelListInput 陪玩师等级表列表查询输入
type CoachLevelListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Level int `json:"level"`
	Status int `json:"status"`
}

