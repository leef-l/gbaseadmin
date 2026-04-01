// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayRechargePlan is the golang structure of table play_recharge_plan for DAO operations like Where/Data.
type PlayRechargePlan struct {
	g.Meta     `orm:"table:play_recharge_plan, do:true"`
	Id         any         // 方案ID（Snowflake）
	Title      any         // 方案名称
	Amount     any         // 充值金额（分）
	GiftAmount any         // 赠送金额（分）
	CoverImage any         // 方案封面图
	Sort       any         // 排序（升序）
	Status     any         // 状态:0=关闭,1=开启
	CreatedBy  any         // 创建人ID
	DeptId     any         // 所属部门ID
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 软删除时间
}
