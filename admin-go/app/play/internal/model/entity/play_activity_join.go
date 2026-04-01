// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityJoin is the golang structure for table play_activity_join.
type PlayActivityJoin struct {
	Id          uint64      `orm:"id"           description:"记录ID（Snowflake）"`              // 记录ID（Snowflake）
	ActivityId  uint64      `orm:"activity_id"  description:"活动ID"`                         // 活动ID
	MemberId    uint64      `orm:"member_id"    description:"会员ID"`                         // 会员ID
	JoinStatus  int         `orm:"join_status"  description:"参与状态:0=已报名,1=进行中,2=已完成,3=已领奖"` // 参与状态:0=已报名,1=进行中,2=已完成,3=已领奖
	CurrentStep int         `orm:"current_step" description:"当前完成到第几步（步骤活动用）"`              // 当前完成到第几步（步骤活动用）
	FinishAt    *gtime.Time `orm:"finish_at"    description:"完成时间"`                         // 完成时间
	RewardAt    *gtime.Time `orm:"reward_at"    description:"领奖时间"`                         // 领奖时间
	Remark      string      `orm:"remark"       description:"备注"`                           // 备注
	CreatedBy   uint64      `orm:"created_by"   description:"创建人ID"`                        // 创建人ID
	DeptId      uint64      `orm:"dept_id"      description:"所属部门ID"`                       // 所属部门ID
	CreatedAt   *gtime.Time `orm:"created_at"   description:"创建时间"`                         // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"更新时间"`                         // 更新时间
	DeletedAt   *gtime.Time `orm:"deleted_at"   description:"软删除时间"`                        // 软删除时间
}
