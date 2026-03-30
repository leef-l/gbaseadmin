package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IBalanceLog interface {
	Create(ctx context.Context, in *model.BalanceLogCreateInput) error
	Update(ctx context.Context, in *model.BalanceLogUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.BalanceLogDetailOutput, err error)
	List(ctx context.Context, in *model.BalanceLogListInput) (list []*model.BalanceLogListOutput, total int, err error)
}

var localBalanceLog IBalanceLog

func BalanceLog() IBalanceLog {
	return localBalanceLog
}

func RegisterBalanceLog(i IBalanceLog) {
	localBalanceLog = i
}
