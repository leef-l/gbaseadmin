package model

import "gbaseadmin/utility/snowflake"

// CoachApplyAuditInput 审核陪玩师申请输入
type CoachApplyAuditInput struct {
	ID          snowflake.JsonInt64 `json:"id"`
	AuditStatus int                 `json:"auditStatus"`
	AuditRemark string              `json:"auditRemark"`
}
