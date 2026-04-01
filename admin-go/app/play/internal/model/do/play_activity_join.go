// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityJoin is the golang structure of table play_activity_join for DAO operations like Where/Data.
type PlayActivityJoin struct {
	g.Meta      `orm:"table:play_activity_join, do:true"`
	Id          any         // 记录ID（Snowflake）
	ActivityId  any         // 活动ID
	MemberId    any         // 会员ID
	JoinStatus  any         // 参与状态:0=已报名,1=进行中,2=已完成,3=已领奖
	CurrentStep any         // 当前完成到第几步（步骤活动用）
	FinishAt    *gtime.Time // 完成时间
	RewardAt    *gtime.Time // 领奖时间
	Remark      any         // 备注
	CreatedBy   any         // 创建人ID
	DeptId      any         // 所属部门ID
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 软删除时间
}
