package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// MemberLevel DTO 模型

// MemberLevelCreateInput 创建会员等级表输入
type MemberLevelCreateInput struct {
	Title string `json:"title"`
	Level int `json:"level"`
	Icon string `json:"icon"`
	MinExp int `json:"minExp"`
	Discount int `json:"discount"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// MemberLevelUpdateInput 更新会员等级表输入
type MemberLevelUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Level int `json:"level"`
	Icon string `json:"icon"`
	MinExp int `json:"minExp"`
	Discount int `json:"discount"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// MemberLevelDetailOutput 会员等级表详情输出
type MemberLevelDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Level int `json:"level"`
	Icon string `json:"icon"`
	MinExp int `json:"minExp"`
	Discount int `json:"discount"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// MemberLevelListOutput 会员等级表列表输出
type MemberLevelListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Level int `json:"level"`
	Icon string `json:"icon"`
	MinExp int `json:"minExp"`
	Discount int `json:"discount"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// MemberLevelListInput 会员等级表列表查询输入
type MemberLevelListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Level int `json:"level"`
	Status int `json:"status"`
}

