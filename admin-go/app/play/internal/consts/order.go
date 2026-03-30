package consts

// OrderPayType 支付方式
const (
	OrderPayType0 = 0 // 未支付
	OrderPayType1 = 1 // 微信支付
	OrderPayType2 = 2 // 支付宝支付
	OrderPayType3 = 3 // 余额支付
)

// OrderOrderStatus 订单状态
const (
	OrderOrderStatus0 = 0 // 待支付
	OrderOrderStatus1 = 1 // 已支付
	OrderOrderStatus2 = 2 // 进行中
	OrderOrderStatus3 = 3 // 已完成
	OrderOrderStatus4 = 4 // 已取消
	OrderOrderStatus5 = 5 // 退款中
	OrderOrderStatus6 = 6 // 已退款
)

