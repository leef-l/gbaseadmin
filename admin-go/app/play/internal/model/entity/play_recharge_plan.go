// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayRechargePlan is the golang structure for table play_recharge_plan.
type PlayRechargePlan struct {
	Id         uint64      `orm:"id"          description:"方案ID（Snowflake）"` // 方案ID（Snowflake）
	Title      string      `orm:"title"       description:"方案名称"`            // 方案名称
	Amount     int64       `orm:"amount"      description:"充值金额（分）"`         // 充值金额（分）
	GiftAmount int64       `orm:"gift_amount" description:"赠送金额（分）"`         // 赠送金额（分）
	CoverImage string      `orm:"cover_image" description:"方案封面图"`           // 方案封面图
	Sort       int         `orm:"sort"        description:"排序（升序）"`          // 排序（升序）
	Status     int         `orm:"status"      description:"状态:0=关闭,1=开启"`    // 状态:0=关闭,1=开启
	CreatedBy  uint64      `orm:"created_by"  description:"创建人ID"`           // 创建人ID
	DeptId     uint64      `orm:"dept_id"     description:"所属部门ID"`          // 所属部门ID
	CreatedAt  *gtime.Time `orm:"created_at"  description:"创建时间"`            // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at"  description:"更新时间"`            // 更新时间
	DeletedAt  *gtime.Time `orm:"deleted_at"  description:"软删除时间"`           // 软删除时间
}
