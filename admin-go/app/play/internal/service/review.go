package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IReview interface {
	Create(ctx context.Context, in *model.ReviewCreateInput) error
	Update(ctx context.Context, in *model.ReviewUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ReviewDetailOutput, err error)
	List(ctx context.Context, in *model.ReviewListInput) (list []*model.ReviewListOutput, total int, err error)
}

var localReview IReview

func Review() IReview {
	return localReview
}

func RegisterReview(i IReview) {
	localReview = i
}
