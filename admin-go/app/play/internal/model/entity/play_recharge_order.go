// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayRechargeOrder is the golang structure for table play_recharge_order.
type PlayRechargeOrder struct {
	Id             uint64      `orm:"id"               description:"充值订单ID（Snowflake）"`        // 充值订单ID（Snowflake）
	OrderNo        string      `orm:"order_no"         description:"充值订单号"`                    // 充值订单号
	MemberId       uint64      `orm:"member_id"        description:"会员ID"`                     // 会员ID
	RechargePlanId uint64      `orm:"recharge_plan_id" description:"充值方案ID"`                   // 充值方案ID
	Amount         int64       `orm:"amount"           description:"充值金额（分）"`                  // 充值金额（分）
	GiftAmount     int64       `orm:"gift_amount"      description:"赠送金额（分）"`                  // 赠送金额（分）
	PayType        int         `orm:"pay_type"         description:"支付方式:1=微信支付,2=支付宝支付"`      // 支付方式:1=微信支付,2=支付宝支付
	TradeNo        string      `orm:"trade_no"         description:"第三方交易号"`                   // 第三方交易号
	PayStatus      int         `orm:"pay_status"       description:"支付状态:0=待支付,1=支付成功,2=支付失败"` // 支付状态:0=待支付,1=支付成功,2=支付失败
	PayAt          *gtime.Time `orm:"pay_at"           description:"支付时间"`                     // 支付时间
	CreatedBy      uint64      `orm:"created_by"       description:"创建人ID"`                    // 创建人ID
	DeptId         uint64      `orm:"dept_id"          description:"所属部门ID"`                   // 所属部门ID
	CreatedAt      *gtime.Time `orm:"created_at"       description:"创建时间"`                     // 创建时间
	UpdatedAt      *gtime.Time `orm:"updated_at"       description:"更新时间"`                     // 更新时间
	DeletedAt      *gtime.Time `orm:"deleted_at"       description:"软删除时间"`                    // 软删除时间
}
