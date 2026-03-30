// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoachLevel is the golang structure of table play_coach_level for DAO operations like Where/Data.
type PlayCoachLevel struct {
	g.Meta         `orm:"table:play_coach_level, do:true"`
	Id             any         // 等级ID（Snowflake）
	Title          any         // 等级名称
	Level          any         // 等级:1=青铜,2=白银,3=黄金,4=铂金,5=钻石
	Icon           any         // 等级图标
	MinOrders      any         // 所需最低接单数
	MinScore       any         // 所需最低评分（乘100存储，如 450=4.50分）
	CommissionRate any         // 平台抽成比例（百分比，如 20 表示 20%）
	Sort           any         // 排序（升序）
	Status         any         // 状态:0=关闭,1=开启
	CreatedBy      any         // 创建人ID
	DeptId         any         // 所属部门ID
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 软删除时间
}
