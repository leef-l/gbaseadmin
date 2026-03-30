// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayMemberLevel is the golang structure of table play_member_level for DAO operations like Where/Data.
type PlayMemberLevel struct {
	g.Meta    `orm:"table:play_member_level, do:true"`
	Id        any         // 等级ID（Snowflake）
	Title     any         // 等级名称
	Level     any         // 等级:1=普通会员,2=白银会员,3=黄金会员,4=铂金会员,5=钻石会员
	Icon      any         // 等级图标
	MinExp    any         // 所需最低经验值
	Discount  any         // 折扣（百分比，如 90 表示九折）
	Sort      any         // 排序（升序）
	Status    any         // 状态:0=关闭,1=开启
	CreatedBy any         // 创建人ID
	DeptId    any         // 所属部门ID
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 软删除时间
}
