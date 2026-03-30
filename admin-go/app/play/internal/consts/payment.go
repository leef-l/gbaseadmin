package consts

// PaymentPayType 支付方式
const (
	PaymentPayTypeWechat  = 1 // 微信支付
	PaymentPayTypeAlipay  = 2 // 支付宝支付
	PaymentPayTypeBalance = 3 // 余额支付
)

// PaymentPayStatus 支付状态
const (
	PaymentPayStatusPending = 0 // 待支付
	PaymentPayStatusSuccess = 1 // 支付成功
	PaymentPayStatusFailed  = 2 // 支付失败
	PaymentPayStatusRefunded = 3 // 已退款
)
