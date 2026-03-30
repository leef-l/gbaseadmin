package payment

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Payment = cPayment{}

type cPayment struct{}

// Create 创建支付记录表
func (c *cPayment) Create(ctx context.Context, req *v1.PaymentCreateReq) (res *v1.PaymentCreateRes, err error) {
	err = service.Payment().Create(ctx, &model.PaymentCreateInput{
		OrderID: req.OrderID,
		MemberID: req.MemberID,
		PaymentNo: req.PaymentNo,
		TradeNo: req.TradeNo,
		PayType: req.PayType,
		PayAmount: req.PayAmount,
		PayStatus: req.PayStatus,
		PayAt: req.PayAt,
		RefundAt: req.RefundAt,
		RefundAmount: req.RefundAmount,
		CallbackContent: req.CallbackContent,
	})
	return
}

// Update 更新支付记录表
func (c *cPayment) Update(ctx context.Context, req *v1.PaymentUpdateReq) (res *v1.PaymentUpdateRes, err error) {
	err = service.Payment().Update(ctx, &model.PaymentUpdateInput{
		ID: req.ID,
		OrderID: req.OrderID,
		MemberID: req.MemberID,
		PaymentNo: req.PaymentNo,
		TradeNo: req.TradeNo,
		PayType: req.PayType,
		PayAmount: req.PayAmount,
		PayStatus: req.PayStatus,
		PayAt: req.PayAt,
		RefundAt: req.RefundAt,
		RefundAmount: req.RefundAmount,
		CallbackContent: req.CallbackContent,
	})
	return
}

// Delete 删除支付记录表
func (c *cPayment) Delete(ctx context.Context, req *v1.PaymentDeleteReq) (res *v1.PaymentDeleteRes, err error) {
	err = service.Payment().Delete(ctx, req.ID)
	return
}

// Detail 获取支付记录表详情
func (c *cPayment) Detail(ctx context.Context, req *v1.PaymentDetailReq) (res *v1.PaymentDetailRes, err error) {
	res = &v1.PaymentDetailRes{}
	res.PaymentDetailOutput, err = service.Payment().Detail(ctx, req.ID)
	return
}

// List 获取支付记录表列表
func (c *cPayment) List(ctx context.Context, req *v1.PaymentListReq) (res *v1.PaymentListRes, err error) {
	res = &v1.PaymentListRes{}
	res.List, res.Total, err = service.Payment().List(ctx, &model.PaymentListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		PayType: req.PayType,
		PayStatus: req.PayStatus,
	})
	return
}

