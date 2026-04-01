// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayRechargeOrder is the golang structure of table play_recharge_order for DAO operations like Where/Data.
type PlayRechargeOrder struct {
	g.Meta         `orm:"table:play_recharge_order, do:true"`
	Id             any         // 充值订单ID（Snowflake）
	OrderNo        any         // 充值订单号
	MemberId       any         // 会员ID
	RechargePlanId any         // 充值方案ID
	Amount         any         // 充值金额（分）
	GiftAmount     any         // 赠送金额（分）
	PayType        any         // 支付方式:1=微信支付,2=支付宝支付
	TradeNo        any         // 第三方交易号
	PayStatus      any         // 支付状态:0=待支付,1=支付成功,2=支付失败
	PayAt          *gtime.Time // 支付时间
	CreatedBy      any         // 创建人ID
	DeptId         any         // 所属部门ID
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 软删除时间
}
