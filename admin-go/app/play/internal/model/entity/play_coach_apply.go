// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoachApply is the golang structure for table play_coach_apply.
type PlayCoachApply struct {
	Id               uint64      `orm:"id"                  description:"申请ID（Snowflake）"`      // 申请ID（Snowflake）
	MemberId         uint64      `orm:"member_id"           description:"会员ID"`                 // 会员ID
	RealName         string      `orm:"real_name"           description:"真实姓名"`                 // 真实姓名
	IdCard           string      `orm:"id_card"             description:"身份证号"`                 // 身份证号
	IdCardFrontImage string      `orm:"id_card_front_image" description:"身份证正面照"`               // 身份证正面照
	IdCardBackImage  string      `orm:"id_card_back_image"  description:"身份证反面照"`               // 身份证反面照
	SkillDesc        string      `orm:"skill_desc"          description:"技能描述"`                 // 技能描述
	AuditStatus      int         `orm:"audit_status"        description:"审核状态:0=待审核,1=通过,2=拒绝"` // 审核状态:0=待审核,1=通过,2=拒绝
	AuditRemark      string      `orm:"audit_remark"        description:"审核备注"`                 // 审核备注
	AuditAt          *gtime.Time `orm:"audit_at"            description:"审核时间"`                 // 审核时间
	CreatedBy        uint64      `orm:"created_by"          description:"创建人ID"`                // 创建人ID
	DeptId           uint64      `orm:"dept_id"             description:"所属部门ID"`               // 所属部门ID
	CreatedAt        *gtime.Time `orm:"created_at"          description:"创建时间"`                 // 创建时间
	UpdatedAt        *gtime.Time `orm:"updated_at"          description:"更新时间"`                 // 更新时间
	DeletedAt        *gtime.Time `orm:"deleted_at"          description:"软删除时间"`                // 软删除时间
}
