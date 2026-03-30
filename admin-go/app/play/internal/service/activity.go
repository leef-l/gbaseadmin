package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IActivity interface {
	Create(ctx context.Context, in *model.ActivityCreateInput) error
	Update(ctx context.Context, in *model.ActivityUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityDetailOutput, err error)
	List(ctx context.Context, in *model.ActivityListInput) (list []*model.ActivityListOutput, total int, err error)
}

var localActivity IActivity

func Activity() IActivity {
	return localActivity
}

func RegisterActivity(i IActivity) {
	localActivity = i
}
