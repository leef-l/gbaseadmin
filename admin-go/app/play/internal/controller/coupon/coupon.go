package coupon

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Coupon = cCoupon{}

type cCoupon struct{}

// Create 创建优惠券模板表
func (c *cCoupon) Create(ctx context.Context, req *v1.CouponCreateReq) (res *v1.CouponCreateRes, err error) {
	err = service.Coupon().Create(ctx, &model.CouponCreateInput{
		Title: req.Title,
		Type: req.Type,
		IsNewMember: req.IsNewMember,
		FaceValue: req.FaceValue,
		MinAmount: req.MinAmount,
		TotalNum: req.TotalNum,
		UsedNum: req.UsedNum,
		ClaimNum: req.ClaimNum,
		PerLimit: req.PerLimit,
		ValidStartAt: req.ValidStartAt,
		ValidEndAt: req.ValidEndAt,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新优惠券模板表
func (c *cCoupon) Update(ctx context.Context, req *v1.CouponUpdateReq) (res *v1.CouponUpdateRes, err error) {
	err = service.Coupon().Update(ctx, &model.CouponUpdateInput{
		ID: req.ID,
		Title: req.Title,
		Type: req.Type,
		IsNewMember: req.IsNewMember,
		FaceValue: req.FaceValue,
		MinAmount: req.MinAmount,
		TotalNum: req.TotalNum,
		UsedNum: req.UsedNum,
		ClaimNum: req.ClaimNum,
		PerLimit: req.PerLimit,
		ValidStartAt: req.ValidStartAt,
		ValidEndAt: req.ValidEndAt,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除优惠券模板表
func (c *cCoupon) Delete(ctx context.Context, req *v1.CouponDeleteReq) (res *v1.CouponDeleteRes, err error) {
	err = service.Coupon().Delete(ctx, req.ID)
	return
}

// Detail 获取优惠券模板表详情
func (c *cCoupon) Detail(ctx context.Context, req *v1.CouponDetailReq) (res *v1.CouponDetailRes, err error) {
	res = &v1.CouponDetailRes{}
	res.CouponDetailOutput, err = service.Coupon().Detail(ctx, req.ID)
	return
}

// List 获取优惠券模板表列表
func (c *cCoupon) List(ctx context.Context, req *v1.CouponListReq) (res *v1.CouponListRes, err error) {
	res = &v1.CouponListRes{}
	res.List, res.Total, err = service.Coupon().List(ctx, &model.CouponListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Type: req.Type,
		IsNewMember: req.IsNewMember,
		Status: req.Status,
	})
	return
}

