package coupon_member

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var CouponMember = cCouponMember{}

type cCouponMember struct{}

// Create 创建ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨
func (c *cCouponMember) Create(ctx context.Context, req *v1.CouponMemberCreateReq) (res *v1.CouponMemberCreateRes, err error) {
	err = service.CouponMember().Create(ctx, &model.CouponMemberCreateInput{
		CouponID: req.CouponID,
		MemberID: req.MemberID,
		OrderID: req.OrderID,
		UseStatus: req.UseStatus,
		ClaimAt: req.ClaimAt,
		UseAt: req.UseAt,
		ExpireAt: req.ExpireAt,
	})
	return
}

// Update 更新ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨
func (c *cCouponMember) Update(ctx context.Context, req *v1.CouponMemberUpdateReq) (res *v1.CouponMemberUpdateRes, err error) {
	err = service.CouponMember().Update(ctx, &model.CouponMemberUpdateInput{
		ID: req.ID,
		CouponID: req.CouponID,
		MemberID: req.MemberID,
		OrderID: req.OrderID,
		UseStatus: req.UseStatus,
		ClaimAt: req.ClaimAt,
		UseAt: req.UseAt,
		ExpireAt: req.ExpireAt,
	})
	return
}

// Delete 删除ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨
func (c *cCouponMember) Delete(ctx context.Context, req *v1.CouponMemberDeleteReq) (res *v1.CouponMemberDeleteRes, err error) {
	err = service.CouponMember().Delete(ctx, req.ID)
	return
}

// Detail 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨详情
func (c *cCouponMember) Detail(ctx context.Context, req *v1.CouponMemberDetailReq) (res *v1.CouponMemberDetailRes, err error) {
	res = &v1.CouponMemberDetailRes{}
	res.CouponMemberDetailOutput, err = service.CouponMember().Detail(ctx, req.ID)
	return
}

// List 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨列表
func (c *cCouponMember) List(ctx context.Context, req *v1.CouponMemberListReq) (res *v1.CouponMemberListRes, err error) {
	res = &v1.CouponMemberListRes{}
	res.List, res.Total, err = service.CouponMember().List(ctx, &model.CouponMemberListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		UseStatus: req.UseStatus,
	})
	return
}

