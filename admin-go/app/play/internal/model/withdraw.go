package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Withdraw DTO 模型

// WithdrawCreateInput 创建陪玩师提现记录输入
type WithdrawCreateInput struct {
	CoachID snowflake.JsonInt64 `json:"coachID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	Amount int `json:"amount"`
	Status int `json:"status"`
	Reason string `json:"reason"`
	AuditedAt *gtime.Time `json:"auditedAt"`
}

// WithdrawUpdateInput 更新陪玩师提现记录输入
type WithdrawUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	Amount int `json:"amount"`
	Status int `json:"status"`
	Reason string `json:"reason"`
	AuditedAt *gtime.Time `json:"auditedAt"`
}

// WithdrawDetailOutput 陪玩师提现记录详情输出
type WithdrawDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	CoachRealName string `json:"coachRealName"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	MemberNickname string `json:"memberNickname"`
	Amount int `json:"amount"`
	Status int `json:"status"`
	Reason string `json:"reason"`
	AuditedAt *gtime.Time `json:"auditedAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// WithdrawListOutput 陪玩师提现记录列表输出
type WithdrawListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	CoachRealName string `json:"coachRealName"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	MemberNickname string `json:"memberNickname"`
	Amount int `json:"amount"`
	Status int `json:"status"`
	Reason string `json:"reason"`
	AuditedAt *gtime.Time `json:"auditedAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// WithdrawListInput 陪玩师提现记录列表查询输入
type WithdrawListInput struct {
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
	OrderBy   string `json:"orderBy"`
	OrderDir  string `json:"orderDir"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

