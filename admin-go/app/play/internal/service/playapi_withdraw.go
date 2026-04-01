package service

import (
	"context"

	v1 "gbaseadmin/app/play/api/playapi/v1"
)

type IPlayapiWithdraw interface {
	Withdraw(ctx context.Context, coachID int64, memberID int64, amount int64) (withdrawID string, err error)
	WithdrawList(ctx context.Context, coachID int64, req *v1.CoachWithdrawListApiReq) (list []v1.CoachWithdrawItem, total int, err error)
}

var localPlayapiWithdraw IPlayapiWithdraw

func PlayapiWithdraw() IPlayapiWithdraw {
	return localPlayapiWithdraw
}

func RegisterPlayapiWithdraw(i IPlayapiWithdraw) {
	localPlayapiWithdraw = i
}
