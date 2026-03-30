package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IActivityStep interface {
	Create(ctx context.Context, in *model.ActivityStepCreateInput) error
	Update(ctx context.Context, in *model.ActivityStepUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityStepDetailOutput, err error)
	List(ctx context.Context, in *model.ActivityStepListInput) (list []*model.ActivityStepListOutput, total int, err error)
}

var localActivityStep IActivityStep

func ActivityStep() IActivityStep {
	return localActivityStep
}

func RegisterActivityStep(i IActivityStep) {
	localActivityStep = i
}
