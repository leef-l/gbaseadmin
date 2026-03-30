package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// ActivityReward DTO 模型

// ActivityRewardCreateInput 创建活动奖励表输入
type ActivityRewardCreateInput struct {
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	RewardType int `json:"rewardType"`
	RewardValue int64 `json:"rewardValue"`
	RewardName string `json:"rewardName"`
	Sort int `json:"sort"`
}

// ActivityRewardUpdateInput 更新活动奖励表输入
type ActivityRewardUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	RewardType int `json:"rewardType"`
	RewardValue int64 `json:"rewardValue"`
	RewardName string `json:"rewardName"`
	Sort int `json:"sort"`
}

// ActivityRewardDetailOutput 活动奖励表详情输出
type ActivityRewardDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	ActivityTitle string `json:"activityTitle"`
	RewardType int `json:"rewardType"`
	RewardValue int64 `json:"rewardValue"`
	RewardName string `json:"rewardName"`
	Sort int `json:"sort"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityRewardListOutput 活动奖励表列表输出
type ActivityRewardListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	ActivityTitle string `json:"activityTitle"`
	RewardType int `json:"rewardType"`
	RewardValue int64 `json:"rewardValue"`
	RewardName string `json:"rewardName"`
	Sort int `json:"sort"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityRewardListInput 活动奖励表列表查询输入
type ActivityRewardListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	RewardType int `json:"rewardType"`
}

