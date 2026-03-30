package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type ICoupon interface {
	Create(ctx context.Context, in *model.CouponCreateInput) error
	Update(ctx context.Context, in *model.CouponUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CouponDetailOutput, err error)
	List(ctx context.Context, in *model.CouponListInput) (list []*model.CouponListOutput, total int, err error)
}

var localCoupon ICoupon

func Coupon() ICoupon {
	return localCoupon
}

func RegisterCoupon(i ICoupon) {
	localCoupon = i
}
