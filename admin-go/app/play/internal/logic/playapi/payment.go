package playapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/alipay"
	"gbaseadmin/utility/snowflake"
	"gbaseadmin/utility/wxpay"
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

	case "wechat":
		devMode := g.Cfg().MustGet(ctx, "pay.devMode").Bool()
		if devMode {
			// 开发模式：直接模拟支付成功
			err = s.markOrderPaid(ctx, order[oc.Id].Int64(), memberID, order[oc.PayAmount].Int64(), 1)
			if err != nil {
				return
			}
			res.PayResult = "success"
			res.PayParams = `{"mock":"devMode wechat pay success"}`
		} else {
			client, e := wxpay.New(ctx)
			if e != nil {
				err = gerror.Newf("初始化微信支付失败: %v", e)
				return
			}
			orderNo := order[oc.OrderNo].String()
			amount := order[oc.PayAmount].Int64()
			desc := order[oc.GoodsTitle].String()
			payParams, e := client.CreateOrder(ctx, orderNo, amount, desc)
			if e != nil {
				err = gerror.Newf("创建微信支付订单失败: %v", e)
				return
			}
			paramsBytes, _ := json.Marshal(payParams)
			res.PayResult = "pending"
			res.PayParams = string(paramsBytes)
		}

	case "alipay":
		devMode := g.Cfg().MustGet(ctx, "pay.devMode").Bool()
		if devMode {
			// 开发模式：直接模拟支付成功
			err = s.markOrderPaid(ctx, order[oc.Id].Int64(), memberID, order[oc.PayAmount].Int64(), 2)
			if err != nil {
				return
			}
			res.PayResult = "success"
			res.PayParams = `{"mock":"devMode alipay pay success"}`
		} else {
			client, e := alipay.New(ctx)
			if e != nil {
				err = gerror.Newf("初始化支付宝失败: %v", e)
				return
			}
			orderNo := order[oc.OrderNo].String()
			amount := order[oc.PayAmount].Int64()
			desc := order[oc.GoodsTitle].String()
			payURL, e := client.CreateOrder(ctx, orderNo, amount, desc)
			if e != nil {
				err = gerror.Newf("创建支付宝订单失败: %v", e)
				return
			}
			res.PayResult = "pending"
			res.PayParams = fmt.Sprintf(`{"pay_url":"%s"}`, payURL)
		}
	}
	return
}

// markOrderPaid 将订单标记为已支付，并创建支付记录
// payTypeInt: 1=微信, 2=支付宝, 3=余额
func (s *sPlayapiPayment) markOrderPaid(ctx context.Context, orderID, memberID, payAmount int64, payTypeInt int) error {
	return dao.PlayOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		oc := dao.PlayOrder.Columns()

		// 防重：再次确认订单状态
		order, err := dao.PlayOrder.Ctx(ctx).Where(oc.Id, orderID).One()
		if err != nil {
			return err
		}
		if order.IsEmpty() {
			return gerror.New("订单不存在")
		}
		if order[oc.OrderStatus].Int() != 0 {
			// 已支付，幂等处理
			return nil
		}

		// 创建支付记录
		paymentID := snowflake.Generate()
		paymentNo := fmt.Sprintf("PAY%d", paymentID)
		_, err = dao.PlayPayment.Ctx(ctx).Data(g.Map{
			dao.PlayPayment.Columns().Id:        paymentID,
			dao.PlayPayment.Columns().OrderId:   orderID,
			dao.PlayPayment.Columns().MemberId:  memberID,
			dao.PlayPayment.Columns().PaymentNo: paymentNo,
			dao.PlayPayment.Columns().PayType:   payTypeInt,
			dao.PlayPayment.Columns().PayAmount: payAmount,
			dao.PlayPayment.Columns().PayStatus: 1,
			dao.PlayPayment.Columns().PayAt:     gtime.Now(),
			dao.PlayPayment.Columns().CreatedAt: gtime.Now(),
			dao.PlayPayment.Columns().UpdatedAt: gtime.Now(),
		}).Insert()
		if err != nil {
			return err
		}

		// 更新订单状态为已支付(1)
		_, err = dao.PlayOrder.Ctx(ctx).Where(oc.Id, orderID).Data(g.Map{
			oc.OrderStatus: 1,
			oc.PayType:     payTypeInt,
			oc.PayAt:       gtime.Now(),
			oc.UpdatedAt:   gtime.Now(),
		}).Update()
		return err
	})
}

func (s *sPlayapiPayment) WxCallback(ctx context.Context, req *v1.PaymentWxCallbackReq) error {
	devMode := g.Cfg().MustGet(ctx, "pay.devMode").Bool()
	if devMode {
		g.Log().Info(ctx, "WxCallback: devMode 模式，跳过真实回调处理")
		return nil
	}

	r := g.RequestFromCtx(ctx)
	client, err := wxpay.New(ctx)
	if err != nil {
		return gerror.Newf("初始化微信支付失败: %v", err)
	}

	outTradeNo, err := client.VerifyNotify(ctx, r.Request)
	if err != nil {
		return gerror.Newf("微信回调验签失败: %v", err)
	}

	// 根据订单号查询订单
	oc := dao.PlayOrder.Columns()
	order, err := dao.PlayOrder.Ctx(ctx).Where(oc.OrderNo, outTradeNo).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return gerror.Newf("订单不存在: %s", outTradeNo)
	}

	return s.markOrderPaid(ctx,
		order[oc.Id].Int64(),
		order[oc.MemberId].Int64(),
		order[oc.PayAmount].Int64(),
		1, // 微信
	)
}

func (s *sPlayapiPayment) AlipayCallback(ctx context.Context, req *v1.PaymentAlipayCallbackReq) error {
	devMode := g.Cfg().MustGet(ctx, "pay.devMode").Bool()
	if devMode {
		g.Log().Info(ctx, "AlipayCallback: devMode 模式，跳过真实回调处理")
		return nil
	}

	r := g.RequestFromCtx(ctx)
	client, err := alipay.New(ctx)
	if err != nil {
		return gerror.Newf("初始化支付宝失败: %v", err)
	}

	// 获取所有 POST 参数
	postForm := url.Values{}
	for k, v := range r.GetFormMapStrStr() {
		postForm.Set(k, v)
	}

	outTradeNo, err := client.VerifyNotify(ctx, postForm)
	if err != nil {
		return gerror.Newf("支付宝回调验签失败: %v", err)
	}

	// 根据订单号查询订单
	oc := dao.PlayOrder.Columns()
	order, err := dao.PlayOrder.Ctx(ctx).Where(oc.OrderNo, outTradeNo).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return gerror.Newf("订单不存在: %s", outTradeNo)
	}

	return s.markOrderPaid(ctx,
		order[oc.Id].Int64(),
		order[oc.MemberId].Int64(),
		order[oc.PayAmount].Int64(),
		2, // 支付宝
	)
}
