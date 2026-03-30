package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IRechargePlan interface {
	Create(ctx context.Context, in *model.RechargePlanCreateInput) error
	Update(ctx context.Context, in *model.RechargePlanUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.RechargePlanDetailOutput, err error)
	List(ctx context.Context, in *model.RechargePlanListInput) (list []*model.RechargePlanListOutput, total int, err error)
}

var localRechargePlan IRechargePlan

func RechargePlan() IRechargePlan {
	return localRechargePlan
}

func RegisterRechargePlan(i IRechargePlan) {
	localRechargePlan = i
}
