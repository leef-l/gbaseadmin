// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityReward is the golang structure for table play_activity_reward.
type PlayActivityReward struct {
	Id          uint64      `orm:"id"           description:"奖励ID（Snowflake）"`                      // 奖励ID（Snowflake）
	ActivityId  uint64      `orm:"activity_id"  description:"活动ID"`                                 // 活动ID
	RewardType  int         `orm:"reward_type"  description:"奖励类型:1=余额,2=优惠券,3=经验值,4=会员等级天数"`       // 奖励类型:1=余额,2=优惠券,3=经验值,4=会员等级天数
	RewardValue int64       `orm:"reward_value" description:"奖励数值（余额=分，优惠券=coupon_id，经验=值，等级天数=天）"` // 奖励数值（余额=分，优惠券=coupon_id，经验=值，等级天数=天）
	RewardName  string      `orm:"reward_name"  description:"奖励名称（展示用，如\"送50元余额\"）"`                // 奖励名称（展示用，如"送50元余额"）
	Sort        int         `orm:"sort"         description:"排序（升序）"`                               // 排序（升序）
	CreatedBy   uint64      `orm:"created_by"   description:"创建人ID"`                                // 创建人ID
	DeptId      uint64      `orm:"dept_id"      description:"所属部门ID"`                               // 所属部门ID
	CreatedAt   *gtime.Time `orm:"created_at"   description:"创建时间"`                                 // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"更新时间"`                                 // 更新时间
	DeletedAt   *gtime.Time `orm:"deleted_at"   description:"软删除时间"`                                // 软删除时间
}
