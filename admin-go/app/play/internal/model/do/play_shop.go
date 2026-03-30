// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayShop is the golang structure of table play_shop for DAO operations like Where/Data.
type PlayShop struct {
	g.Meta         `orm:"table:play_shop, do:true"`
	Id             any         // 店铺ID（Snowflake）
	Title          any         // 店铺名称
	LogoImage      any         // 店铺LOGO
	CoverImage     any         // 封面图
	ContactName    any         // 联系人姓名
	ContactPhone   any         // 联系电话
	Intro          any         // 店铺简介
	CommissionRate any         // 店铺抽成比例（百分比，如 10 表示 10%）
	CoachNum       any         // 陪玩师数量
	Sort           any         // 排序（升序）
	Status         any         // 状态:0=关闭,1=开启
	CreatedBy      any         // 创建人ID
	DeptId         any         // 所属部门ID
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 软删除时间
}
