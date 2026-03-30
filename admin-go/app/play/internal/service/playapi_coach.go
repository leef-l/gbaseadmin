package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiCoach interface {
	List(ctx context.Context, req *v1.CoachListReq) (list []v1.CoachListItem, total int, err error)
	Detail(ctx context.Context, req *v1.CoachDetailReq) (res *v1.CoachDetailRes, err error)
	Apply(ctx context.Context, req *v1.CoachApplyReq) error
	ApplyStatus(ctx context.Context, req *v1.CoachApplyStatusReq) (res *v1.CoachApplyStatusRes, err error)
	SetOnline(ctx context.Context, req *v1.CoachOnlineReq) error
	MyGoods(ctx context.Context, req *v1.CoachMyGoodsReq) (list []v1.CoachGoodsItem, total int, err error)
	GoodsCreate(ctx context.Context, req *v1.CoachGoodsCreateReq) error
	GoodsUpdate(ctx context.Context, req *v1.CoachGoodsUpdateReq) error
	GoodsStatus(ctx context.Context, req *v1.CoachGoodsStatusReq) error
	Income(ctx context.Context, req *v1.CoachIncomeReq) (res *v1.CoachIncomeRes, err error)
	Orders(ctx context.Context, req *v1.CoachOrdersReq) (list []v1.OrderListItem, total int, err error)
}

var localPlayapiCoach IPlayapiCoach

func PlayapiCoach() IPlayapiCoach {
	return localPlayapiCoach
}

func RegisterPlayapiCoach(i IPlayapiCoach) {
	localPlayapiCoach = i
}
