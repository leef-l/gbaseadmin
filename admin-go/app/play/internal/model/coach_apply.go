package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// CoachApply DTO 模型

// CoachApplyCreateInput 创建陪玩师申请表输入
type CoachApplyCreateInput struct {
	MemberID snowflake.JsonInt64 `json:"memberID"`
	RealName string `json:"realName"`
	IDCard string `json:"idCard"`
	IDCardFrontImage string `json:"idCardFrontImage"`
	IDCardBackImage string `json:"idCardBackImage"`
	SkillDesc string `json:"skillDesc"`
	AuditStatus int `json:"auditStatus"`
	AuditRemark string `json:"auditRemark"`
	AuditAt *gtime.Time `json:"auditAt"`
}

// CoachApplyUpdateInput 更新陪玩师申请表输入
type CoachApplyUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	RealName string `json:"realName"`
	IDCard string `json:"idCard"`
	IDCardFrontImage string `json:"idCardFrontImage"`
	IDCardBackImage string `json:"idCardBackImage"`
	SkillDesc string `json:"skillDesc"`
	AuditStatus int `json:"auditStatus"`
	AuditRemark string `json:"auditRemark"`
	AuditAt *gtime.Time `json:"auditAt"`
}

// CoachApplyDetailOutput 陪玩师申请表详情输出
type CoachApplyDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	RealName string `json:"realName"`
	IDCard string `json:"idCard"`
	IDCardFrontImage string `json:"idCardFrontImage"`
	IDCardBackImage string `json:"idCardBackImage"`
	SkillDesc string `json:"skillDesc"`
	AuditStatus int `json:"auditStatus"`
	AuditRemark string `json:"auditRemark"`
	AuditAt *gtime.Time `json:"auditAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CoachApplyListOutput 陪玩师申请表列表输出
type CoachApplyListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	RealName string `json:"realName"`
	IDCard string `json:"idCard"`
	IDCardFrontImage string `json:"idCardFrontImage"`
	IDCardBackImage string `json:"idCardBackImage"`
	SkillDesc string `json:"skillDesc"`
	AuditStatus int `json:"auditStatus"`
	AuditRemark string `json:"auditRemark"`
	AuditAt *gtime.Time `json:"auditAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CoachApplyListInput 陪玩师申请表列表查询输入
type CoachApplyListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	AuditStatus int `json:"auditStatus"`
}

