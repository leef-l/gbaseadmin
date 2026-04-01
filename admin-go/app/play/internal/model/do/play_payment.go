// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayPayment is the golang structure of table play_payment for DAO operations like Where/Data.
type PlayPayment struct {
	g.Meta          `orm:"table:play_payment, do:true"`
	Id              any         // 支付记录ID（Snowflake）
	OrderId         any         // 订单ID
	MemberId        any         // 会员ID
	PaymentNo       any         // 支付流水号（平台内部）
	TradeNo         any         // 第三方交易号
	PayType         any         // 支付方式:1=微信支付,2=支付宝支付,3=余额支付
	PayAmount       any         // 支付金额（分）
	PayStatus       any         // 支付状态:0=待支付,1=支付成功,2=支付失败,3=已退款
	PayAt           *gtime.Time // 支付成功时间
	RefundAt        *gtime.Time // 退款时间
	RefundAmount    any         // 退款金额（分）
	CallbackContent any         // 回调报文
	CreatedBy       any         // 创建人ID
	DeptId          any         // 所属部门ID
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 软删除时间
}
