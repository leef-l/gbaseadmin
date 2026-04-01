// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayPayment is the golang structure for table play_payment.
type PlayPayment struct {
	Id              uint64      `orm:"id"               description:"支付记录ID（Snowflake）"`              // 支付记录ID（Snowflake）
	OrderId         uint64      `orm:"order_id"         description:"订单ID"`                           // 订单ID
	MemberId        uint64      `orm:"member_id"        description:"会员ID"`                           // 会员ID
	PaymentNo       string      `orm:"payment_no"       description:"支付流水号（平台内部）"`                    // 支付流水号（平台内部）
	TradeNo         string      `orm:"trade_no"         description:"第三方交易号"`                         // 第三方交易号
	PayType         int         `orm:"pay_type"         description:"支付方式:1=微信支付,2=支付宝支付,3=余额支付"`     // 支付方式:1=微信支付,2=支付宝支付,3=余额支付
	PayAmount       int64       `orm:"pay_amount"       description:"支付金额（分）"`                        // 支付金额（分）
	PayStatus       int         `orm:"pay_status"       description:"支付状态:0=待支付,1=支付成功,2=支付失败,3=已退款"` // 支付状态:0=待支付,1=支付成功,2=支付失败,3=已退款
	PayAt           *gtime.Time `orm:"pay_at"           description:"支付成功时间"`                         // 支付成功时间
	RefundAt        *gtime.Time `orm:"refund_at"        description:"退款时间"`                           // 退款时间
	RefundAmount    int64       `orm:"refund_amount"    description:"退款金额（分）"`                        // 退款金额（分）
	CallbackContent string      `orm:"callback_content" description:"回调报文"`                           // 回调报文
	CreatedBy       uint64      `orm:"created_by"       description:"创建人ID"`                          // 创建人ID
	DeptId          uint64      `orm:"dept_id"          description:"所属部门ID"`                         // 所属部门ID
	CreatedAt       *gtime.Time `orm:"created_at"       description:"创建时间"`                           // 创建时间
	UpdatedAt       *gtime.Time `orm:"updated_at"       description:"更新时间"`                           // 更新时间
	DeletedAt       *gtime.Time `orm:"deleted_at"       description:"软删除时间"`                          // 软删除时间
}
