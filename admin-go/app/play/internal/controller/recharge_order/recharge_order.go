package recharge_order

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var RechargeOrder = cRechargeOrder{}

type cRechargeOrder struct{}

// Create 创建充值订单表
func (c *cRechargeOrder) Create(ctx context.Context, req *v1.RechargeOrderCreateReq) (res *v1.RechargeOrderCreateRes, err error) {
	err = service.RechargeOrder().Create(ctx, &model.RechargeOrderCreateInput{
		OrderNo: req.OrderNo,
		MemberID: req.MemberID,
		RechargePlanID: req.RechargePlanID,
		Amount: req.Amount,
		GiftAmount: req.GiftAmount,
		PayType: req.PayType,
		TradeNo: req.TradeNo,
		PayStatus: req.PayStatus,
		PayAt: req.PayAt,
	})
	return
}

// Update 更新充值订单表
func (c *cRechargeOrder) Update(ctx context.Context, req *v1.RechargeOrderUpdateReq) (res *v1.RechargeOrderUpdateRes, err error) {
	err = service.RechargeOrder().Update(ctx, &model.RechargeOrderUpdateInput{
		ID: req.ID,
		OrderNo: req.OrderNo,
		MemberID: req.MemberID,
		RechargePlanID: req.RechargePlanID,
		Amount: req.Amount,
		GiftAmount: req.GiftAmount,
		PayType: req.PayType,
		TradeNo: req.TradeNo,
		PayStatus: req.PayStatus,
		PayAt: req.PayAt,
	})
	return
}

// Delete 删除充值订单表
func (c *cRechargeOrder) Delete(ctx context.Context, req *v1.RechargeOrderDeleteReq) (res *v1.RechargeOrderDeleteRes, err error) {
	err = service.RechargeOrder().Delete(ctx, req.ID)
	return
}

// Detail 获取充值订单表详情
func (c *cRechargeOrder) Detail(ctx context.Context, req *v1.RechargeOrderDetailReq) (res *v1.RechargeOrderDetailRes, err error) {
	res = &v1.RechargeOrderDetailRes{}
	res.RechargeOrderDetailOutput, err = service.RechargeOrder().Detail(ctx, req.ID)
	return
}

// List 获取充值订单表列表
func (c *cRechargeOrder) List(ctx context.Context, req *v1.RechargeOrderListReq) (res *v1.RechargeOrderListRes, err error) {
	res = &v1.RechargeOrderListRes{}
	res.List, res.Total, err = service.RechargeOrder().List(ctx, &model.RechargeOrderListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		PayType: req.PayType,
		PayStatus: req.PayStatus,
	})
	return
}

