package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Activity DTO 模型

// ActivityCreateInput 创建æ´»åŠ¨è¡¨输入
type ActivityCreateInput struct {
	Title string `json:"title"`
	CoverImage string `json:"coverImage"`
	DescContent string `json:"descContent"`
	Type int `json:"type"`
	ConditionType int `json:"conditionType"`
	ConditionValue int64 `json:"conditionValue"`
	IsAutoReward int `json:"isAutoReward"`
	StartAt *gtime.Time `json:"startAt"`
	EndAt *gtime.Time `json:"endAt"`
	MaxNum int `json:"maxNum"`
	JoinNum int `json:"joinNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// ActivityUpdateInput 更新æ´»åŠ¨è¡¨输入
type ActivityUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	CoverImage string `json:"coverImage"`
	DescContent string `json:"descContent"`
	Type int `json:"type"`
	ConditionType int `json:"conditionType"`
	ConditionValue int64 `json:"conditionValue"`
	IsAutoReward int `json:"isAutoReward"`
	StartAt *gtime.Time `json:"startAt"`
	EndAt *gtime.Time `json:"endAt"`
	MaxNum int `json:"maxNum"`
	JoinNum int `json:"joinNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// ActivityDetailOutput æ´»åŠ¨è¡¨详情输出
type ActivityDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	CoverImage string `json:"coverImage"`
	DescContent string `json:"descContent"`
	Type int `json:"type"`
	ConditionType int `json:"conditionType"`
	ConditionValue int64 `json:"conditionValue"`
	IsAutoReward int `json:"isAutoReward"`
	StartAt *gtime.Time `json:"startAt"`
	EndAt *gtime.Time `json:"endAt"`
	MaxNum int `json:"maxNum"`
	JoinNum int `json:"joinNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityListOutput æ´»åŠ¨è¡¨列表输出
type ActivityListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	CoverImage string `json:"coverImage"`
	DescContent string `json:"descContent"`
	Type int `json:"type"`
	ConditionType int `json:"conditionType"`
	ConditionValue int64 `json:"conditionValue"`
	IsAutoReward int `json:"isAutoReward"`
	StartAt *gtime.Time `json:"startAt"`
	EndAt *gtime.Time `json:"endAt"`
	MaxNum int `json:"maxNum"`
	JoinNum int `json:"joinNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityListInput æ´»åŠ¨è¡¨列表查询输入
type ActivityListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Type int `json:"type"`
	ConditionType int `json:"conditionType"`
	IsAutoReward int `json:"isAutoReward"`
	Status int `json:"status"`
}

