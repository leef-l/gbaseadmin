package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type ICoachApply interface {
	Create(ctx context.Context, in *model.CoachApplyCreateInput) error
	Update(ctx context.Context, in *model.CoachApplyUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CoachApplyDetailOutput, err error)
	List(ctx context.Context, in *model.CoachApplyListInput) (list []*model.CoachApplyListOutput, total int, err error)
}

var localCoachApply ICoachApply

func CoachApply() ICoachApply {
	return localCoachApply
}

func RegisterCoachApply(i ICoachApply) {
	localCoachApply = i
}
