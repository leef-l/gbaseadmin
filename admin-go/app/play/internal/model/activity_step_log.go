package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// ActivityStepLog DTO 模型

// ActivityStepLogCreateInput 创建活动步骤提交记录输入
type ActivityStepLogCreateInput struct {
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	StepID snowflake.JsonInt64 `json:"stepID"`
	JoinID snowflake.JsonInt64 `json:"joinID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	StepType int `json:"stepType"`
	SubmitText string `json:"submitText"`
	SubmitImage string `json:"submitImage"`
	AuditStatus int `json:"auditStatus"`
	AuditRemark string `json:"auditRemark"`
	AuditBy int64 `json:"auditBy"`
	AuditAt *gtime.Time `json:"auditAt"`
}

// ActivityStepLogUpdateInput 更新活动步骤提交记录输入
type ActivityStepLogUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	StepID snowflake.JsonInt64 `json:"stepID"`
	JoinID snowflake.JsonInt64 `json:"joinID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	StepType int `json:"stepType"`
	SubmitText string `json:"submitText"`
	SubmitImage string `json:"submitImage"`
	AuditStatus int `json:"auditStatus"`
	AuditRemark string `json:"auditRemark"`
	AuditBy int64 `json:"auditBy"`
	AuditAt *gtime.Time `json:"auditAt"`
}

// ActivityStepLogDetailOutput 活动步骤提交记录详情输出
type ActivityStepLogDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	ActivityTitle string `json:"activityTitle"`
	StepID snowflake.JsonInt64 `json:"stepID"`
	JoinID snowflake.JsonInt64 `json:"joinID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	StepType int `json:"stepType"`
	SubmitText string `json:"submitText"`
	SubmitImage string `json:"submitImage"`
	AuditStatus int `json:"auditStatus"`
	AuditRemark string `json:"auditRemark"`
	AuditBy int64 `json:"auditBy"`
	AuditAt *gtime.Time `json:"auditAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityStepLogListOutput 活动步骤提交记录列表输出
type ActivityStepLogListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ActivityID snowflake.JsonInt64 `json:"activityID"`
	ActivityTitle string `json:"activityTitle"`
	StepID snowflake.JsonInt64 `json:"stepID"`
	JoinID snowflake.JsonInt64 `json:"joinID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	StepType int `json:"stepType"`
	SubmitText string `json:"submitText"`
	SubmitImage string `json:"submitImage"`
	AuditStatus int `json:"auditStatus"`
	AuditRemark string `json:"auditRemark"`
	AuditBy int64 `json:"auditBy"`
	AuditAt *gtime.Time `json:"auditAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ActivityStepLogListInput 活动步骤提交记录列表查询输入
type ActivityStepLogListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	StepType int `json:"stepType"`
	AuditStatus int `json:"auditStatus"`
}

