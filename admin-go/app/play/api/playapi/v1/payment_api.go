package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ========== MemberAuth 接口 ==========

type PaymentPayReq struct {
	g.Meta  `path:"/payment/pay" method:"post" tags:"C端支付" summary:"发起支付"`
	OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
	PayType string `json:"payType" v:"required|in:balance,wechat,alipay#支付方式不能为空|支付方式不合法" dc:"支付方式:balance=余额,wechat=微信,alipay=支付宝"`
}

type PaymentPayRes struct {
	g.Meta    `mime:"application/json"`
	PayResult string `json:"payResult" dc:"支付结果:success=余额支付成功,pending=等待第三方支付"`
	PayParams string `json:"payParams" dc:"第三方支付参数(JSON字符串)"`
}

// ========== 公开回调接口 ==========

type PaymentWxCallbackReq struct {
	g.Meta `path:"/payment/wx_callback" method:"post" tags:"C端支付" summary:"微信支付回调"`
}

type PaymentWxCallbackRes struct {
	g.Meta `mime:"application/xml"`
}

type PaymentAlipayCallbackReq struct {
	g.Meta `path:"/payment/alipay_callback" method:"post" tags:"C端支付" summary:"支付宝支付回调"`
}

type PaymentAlipayCallbackRes struct {
	g.Meta `mime:"text/plain"`
}
