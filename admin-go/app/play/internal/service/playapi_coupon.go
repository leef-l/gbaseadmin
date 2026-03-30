package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiCoupon interface {
	Available(ctx context.Context, page, pageSize int) (list []v1.CouponAvailableItem, total int, err error)
	Receive(ctx context.Context, memberID int64, couponID string) error
	Mine(ctx context.Context, memberID int64, status *int, page, pageSize int) (list []v1.CouponMineItem, total int, err error)
	Usable(ctx context.Context, memberID int64, orderAmount int64) (list []v1.CouponMineItem, err error)
}

var localPlayapiCoupon IPlayapiCoupon

func PlayapiCoupon() IPlayapiCoupon {
	return localPlayapiCoupon
}

func RegisterPlayapiCoupon(i IPlayapiCoupon) {
	localPlayapiCoupon = i
}
