package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IActivityStepLog interface {
	Create(ctx context.Context, in *model.ActivityStepLogCreateInput) error
	Update(ctx context.Context, in *model.ActivityStepLogUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityStepLogDetailOutput, err error)
	List(ctx context.Context, in *model.ActivityStepLogListInput) (list []*model.ActivityStepLogListOutput, total int, err error)
}

var localActivityStepLog IActivityStepLog

func ActivityStepLog() IActivityStepLog {
	return localActivityStepLog
}

func RegisterActivityStepLog(i IActivityStepLog) {
	localActivityStepLog = i
}
