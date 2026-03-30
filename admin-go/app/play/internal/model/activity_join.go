package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// ActivityJoin DTO 模型

// ActivityJoinCreateInput 创建æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨输入
type ActivityJoinCreateInput struct {
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	JoinStatus int `json:"joinStatus"`
	CurrentStep int `json:"currentStep"`
	FinishAt *gtime.Time `json:"finishAt"`
	RewardAt *gtime.Time `json:"rewardAt"`
	Remark string `json:"remark"`
}

// ActivityJoinUpdateInput 更新æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨输入
type ActivityJoinUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	JoinStatus int `json:"joinStatus"`
	CurrentStep int `json:"currentStep"`
	FinishAt *gtime.Time `json:"finishAt"`
	RewardAt *gtime.Time `json:"rewardAt"`
	Remark string `json:"remark"`
}

// ActivityJoinDetailOutput æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨详情输出
type ActivityJoinDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	ActivityTitle string `json:"activityTitle"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	JoinStatus int `json:"joinStatus"`
	CurrentStep int `json:"currentStep"`
	FinishAt *gtime.Time `json:"finishAt"`
	RewardAt *gtime.Time `json:"rewardAt"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityJoinListOutput æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨列表输出
type ActivityJoinListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	ActivityTitle string `json:"activityTitle"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	JoinStatus int `json:"joinStatus"`
	CurrentStep int `json:"currentStep"`
	FinishAt *gtime.Time `json:"finishAt"`
	RewardAt *gtime.Time `json:"rewardAt"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityJoinListInput æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨列表查询输入
type ActivityJoinListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	JoinStatus int `json:"joinStatus"`
}

