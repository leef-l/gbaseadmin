package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiRecharge interface {
	Plans(ctx context.Context) (list []v1.RechargePlanItem, err error)
	Create(ctx context.Context, memberID int64, planID string, payType string) (orderID string, payParams string, err error)
	WxNotify(ctx context.Context) error
	AlipayNotify(ctx context.Context) error
}

var localPlayapiRecharge IPlayapiRecharge

func PlayapiRecharge() IPlayapiRecharge {
	return localPlayapiRecharge
}

func RegisterPlayapiRecharge(i IPlayapiRecharge) {
	localPlayapiRecharge = i
}
