package playapi

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

type sPlayapiPayment struct{}

func init() {
	service.RegisterPlayapiPayment(&sPlayapiPayment{})
}

func (s *sPlayapiPayment) Pay(ctx context.Context, req *v1.PaymentPayReq) (res *v1.PaymentPayRes, err error) {
	r := g.RequestFromCtx(ctx)
	memberID := r.GetCtxVar("jwt_member_id").Int64()
	oc := dao.PlayOrder.Columns()

	order, err := dao.PlayOrder.Ctx(ctx).Where(oc.Id, req.OrderID).Where(oc.MemberId, memberID).One()
	if err != nil {
		return
	}
	if order.IsEmpty() {
		err = gerror.New("订单不存在")
		return
	}
	if order[oc.OrderStatus].Int() != 0 {
		err = gerror.New("订单状态不可支付")
		return
	}

	res = &v1.PaymentPayRes{}

	switch req.PayType {
	case "balance":
		err = service.PaymentEnhance().BalancePay(ctx, &model.BalancePayInput{
			OrderID:  snowflake.JsonInt64(order[oc.Id].Int64()),
			MemberID: snowflake.JsonInt64(memberID),
		})
		if err != nil {
			return
		}
		res.PayResult = "success"
	case "wechat", "alipay":
		// TODO: 接入第三方支付SDK
		res.PayResult = "pending"
		res.PayParams = "{}"
	}
	return
}

func (s *sPlayapiPayment) WxCallback(ctx context.Context, req *v1.PaymentWxCallbackReq) error {
	// TODO: 验证微信签名 → 解析回调 → 更新订单状态
	return nil
}

func (s *sPlayapiPayment) AlipayCallback(ctx context.Context, req *v1.PaymentAlipayCallbackReq) error {
	// TODO: 验证支付宝签名 → 解析回调 → 更新订单状态
	return nil
}
