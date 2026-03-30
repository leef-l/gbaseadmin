package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type ICouponMember interface {
	Create(ctx context.Context, in *model.CouponMemberCreateInput) error
	Update(ctx context.Context, in *model.CouponMemberUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CouponMemberDetailOutput, err error)
	List(ctx context.Context, in *model.CouponMemberListInput) (list []*model.CouponMemberListOutput, total int, err error)
}

var localCouponMember ICouponMember

func CouponMember() ICouponMember {
	return localCouponMember
}

func RegisterCouponMember(i ICouponMember) {
	localCouponMember = i
}
