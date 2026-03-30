package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// ActivityStep DTO 模型

// ActivityStepCreateInput 创建活动步骤表输入
type ActivityStepCreateInput struct {
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	StepNum int `json:"stepNum"`
	Title string `json:"title"`
	DescContent string `json:"descContent"`
	StepImage string `json:"stepImage"`
	Sort int `json:"sort"`
}

// ActivityStepUpdateInput 更新活动步骤表输入
type ActivityStepUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	StepNum int `json:"stepNum"`
	Title string `json:"title"`
	DescContent string `json:"descContent"`
	StepImage string `json:"stepImage"`
	Sort int `json:"sort"`
}

// ActivityStepDetailOutput 活动步骤表详情输出
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

// ActivityStepListOutput 活动步骤表列表输出
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

// ActivityStepListInput 活动步骤表列表查询输入
type ActivityStepListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

