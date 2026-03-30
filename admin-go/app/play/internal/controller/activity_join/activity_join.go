package activity_join

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var ActivityJoin = cActivityJoin{}

type cActivityJoin struct{}

// Create 创建æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨
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

// Update 更新æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨
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

// Delete 删除æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨
func (c *cActivityJoin) Delete(ctx context.Context, req *v1.ActivityJoinDeleteReq) (res *v1.ActivityJoinDeleteRes, err error) {
	err = service.ActivityJoin().Delete(ctx, req.ID)
	return
}

// Detail 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨详情
func (c *cActivityJoin) Detail(ctx context.Context, req *v1.ActivityJoinDetailReq) (res *v1.ActivityJoinDetailRes, err error) {
	res = &v1.ActivityJoinDetailRes{}
	res.ActivityJoinDetailOutput, err = service.ActivityJoin().Detail(ctx, req.ID)
	return
}

// List 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨列表
func (c *cActivityJoin) List(ctx context.Context, req *v1.ActivityJoinListReq) (res *v1.ActivityJoinListRes, err error) {
	res = &v1.ActivityJoinListRes{}
	res.List, res.Total, err = service.ActivityJoin().List(ctx, &model.ActivityJoinListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		JoinStatus: req.JoinStatus,
	})
	return
}

