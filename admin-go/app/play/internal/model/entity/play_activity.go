// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivity is the golang structure for table play_activity.
type PlayActivity struct {
	Id             uint64      `orm:"id"              description:"活动ID（Snowflake）"`                            // 活动ID（Snowflake）
	Title          string      `orm:"title"           description:"活动名称"`                                       // 活动名称
	CoverImage     string      `orm:"cover_image"     description:"活动封面图"`                                      // 活动封面图
	DescContent    string      `orm:"desc_content"    description:"活动详情描述（富文本，支持图文混排）"`                         // 活动详情描述（富文本，支持图文混排）
	Type           int         `orm:"type"            description:"活动类型:1=充值活动,2=下单活动,3=注册活动,4=图文步骤活动,5=自定义活动"` // 活动类型:1=充值活动,2=下单活动,3=注册活动,4=图文步骤活动,5=自定义活动
	ConditionType  int         `orm:"condition_type"  description:"参与条件:0=无条件,1=需报名,2=充值满额,3=下单满额,4=完成步骤"`      // 参与条件:0=无条件,1=需报名,2=充值满额,3=下单满额,4=完成步骤
	ConditionValue int64       `orm:"condition_value" description:"条件值（分/次，如充值满5000分、下单满3次）"`                   // 条件值（分/次，如充值满5000分、下单满3次）
	IsAutoReward   int         `orm:"is_auto_reward"  description:"是否自动发奖:0=否（需审核）,1=是（用户完成即发）"`                // 是否自动发奖:0=否（需审核）,1=是（用户完成即发）
	StartAt        *gtime.Time `orm:"start_at"        description:"活动开始时间"`                                     // 活动开始时间
	EndAt          *gtime.Time `orm:"end_at"          description:"活动结束时间"`                                     // 活动结束时间
	MaxNum         int         `orm:"max_num"         description:"参与人数上限（0表示不限）"`                              // 参与人数上限（0表示不限）
	JoinNum        int         `orm:"join_num"        description:"已参与人数"`                                      // 已参与人数
	Sort           int         `orm:"sort"            description:"排序（升序）"`                                     // 排序（升序）
	Status         int         `orm:"status"          description:"状态:0=关闭,1=开启"`                               // 状态:0=关闭,1=开启
	CreatedBy      uint64      `orm:"created_by"      description:"创建人ID"`                                      // 创建人ID
	DeptId         uint64      `orm:"dept_id"         description:"所属部门ID"`                                     // 所属部门ID
	CreatedAt      *gtime.Time `orm:"created_at"      description:"创建时间"`                                       // 创建时间
	UpdatedAt      *gtime.Time `orm:"updated_at"      description:"更新时间"`                                       // 更新时间
	DeletedAt      *gtime.Time `orm:"deleted_at"      description:"软删除时间"`                                      // 软删除时间
}
