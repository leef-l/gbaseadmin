// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayProfitLog is the golang structure for table play_profit_log.
type PlayProfitLog struct {
	Id             uint64      `orm:"id"              description:"流水ID（Snowflake）"`  // 流水ID（Snowflake）
	OrderId        uint64      `orm:"order_id"        description:"订单ID"`             // 订单ID
	OrderNo        string      `orm:"order_no"        description:"订单编号"`             // 订单编号
	PayAmount      int64       `orm:"pay_amount"      description:"实付金额（分）"`          // 实付金额（分）
	CoachId        uint64      `orm:"coach_id"        description:"陪玩师ID"`            // 陪玩师ID
	ShopId         uint64      `orm:"shop_id"         description:"店铺ID（0表示无店铺）"`     // 店铺ID（0表示无店铺）
	PlatformRate   int         `orm:"platform_rate"   description:"平台抽成比例（百分比）"`      // 平台抽成比例（百分比）
	PlatformAmount int64       `orm:"platform_amount" description:"平台抽成金额（分）"`        // 平台抽成金额（分）
	ShopRate       int         `orm:"shop_rate"       description:"店铺抽成比例（百分比）"`      // 店铺抽成比例（百分比）
	ShopAmount     int64       `orm:"shop_amount"     description:"店铺抽成金额（分）"`        // 店铺抽成金额（分）
	CoachAmount    int64       `orm:"coach_amount"    description:"陪玩师收入（分）"`         // 陪玩师收入（分）
	SettleStatus   int         `orm:"settle_status"   description:"结算状态:0=待结算,1=已结算"` // 结算状态:0=待结算,1=已结算
	SettleAt       *gtime.Time `orm:"settle_at"       description:"结算时间"`             // 结算时间
	CreatedBy      uint64      `orm:"created_by"      description:"创建人ID"`            // 创建人ID
	DeptId         uint64      `orm:"dept_id"         description:"所属部门ID"`           // 所属部门ID
	CreatedAt      *gtime.Time `orm:"created_at"      description:"创建时间"`             // 创建时间
	UpdatedAt      *gtime.Time `orm:"updated_at"      description:"更新时间"`             // 更新时间
	DeletedAt      *gtime.Time `orm:"deleted_at"      description:"软删除时间"`            // 软删除时间
}
