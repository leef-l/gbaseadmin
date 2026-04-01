// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivity is the golang structure of table play_activity for DAO operations like Where/Data.
type PlayActivity struct {
	g.Meta         `orm:"table:play_activity, do:true"`
	Id             any         // 活动ID（Snowflake）
	Title          any         // 活动名称
	CoverImage     any         // 活动封面图
	DescContent    any         // 活动详情描述（富文本，支持图文混排）
	Type           any         // 活动类型:1=充值活动,2=下单活动,3=注册活动,4=图文步骤活动,5=自定义活动
	ConditionType  any         // 参与条件:0=无条件,1=需报名,2=充值满额,3=下单满额,4=完成步骤
	ConditionValue any         // 条件值（分/次，如充值满5000分、下单满3次）
	IsAutoReward   any         // 是否自动发奖:0=否（需审核）,1=是（用户完成即发）
	StartAt        *gtime.Time // 活动开始时间
	EndAt          *gtime.Time // 活动结束时间
	MaxNum         any         // 参与人数上限（0表示不限）
	JoinNum        any         // 已参与人数
	Sort           any         // 排序（升序）
	Status         any         // 状态:0=关闭,1=开启
	CreatedBy      any         // 创建人ID
	DeptId         any         // 所属部门ID
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 软删除时间
}
