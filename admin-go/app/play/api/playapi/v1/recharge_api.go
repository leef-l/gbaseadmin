package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 充值方案列表（MemberAuth） ====================

type RechargePlansReq struct {
	g.Meta `path:"/recharge/plans" method:"get" tags:"C端充值" summary:"充值方案列表"`
}

type RechargePlansRes struct {
	g.Meta `mime:"application/json"`
	List   []RechargePlanItem `json:"list" dc:"充值方案列表"`
}

type RechargePlanItem struct {
	PlanID     string `json:"planId" dc:"方案ID"`
	Title      string `json:"title" dc:"方案标题"`
	Amount     int64  `json:"amount" dc:"充值金额(分)"`
	GiveAmount int64  `json:"giveAmount" dc:"赠送金额(分)"`
	Tag        string `json:"tag" dc:"标签(如:推荐/热门)"`
	Sort       int    `json:"sort" dc:"排序"`
}

// ==================== 创建充值订单（MemberAuth） ====================

type RechargeCreateReq struct {
	g.Meta  `path:"/recharge/create" method:"post" tags:"C端充值" summary:"创建充值订单"`
	PlanID  string `json:"planId" v:"required#充值方案ID不能为空" dc:"充值方案ID"`
	PayType string `json:"payType" v:"required|in:wechat,alipay#支付方式不能为空|支付方式不合法" dc:"支付方式:wechat=微信,alipay=支付宝"`
}

type RechargeCreateRes struct {
	g.Meta    `mime:"application/json"`
	OrderID   string `json:"orderId" dc:"充值订单ID"`
	PayParams string `json:"payParams" dc:"第三方支付参数(JSON字符串)"`
}

// ==================== 充值微信回调（公开） ====================

type RechargeWxNotifyReq struct {
	g.Meta `path:"/recharge/wx_callback" method:"post" tags:"C端充值" summary:"充值微信回调"`
}

type RechargeWxNotifyRes struct {
	g.Meta `mime:"application/xml"`
}

// ==================== 充值支付宝回调（公开） ====================

type RechargeAlipayNotifyReq struct {
	g.Meta `path:"/recharge/alipay_callback" method:"post" tags:"C端充值" summary:"充值支付宝回调"`
}

type RechargeAlipayNotifyRes struct {
	g.Meta `mime:"text/plain"`
}
