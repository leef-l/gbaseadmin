package playapi

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Coupon = &cCoupon{}

type cCoupon struct{}

// Receive 领取优惠券
func (c *cCoupon) Receive(ctx context.Context, req *v1.CouponReceiveReq) (res *v1.CouponReceiveRes, err error) {
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	err = service.PlayapiCoupon().Receive(ctx, memberID, req.CouponID)
	return
}

// Mine 我的优惠券
func (c *cCoupon) Mine(ctx context.Context, req *v1.CouponMineReq) (res *v1.CouponMineRes, err error) {
	res = &v1.CouponMineRes{}
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	res.List, res.Total, err = service.PlayapiCoupon().Mine(ctx, memberID, req.Status, req.Page, req.PageSize)
	return
}

// Usable 下单可用优惠券
func (c *cCoupon) Usable(ctx context.Context, req *v1.CouponUsableApiReq) (res *v1.CouponUsableApiRes, err error) {
	res = &v1.CouponUsableApiRes{}
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	res.List, err = service.PlayapiCoupon().Usable(ctx, memberID, req.OrderAmount)
	return
}

var CouponPublic = &cCouponPublic{}

type cCouponPublic struct{}

// Available 可领取优惠券列表
func (c *cCouponPublic) Available(ctx context.Context, req *v1.CouponAvailableReq) (res *v1.CouponAvailableRes, err error) {
	res = &v1.CouponAvailableRes{}
	res.List, res.Total, err = service.PlayapiCoupon().Available(ctx, req.Page, req.PageSize)
	return
}
