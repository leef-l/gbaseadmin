package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type ICoach interface {
	Create(ctx context.Context, in *model.CoachCreateInput) error
	Update(ctx context.Context, in *model.CoachUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CoachDetailOutput, err error)
	List(ctx context.Context, in *model.CoachListInput) (list []*model.CoachListOutput, total int, err error)
}

var localCoach ICoach

func Coach() ICoach {
	return localCoach
}

func RegisterCoach(i ICoach) {
	localCoach = i
}
