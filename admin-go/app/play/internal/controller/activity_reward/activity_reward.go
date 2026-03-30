package activity_reward

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var ActivityReward = cActivityReward{}

type cActivityReward struct{}

// Create 创建活动奖励表
func (c *cActivityReward) Create(ctx context.Context, req *v1.ActivityRewardCreateReq) (res *v1.ActivityRewardCreateRes, err error) {
	err = service.ActivityReward().Create(ctx, &model.ActivityRewardCreateInput{
		ActivityID: req.ActivityID,
		RewardType: req.RewardType,
		RewardValue: req.RewardValue,
		RewardName: req.RewardName,
		Sort: req.Sort,
	})
	return
}

// Update 更新活动奖励表
func (c *cActivityReward) Update(ctx context.Context, req *v1.ActivityRewardUpdateReq) (res *v1.ActivityRewardUpdateRes, err error) {
	err = service.ActivityReward().Update(ctx, &model.ActivityRewardUpdateInput{
		ID: req.ID,
		ActivityID: req.ActivityID,
		RewardType: req.RewardType,
		RewardValue: req.RewardValue,
		RewardName: req.RewardName,
		Sort: req.Sort,
	})
	return
}

// Delete 删除活动奖励表
func (c *cActivityReward) Delete(ctx context.Context, req *v1.ActivityRewardDeleteReq) (res *v1.ActivityRewardDeleteRes, err error) {
	err = service.ActivityReward().Delete(ctx, req.ID)
	return
}

// Detail 获取活动奖励表详情
func (c *cActivityReward) Detail(ctx context.Context, req *v1.ActivityRewardDetailReq) (res *v1.ActivityRewardDetailRes, err error) {
	res = &v1.ActivityRewardDetailRes{}
	res.ActivityRewardDetailOutput, err = service.ActivityReward().Detail(ctx, req.ID)
	return
}

// List 获取活动奖励表列表
func (c *cActivityReward) List(ctx context.Context, req *v1.ActivityRewardListReq) (res *v1.ActivityRewardListRes, err error) {
	res = &v1.ActivityRewardListRes{}
	res.List, res.Total, err = service.ActivityReward().List(ctx, &model.ActivityRewardListInput{
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
		ActivityID: req.ActivityID,
		RewardType: req.RewardType,
	})
	return
}

