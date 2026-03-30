package consts

// RechargeOrderPayType 支付方式
const (
	RechargeOrderPayTypeWechat = 1 // 微信支付
	RechargeOrderPayTypeAlipay = 2 // 支付宝支付
)

// RechargeOrderPayStatus 支付状态
const (
	RechargeOrderPayStatusPending = 0 // 待支付
	RechargeOrderPayStatusSuccess = 1 // 支付成功
	RechargeOrderPayStatusFailed  = 2 // 支付失败
)
