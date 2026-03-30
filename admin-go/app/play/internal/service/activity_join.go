package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IActivityJoin interface {
	Create(ctx context.Context, in *model.ActivityJoinCreateInput) error
	Update(ctx context.Context, in *model.ActivityJoinUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityJoinDetailOutput, err error)
	List(ctx context.Context, in *model.ActivityJoinListInput) (list []*model.ActivityJoinListOutput, total int, err error)
}

var localActivityJoin IActivityJoin

func ActivityJoin() IActivityJoin {
	return localActivityJoin
}

func RegisterActivityJoin(i IActivityJoin) {
	localActivityJoin = i
}
