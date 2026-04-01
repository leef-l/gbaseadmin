// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityStepLog is the golang structure for table play_activity_step_log.
type PlayActivityStepLog struct {
	Id          uint64      `orm:"id"           description:"记录ID（Snowflake）"`      // 记录ID（Snowflake）
	ActivityId  uint64      `orm:"activity_id"  description:"活动ID"`                 // 活动ID
	StepId      uint64      `orm:"step_id"      description:"步骤ID"`                 // 步骤ID
	JoinId      uint64      `orm:"join_id"      description:"参与记录ID"`               // 参与记录ID
	MemberId    uint64      `orm:"member_id"    description:"会员ID"`                 // 会员ID
	StepType    int         `orm:"step_type"    description:"步骤类型：1=文字 2=链接 3=图片"`  // 步骤类型：1=文字 2=链接 3=图片
	SubmitText  string      `orm:"submit_text"  description:"用户提交的文字或链接"`           // 用户提交的文字或链接
	SubmitImage string      `orm:"submit_image" description:"用户提交的图片URL"`           // 用户提交的图片URL
	AuditStatus int         `orm:"audit_status" description:"审核状态：0=待审核 1=通过 2=驳回"` // 审核状态：0=待审核 1=通过 2=驳回
	AuditRemark string      `orm:"audit_remark" description:"审核备注"`                 // 审核备注
	AuditBy     uint64      `orm:"audit_by"     description:"审核人ID"`                // 审核人ID
	AuditAt     *gtime.Time `orm:"audit_at"     description:"审核时间"`                 // 审核时间
	CreatedBy   uint64      `orm:"created_by"   description:"创建人ID"`                // 创建人ID
	DeptId      uint64      `orm:"dept_id"      description:"所属部门ID"`               // 所属部门ID
	CreatedAt   *gtime.Time `orm:"created_at"   description:"创建时间"`                 // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"更新时间"`                 // 更新时间
	DeletedAt   *gtime.Time `orm:"deleted_at"   description:"软删除时间"`                // 软删除时间
}
