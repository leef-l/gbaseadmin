package activity_join

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var ActivityJoin = cActivityJoin{}

type cActivityJoin struct{}

// Create 创建活动参与记录表
func (c *cActivityJoin) Create(ctx context.Context, req *v1.ActivityJoinCreateReq) (res *v1.ActivityJoinCreateRes, err error) {
	err = service.ActivityJoin().Create(ctx, &model.ActivityJoinCreateInput{
		ActivityID: req.ActivityID,
		MemberID: req.MemberID,
		JoinStatus: req.JoinStatus,
		CurrentStep: req.CurrentStep,
		FinishAt: req.FinishAt,
		RewardAt: req.RewardAt,
		Remark: req.Remark,
	})
	return
}

// Update 更新活动参与记录表
func (c *cActivityJoin) Update(ctx context.Context, req *v1.ActivityJoinUpdateReq) (res *v1.ActivityJoinUpdateRes, err error) {
	err = service.ActivityJoin().Update(ctx, &model.ActivityJoinUpdateInput{
		ID: req.ID,
		ActivityID: req.ActivityID,
		MemberID: req.MemberID,
		JoinStatus: req.JoinStatus,
		CurrentStep: req.CurrentStep,
		FinishAt: req.FinishAt,
		RewardAt: req.RewardAt,
		Remark: req.Remark,
	})
	return
}

// Delete 删除活动参与记录表
func (c *cActivityJoin) Delete(ctx context.Context, req *v1.ActivityJoinDeleteReq) (res *v1.ActivityJoinDeleteRes, err error) {
	err = service.ActivityJoin().Delete(ctx, req.ID)
	return
}

// Detail 获取活动参与记录表详情
func (c *cActivityJoin) Detail(ctx context.Context, req *v1.ActivityJoinDetailReq) (res *v1.ActivityJoinDetailRes, err error) {
	res = &v1.ActivityJoinDetailRes{}
	res.ActivityJoinDetailOutput, err = service.ActivityJoin().Detail(ctx, req.ID)
	return
}

// List 获取活动参与记录表列表
func (c *cActivityJoin) List(ctx context.Context, req *v1.ActivityJoinListReq) (res *v1.ActivityJoinListRes, err error) {
	res = &v1.ActivityJoinListRes{}
	res.List, res.Total, err = service.ActivityJoin().List(ctx, &model.ActivityJoinListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		JoinStatus: req.JoinStatus,
	})
	return
}

