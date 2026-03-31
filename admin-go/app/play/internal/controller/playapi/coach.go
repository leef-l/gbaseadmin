package playapi

import (
	"context"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var CoachPublic = &cCoachPublic{}

type cCoachPublic struct{}

func (c *cCoachPublic) CoachList(ctx context.Context, req *v1.CoachListReq) (res *v1.CoachListRes, err error) {
	res = &v1.CoachListRes{}
	res.List, res.Total, err = service.PlayapiCoach().List(ctx, req)
	return
}

func (c *cCoachPublic) CoachDetail(ctx context.Context, req *v1.CoachDetailReq) (res *v1.CoachDetailRes, err error) {
	return service.PlayapiCoach().Detail(ctx, req)
}

var CoachApply = &cCoachApply{}

type cCoachApply struct{}

func (c *cCoachApply) Apply(ctx context.Context, req *v1.CoachApplyReq) (res *v1.CoachApplyRes, err error) {
	err = service.PlayapiCoach().Apply(ctx, req)
	return
}

func (c *cCoachApply) ApplyStatus(ctx context.Context, req *v1.CoachApplyStatusReq) (res *v1.CoachApplyStatusRes, err error) {
	return service.PlayapiCoach().ApplyStatus(ctx, req)
}

var CoachWork = &cCoachWork{}

type cCoachWork struct{}

func (c *cCoachWork) Online(ctx context.Context, req *v1.CoachOnlineReq) (res *v1.CoachOnlineRes, err error) {
	err = service.PlayapiCoach().SetOnline(ctx, req)
	return
}

func (c *cCoachWork) MyGoods(ctx context.Context, req *v1.CoachMyGoodsReq) (res *v1.CoachMyGoodsRes, err error) {
	res = &v1.CoachMyGoodsRes{}
	res.List, res.Total, err = service.PlayapiCoach().MyGoods(ctx, req)
	return
}

func (c *cCoachWork) GoodsCreate(ctx context.Context, req *v1.CoachGoodsCreateReq) (res *v1.CoachGoodsCreateRes, err error) {
	err = service.PlayapiCoach().GoodsCreate(ctx, req)
	return
}

func (c *cCoachWork) GoodsUpdate(ctx context.Context, req *v1.CoachGoodsUpdateReq) (res *v1.CoachGoodsUpdateRes, err error) {
	err = service.PlayapiCoach().GoodsUpdate(ctx, req)
	return
}

func (c *cCoachWork) GoodsStatus(ctx context.Context, req *v1.CoachGoodsStatusReq) (res *v1.CoachGoodsStatusRes, err error) {
	err = service.PlayapiCoach().GoodsStatus(ctx, req)
	return
}

func (c *cCoachWork) Income(ctx context.Context, req *v1.CoachIncomeReq) (res *v1.CoachIncomeRes, err error) {
	return service.PlayapiCoach().Income(ctx, req)
}

func (c *cCoachWork) Orders(ctx context.Context, req *v1.CoachOrdersReq) (res *v1.CoachOrdersRes, err error) {
	res = &v1.CoachOrdersRes{}
	res.List, res.Total, err = service.PlayapiCoach().Orders(ctx, req)
	return
}
