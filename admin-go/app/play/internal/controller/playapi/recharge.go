package playapi

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Recharge = &cRecharge{}

type cRecharge struct{}

// Plans 充值方案列表
func (c *cRecharge) Plans(ctx context.Context, req *v1.RechargePlansReq) (res *v1.RechargePlansRes, err error) {
	res = &v1.RechargePlansRes{}
	res.List, err = service.PlayapiRecharge().Plans(ctx)
	return
}

// Create 创建充值订单
func (c *cRecharge) Create(ctx context.Context, req *v1.RechargeCreateReq) (res *v1.RechargeCreateRes, err error) {
	res = &v1.RechargeCreateRes{}
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	res.OrderID, res.PayParams, err = service.PlayapiRecharge().Create(ctx, memberID, req.PlanID, req.PayType)
	return
}

var RechargeNotify = &cRechargeNotify{}

type cRechargeNotify struct{}

// WxNotify 充值微信回调
func (c *cRechargeNotify) WxNotify(ctx context.Context, req *v1.RechargeWxNotifyReq) (res *v1.RechargeWxNotifyRes, err error) {
	err = service.PlayapiRecharge().WxNotify(ctx)
	return
}

// AlipayNotify 充值支付宝回调
func (c *cRechargeNotify) AlipayNotify(ctx context.Context, req *v1.RechargeAlipayNotifyReq) (res *v1.RechargeAlipayNotifyRes, err error) {
	err = service.PlayapiRecharge().AlipayNotify(ctx)
	return
}
