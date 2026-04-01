package playapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

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

type sRecharge struct{}

// Plans 充值方案列表
func (s *sRecharge) Plans(ctx context.Context) (list []v1.RechargePlanItem, err error) {
	var records []struct {
		Id         uint64 `json:"id"`
		Title      string `json:"title"`
		Amount     int64  `json:"amount"`
		GiftAmount int64  `json:"gift_amount"`
		Sort       int    `json:"sort"`
	}
	err = dao.PlayRechargePlan.Ctx(ctx).
		Where(dao.PlayRechargePlan.Columns().Status, 1).
		OrderAsc(dao.PlayRechargePlan.Columns().Sort).
		Scan(&records)
	if err != nil {
		return
	}
	list = make([]v1.RechargePlanItem, 0, len(records))
	for _, r := range records {
		list = append(list, v1.RechargePlanItem{
			PlanID:     strconv.FormatUint(r.Id, 10),
			Title:      r.Title,
			Amount:     r.Amount,
			GiveAmount: r.GiftAmount,
			Sort:       r.Sort,
		})
	}
	return
}

// Create 创建充值订单
func (s *sRecharge) Create(ctx context.Context, memberID int64, planID string, payType string) (orderID string, payParams string, err error) {
	pid, _ := strconv.ParseUint(planID, 10, 64)
	// 查询充值方案
	rpc := dao.PlayRechargePlan.Columns()
	plan, err := dao.PlayRechargePlan.Ctx(ctx).
		Where(rpc.Id, pid).
		Where(rpc.Status, 1).
		Where(rpc.DeletedAt, nil).
		One()
	if err != nil {
		return
	}
	if plan.IsEmpty() {
		err = gerror.New("充值方案不存在")
		return
	}

	// 生成订单
	id := snowflake.Generate()
	orderNo := fmt.Sprintf("RC%s%06d", gtime.Now().Format("YmdHis"), id%1000000)
	payTypeInt := 1
	if payType == "alipay" {
		payTypeInt = 2
	}
	amount := plan[rpc.Amount].Int64()
	giftAmount := plan[rpc.GiftAmount].Int64()

	roc := dao.PlayRechargeOrder.Columns()
	_, err = dao.PlayRechargeOrder.Ctx(ctx).Data(g.Map{
		roc.Id:             id,
		roc.OrderNo:        orderNo,
		roc.MemberId:       memberID,
		roc.RechargePlanId: pid,
		roc.Amount:         amount,
		roc.GiftAmount:     giftAmount,
		roc.PayType:        payTypeInt,
		roc.PayStatus:      0,
		roc.CreatedAt:      gtime.Now(),
		roc.UpdatedAt:      gtime.Now(),
	}).Insert()
	if err != nil {
		return
	}

	orderID = strconv.FormatInt(int64(id), 10)

	devMode := g.Cfg().MustGet(ctx, "pay.devMode").Bool()
	if devMode {
		// 开发模式：直接充值成功，立即给会员加余额
		totalAmount := amount + giftAmount
		e := s.processRechargeSuccess(ctx, int64(id), memberID, totalAmount, orderNo, "mock_trade_no")
		if e != nil {
			// 充值失败不阻断返回，仅记录日志
			g.Log().Errorf(ctx, "devMode 充值处理失败: %v", e)
			payParams = `{"mock":"devMode recharge success","error":"balance add failed"}`
		} else {
			payParams = `{"mock":"devMode recharge success"}`
		}
		return
	}

	// 真实支付模式：调用第三方 SDK 获取支付参数
	switch payType {
	case "wechat":
		client, e := wxpay.New(ctx)
		if e != nil {
			err = gerror.Newf("初始化微信支付失败: %v", e)
			return
		}
		wxParams, e := client.CreateOrder(ctx, orderNo, amount, plan[rpc.Title].String())
		if e != nil {
			err = gerror.Newf("创建微信支付充值订单失败: %v", e)
			return
		}
		paramsBytes, _ := json.Marshal(wxParams)
		payParams = string(paramsBytes)

	case "alipay":
		client, e := alipay.New(ctx)
		if e != nil {
			err = gerror.Newf("初始化支付宝失败: %v", e)
			return
		}
		payURL, e := client.CreateOrder(ctx, orderNo, amount, plan[rpc.Title].String())
		if e != nil {
			err = gerror.Newf("创建支付宝充值订单失败: %v", e)
			return
		}
		payParams = fmt.Sprintf(`{"pay_url":"%s"}`, payURL)
	}

	return
}

