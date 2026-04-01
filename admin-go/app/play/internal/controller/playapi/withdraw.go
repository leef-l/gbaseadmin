package playapi

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Withdraw = &cWithdraw{}

type cWithdraw struct{}

// CoachWithdraw 申请提现
func (c *cWithdraw) CoachWithdraw(ctx context.Context, req *v1.CoachWithdrawApiReq) (res *v1.CoachWithdrawApiRes, err error) {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	memberID := r.GetCtxVar("jwt_member_id").Int64()

	res = &v1.CoachWithdrawApiRes{}
	res.WithdrawId, err = service.PlayapiWithdraw().Withdraw(ctx, coachID, memberID, req.Amount)
	return
}

// CoachWithdrawList 提现记录列表
func (c *cWithdraw) CoachWithdrawList(ctx context.Context, req *v1.CoachWithdrawListApiReq) (res *v1.CoachWithdrawListApiRes, err error) {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()

	res = &v1.CoachWithdrawListApiRes{}
	res.List, res.Total, err = service.PlayapiWithdraw().WithdrawList(ctx, coachID, req)
	return
}
