// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoachApply is the golang structure of table play_coach_apply for DAO operations like Where/Data.
type PlayCoachApply struct {
	g.Meta           `orm:"table:play_coach_apply, do:true"`
	Id               any         // 申请ID（Snowflake）
	MemberId         any         // 会员ID
	RealName         any         // 真实姓名
	IdCard           any         // 身份证号
	IdCardFrontImage any         // 身份证正面照
	IdCardBackImage  any         // 身份证反面照
	SkillDesc        any         // 技能描述
	AuditStatus      any         // 审核状态:0=待审核,1=通过,2=拒绝
	AuditRemark      any         // 审核备注
	AuditAt          *gtime.Time // 审核时间
	CreatedBy        any         // 创建人ID
	DeptId           any         // 所属部门ID
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 软删除时间
}
