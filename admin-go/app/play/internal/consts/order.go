package consts

// OrderPayType 支付方式
const (
	OrderPayTypeUnpaid  = 0 // 未支付
	OrderPayTypeWechat  = 1 // 微信支付
	OrderPayTypeAlipay  = 2 // 支付宝支付
	OrderPayTypeBalance = 3 // 余额支付
)

// OrderOrderStatus 订单状态
const (
	OrderOrderStatusPending    = 0 // 待支付
	OrderOrderStatusPaid       = 1 // 已支付
	OrderOrderStatusInProgress = 2 // 进行中
	OrderOrderStatusCompleted  = 3 // 已完成
	OrderOrderStatusCancelled  = 4 // 已取消
	OrderOrderStatusRefunding  = 5 // 退款中
	OrderOrderStatusRefunded   = 6 // 已退款
)
