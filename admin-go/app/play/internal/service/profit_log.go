package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IProfitLog interface {
	Create(ctx context.Context, in *model.ProfitLogCreateInput) error
	Update(ctx context.Context, in *model.ProfitLogUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ProfitLogDetailOutput, err error)
	List(ctx context.Context, in *model.ProfitLogListInput) (list []*model.ProfitLogListOutput, total int, err error)
}

var localProfitLog IProfitLog

func ProfitLog() IProfitLog {
	return localProfitLog
}

func RegisterProfitLog(i IProfitLog) {
	localProfitLog = i
}
