package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type ICoachLevel interface {
	Create(ctx context.Context, in *model.CoachLevelCreateInput) error
	Update(ctx context.Context, in *model.CoachLevelUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CoachLevelDetailOutput, err error)
	List(ctx context.Context, in *model.CoachLevelListInput) (list []*model.CoachLevelListOutput, total int, err error)
}

var localCoachLevel ICoachLevel

func CoachLevel() ICoachLevel {
	return localCoachLevel
}

func RegisterCoachLevel(i ICoachLevel) {
	localCoachLevel = i
}
