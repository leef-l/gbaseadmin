// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityReward is the golang structure of table play_activity_reward for DAO operations like Where/Data.
type PlayActivityReward struct {
	g.Meta      `orm:"table:play_activity_reward, do:true"`
	Id          any         // 奖励ID（Snowflake）
	ActivityId  any         // 活动ID
	RewardType  any         // 奖励类型:1=余额,2=优惠券,3=经验值,4=会员等级天数
	RewardValue any         // 奖励数值（余额=分，优惠券=coupon_id，经验=值，等级天数=天）
	RewardName  any         // 奖励名称（展示用，如"送50元余额"）
	Sort        any         // 排序（升序）
	CreatedBy   any         // 创建人ID
	DeptId      any         // 所属部门ID
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 软删除时间
}
