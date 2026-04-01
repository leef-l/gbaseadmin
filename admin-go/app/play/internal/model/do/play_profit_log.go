// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayProfitLog is the golang structure of table play_profit_log for DAO operations like Where/Data.
type PlayProfitLog struct {
	g.Meta         `orm:"table:play_profit_log, do:true"`
	Id             any         // 流水ID（Snowflake）
	OrderId        any         // 订单ID
	OrderNo        any         // 订单编号
	PayAmount      any         // 实付金额（分）
	CoachId        any         // 陪玩师ID
	ShopId         any         // 店铺ID（0表示无店铺）
	PlatformRate   any         // 平台抽成比例（百分比）
	PlatformAmount any         // 平台抽成金额（分）
	ShopRate       any         // 店铺抽成比例（百分比）
	ShopAmount     any         // 店铺抽成金额（分）
	CoachAmount    any         // 陪玩师收入（分）
	SettleStatus   any         // 结算状态:0=待结算,1=已结算
	SettleAt       *gtime.Time // 结算时间
	CreatedBy      any         // 创建人ID
	DeptId         any         // 所属部门ID
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 软删除时间
}
