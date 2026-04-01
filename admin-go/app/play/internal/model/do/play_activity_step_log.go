// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityStepLog is the golang structure of table play_activity_step_log for DAO operations like Where/Data.
type PlayActivityStepLog struct {
	g.Meta      `orm:"table:play_activity_step_log, do:true"`
	Id          any         // 记录ID（Snowflake）
	ActivityId  any         // 活动ID
	StepId      any         // 步骤ID
	JoinId      any         // 参与记录ID
	MemberId    any         // 会员ID
	StepType    any         // 步骤类型：1=文字 2=链接 3=图片
	SubmitText  any         // 用户提交的文字或链接
	SubmitImage any         // 用户提交的图片URL
	AuditStatus any         // 审核状态：0=待审核 1=通过 2=驳回
	AuditRemark any         // 审核备注
	AuditBy     any         // 审核人ID
	AuditAt     *gtime.Time // 审核时间
	CreatedBy   any         // 创建人ID
	DeptId      any         // 所属部门ID
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 软删除时间
}
