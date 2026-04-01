package playapi

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Activity = &cActivity{}

type cActivity struct{}

// Join 报名参与活动
func (c *cActivity) Join(ctx context.Context, req *v1.ActivityJoinApiReq) (res *v1.ActivityJoinApiRes, err error) {
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	err = service.PlayapiActivity().Join(ctx, memberID, req.ActivityID)
	return
}

// CompleteStep 完成活动步骤
func (c *cActivity) CompleteStep(ctx context.Context, req *v1.ActivityStepApiReq) (res *v1.ActivityStepApiRes, err error) {
	res = &v1.ActivityStepApiRes{}
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	res.CurrentStep, res.IsCompleted, err = service.PlayapiActivity().CompleteStep(ctx, memberID, req.ActivityID, req.StepID, req.ImageUrl)
	return
}

// ClaimReward 领取奖励
func (c *cActivity) ClaimReward(ctx context.Context, req *v1.ActivityClaimApiReq) (res *v1.ActivityClaimApiRes, err error) {
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	err = service.PlayapiActivity().ClaimReward(ctx, memberID, req.ActivityID, req.RewardID)
	return
}

// MyJoins 我参与的活动列表
func (c *cActivity) MyJoins(ctx context.Context, req *v1.ActivityMyJoinsReq) (res *v1.ActivityMyJoinsRes, err error) {
	res = &v1.ActivityMyJoinsRes{}
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	res.List, res.Total, err = service.PlayapiActivity().MyJoins(ctx, memberID, req.Page, req.PageSize)
	return
}

// Quit 取消报名
func (c *cActivity) Quit(ctx context.Context, req *v1.ActivityQuitApiReq) (res *v1.ActivityQuitApiRes, err error) {
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	err = service.PlayapiActivity().Quit(ctx, memberID, req.ActivityID)
	return
}

var ActivityPublic = &cActivityPublic{}

type cActivityPublic struct{}

// List 活动列表
func (c *cActivityPublic) List(ctx context.Context, req *v1.ActivityListApiReq) (res *v1.ActivityListApiRes, err error) {
	res = &v1.ActivityListApiRes{}
	res.List, res.Total, err = service.PlayapiActivity().List(ctx, req.Page, req.PageSize)
	return
}

// Detail 活动详情
func (c *cActivityPublic) Detail(ctx context.Context, req *v1.ActivityDetailApiReq) (res *v1.ActivityDetailApiRes, err error) {
	// 尝试获取会员ID（未登录为 0）
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	res, err = service.PlayapiActivity().Detail(ctx, req.ActivityID, memberID)
	return
}
