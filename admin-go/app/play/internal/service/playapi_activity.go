package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiActivity interface {
	List(ctx context.Context, page, pageSize int) (list []v1.ActivityListApiItem, total int, err error)
	Detail(ctx context.Context, activityID string) (out *v1.ActivityDetailApiRes, err error)
	Join(ctx context.Context, memberID int64, activityID string) error
	CompleteStep(ctx context.Context, memberID int64, activityID, stepID string) (currentStep int, isCompleted bool, err error)
	ClaimReward(ctx context.Context, memberID int64, activityID, rewardID string) error
	MyJoins(ctx context.Context, memberID int64, page, pageSize int) (list []v1.ActivityMyJoinsItem, total int, err error)
}

var localPlayapiActivity IPlayapiActivity

func PlayapiActivity() IPlayapiActivity {
	return localPlayapiActivity
}

func RegisterPlayapiActivity(i IPlayapiActivity) {
	localPlayapiActivity = i
}
