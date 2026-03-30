package order

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Order = cOrder{}

type cOrder struct{}

// Create 创建订单表
func (c *cOrder) Create(ctx context.Context, req *v1.OrderCreateReq) (res *v1.OrderCreateRes, err error) {
	err = service.Order().Create(ctx, &model.OrderCreateInput{
		OrderNo: req.OrderNo,
		MemberID: req.MemberID,
		CoachID: req.CoachID,
		ShopID: req.ShopID,
		GoodsID: req.GoodsID,
		GoodsTitle: req.GoodsTitle,
		GoodsPrice: req.GoodsPrice,
		Quantity: req.Quantity,
		TotalAmount: req.TotalAmount,
		DiscountAmount: req.DiscountAmount,
		CouponAmount: req.CouponAmount,
		PayAmount: req.PayAmount,
		CouponMemberID: req.CouponMemberID,
		PayType: req.PayType,
		OrderStatus: req.OrderStatus,
		PayAt: req.PayAt,
		StartAt: req.StartAt,
		FinishAt: req.FinishAt,
		CancelAt: req.CancelAt,
		CancelReason: req.CancelReason,
		Remark: req.Remark,
	})
	return
}

// Update 更新订单表
func (c *cOrder) Update(ctx context.Context, req *v1.OrderUpdateReq) (res *v1.OrderUpdateRes, err error) {
	err = service.Order().Update(ctx, &model.OrderUpdateInput{
		ID: req.ID,
		OrderNo: req.OrderNo,
		MemberID: req.MemberID,
		CoachID: req.CoachID,
		ShopID: req.ShopID,
		GoodsID: req.GoodsID,
		GoodsTitle: req.GoodsTitle,
		GoodsPrice: req.GoodsPrice,
		Quantity: req.Quantity,
		TotalAmount: req.TotalAmount,
		DiscountAmount: req.DiscountAmount,
		CouponAmount: req.CouponAmount,
		PayAmount: req.PayAmount,
		CouponMemberID: req.CouponMemberID,
		PayType: req.PayType,
		OrderStatus: req.OrderStatus,
		PayAt: req.PayAt,
		StartAt: req.StartAt,
		FinishAt: req.FinishAt,
		CancelAt: req.CancelAt,
		CancelReason: req.CancelReason,
		Remark: req.Remark,
	})
	return
}

// Delete 删除订单表
func (c *cOrder) Delete(ctx context.Context, req *v1.OrderDeleteReq) (res *v1.OrderDeleteRes, err error) {
	err = service.Order().Delete(ctx, req.ID)
	return
}

// Detail 获取订单表详情
func (c *cOrder) Detail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error) {
	res = &v1.OrderDetailRes{}
	res.OrderDetailOutput, err = service.Order().Detail(ctx, req.ID)
	return
}

// List 获取订单表列表
func (c *cOrder) List(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	res = &v1.OrderListRes{}
	res.List, res.Total, err = service.Order().List(ctx, &model.OrderListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		PayType: req.PayType,
		OrderStatus: req.OrderStatus,
	})
	return
}