// WxNotify 充值微信回调
func (s *sRecharge) WxNotify(ctx context.Context) error {
	devMode := g.Cfg().MustGet(ctx, "pay.devMode").Bool()
	if devMode {
		g.Log().Info(ctx, "RechargeWxNotify: devMode 模式，跳过真实回调处理")
		return nil
	}

	r := g.RequestFromCtx(ctx)
	client, err := wxpay.New(ctx)
	if err != nil {
		return gerror.Newf("初始化微信支付失败: %v", err)
	}

	outTradeNo, err := client.VerifyNotify(ctx, r.Request)
	if err != nil {
		return gerror.Newf("充值微信回调验签失败: %v", err)
	}

	return s.handleWxRechargeSuccess(ctx, outTradeNo, "")
}

// AlipayNotify 充值支付宝回调
func (s *sRecharge) AlipayNotify(ctx context.Context) error {
	devMode := g.Cfg().MustGet(ctx, "pay.devMode").Bool()
	if devMode {
		g.Log().Info(ctx, "RechargeAlipayNotify: devMode 模式，跳过真实回调处理")
		return nil
	}

	r := g.RequestFromCtx(ctx)
	client, err := alipay.New(ctx)
	if err != nil {
		return gerror.Newf("初始化支付宝失败: %v", err)
	}

	postForm := url.Values{}
	for k, v := range r.GetFormMapStrStr() {
		postForm.Set(k, v)
	}

	outTradeNo, err := client.VerifyNotify(ctx, postForm)
	if err != nil {
		return gerror.Newf("充值支付宝回调验签失败: %v", err)
	}

	tradeNo := r.GetForm("trade_no").String()
	return s.handleAlipayRechargeSuccess(ctx, outTradeNo, tradeNo)
}

// ---- 内部方法 ----

// handleWxRechargeSuccess 处理微信充值成功
func (s *sRecharge) handleWxRechargeSuccess(ctx context.Context, outTradeNo, tradeNo string) error {
	roc := dao.PlayRechargeOrder.Columns()
	order, err := dao.PlayRechargeOrder.Ctx(ctx).Where(roc.OrderNo, outTradeNo).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return gerror.Newf("充值订单不存在: %s", outTradeNo)
	}

	memberID := order[roc.MemberId].Int64()
	amount := order[roc.Amount].Int64()
	giftAmount := order[roc.GiftAmount].Int64()
	totalAmount := amount + giftAmount

	return s.processRechargeSuccess(ctx, order[roc.Id].Int64(), memberID, totalAmount, outTradeNo, tradeNo)
}

// handleAlipayRechargeSuccess 处理支付宝充值成功
func (s *sRecharge) handleAlipayRechargeSuccess(ctx context.Context, outTradeNo, tradeNo string) error {
	return s.handleWxRechargeSuccess(ctx, outTradeNo, tradeNo)
}

// processRechargeSuccess 充值成功核心逻辑：更新订单状态 + 给会员加余额（事务）
func (s *sRecharge) processRechargeSuccess(ctx context.Context, orderID, memberID, totalAmount int64, orderNo, tradeNo string) error {
	return dao.PlayRechargeOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		roc := dao.PlayRechargeOrder.Columns()

		// 防重：检查充值订单是否已处理
		order, err := dao.PlayRechargeOrder.Ctx(ctx).Where(roc.Id, orderID).One()
		if err != nil {
			return err
		}
		if order.IsEmpty() {
			return gerror.New("充值订单不存在")
		}
		if order[roc.PayStatus].Int() == 1 {
			// 已充值，幂等处理
			return nil
		}

		// 更新充值订单状态
		_, err = dao.PlayRechargeOrder.Ctx(ctx).Where(roc.Id, orderID).Data(g.Map{
			roc.PayStatus: 1,
			roc.TradeNo:   tradeNo,
			roc.PayAt:     gtime.Now(),
			roc.UpdatedAt: gtime.Now(),
		}).Update()
		if err != nil {
			return err
		}

		// 给会员加余额（充值业务类型 = 1）
		return service.BalanceLogEnhance().AddLog(ctx, &model.AddBalanceLogInput{
			MemberID:     snowflake.JsonInt64(memberID),
			BizType:      1, // 充值
			BizID:        snowflake.JsonInt64(orderID),
			ChangeAmount: totalAmount,
			Remark:       fmt.Sprintf("充值订单%s到账", orderNo),
		})
	})
}

