package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// ActivityStep DTO 模型

// ActivityStepCreateInput 创建æ´»åŠ¨æ­¥éª¤è¡¨输入
type ActivityStepCreateInput struct {
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	StepNum int `json:"stepNum"`
	Title string `json:"title"`
	DescContent string `json:"descContent"`
	StepImage string `json:"stepImage"`
	Sort int `json:"sort"`
}

// ActivityStepUpdateInput 更新æ´»åŠ¨æ­¥éª¤è¡¨输入
type ActivityStepUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	StepNum int `json:"stepNum"`
	Title string `json:"title"`
	DescContent string `json:"descContent"`
	StepImage string `json:"stepImage"`
	Sort int `json:"sort"`
}

// ActivityStepDetailOutput æ´»åŠ¨æ­¥éª¤è¡¨详情输出
type ActivityStepDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	ActivityTitle string `json:"activityTitle"`
	StepNum int `json:"stepNum"`
	Title string `json:"title"`
	DescContent string `json:"descContent"`
	StepImage string `json:"stepImage"`
	Sort int `json:"sort"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityStepListOutput æ´»åŠ¨æ­¥éª¤è¡¨列表输出
type ActivityStepListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	ActivityTitle string `json:"activityTitle"`
	StepNum int `json:"stepNum"`
	Title string `json:"title"`
	DescContent string `json:"descContent"`
	StepImage string `json:"stepImage"`
	Sort int `json:"sort"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityStepListInput æ´»åŠ¨æ­¥éª¤è¡¨列表查询输入
type ActivityStepListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

