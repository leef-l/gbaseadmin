package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IActivityReward interface {
	Create(ctx context.Context, in *model.ActivityRewardCreateInput) error
	Update(ctx context.Context, in *model.ActivityRewardUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityRewardDetailOutput, err error)
	List(ctx context.Context, in *model.ActivityRewardListInput) (list []*model.ActivityRewardListOutput, total int, err error)
}

var localActivityReward IActivityReward

func ActivityReward() IActivityReward {
	return localActivityReward
}

func RegisterActivityReward(i IActivityReward) {
	localActivityReward = i
}
