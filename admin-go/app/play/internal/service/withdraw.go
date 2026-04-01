package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IWithdraw interface {
	Create(ctx context.Context, in *model.WithdrawCreateInput) error
	Update(ctx context.Context, in *model.WithdrawUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	BatchDelete(ctx context.Context, ids []snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.WithdrawDetailOutput, err error)
	List(ctx context.Context, in *model.WithdrawListInput) (list []*model.WithdrawListOutput, total int, err error)
	Export(ctx context.Context, in *model.WithdrawListInput) (list []*model.WithdrawListOutput, err error)
}

var localWithdraw IWithdraw

func Withdraw() IWithdraw {
	return localWithdraw
}

func RegisterWithdraw(i IWithdraw) {
	localWithdraw = i
}
