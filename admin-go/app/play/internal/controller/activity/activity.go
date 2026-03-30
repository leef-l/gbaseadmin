package activity

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Activity = cActivity{}

type cActivity struct{}

// Create 创建活动表
func (c *cActivity) Create(ctx context.Context, req *v1.ActivityCreateReq) (res *v1.ActivityCreateRes, err error) {
	err = service.Activity().Create(ctx, &model.ActivityCreateInput{
		Title: req.Title,
		CoverImage: req.CoverImage,
		DescContent: req.DescContent,
		Type: req.Type,
		ConditionType: req.ConditionType,
		ConditionValue: req.ConditionValue,
		IsAutoReward: req.IsAutoReward,
		StartAt: req.StartAt,
		EndAt: req.EndAt,
		MaxNum: req.MaxNum,
		JoinNum: req.JoinNum,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新活动表
func (c *cActivity) Update(ctx context.Context, req *v1.ActivityUpdateReq) (res *v1.ActivityUpdateRes, err error) {
	err = service.Activity().Update(ctx, &model.ActivityUpdateInput{
		ID: req.ID,
		Title: req.Title,
		CoverImage: req.CoverImage,
		DescContent: req.DescContent,
		Type: req.Type,
		ConditionType: req.ConditionType,
		ConditionValue: req.ConditionValue,
		IsAutoReward: req.IsAutoReward,
		StartAt: req.StartAt,
		EndAt: req.EndAt,
		MaxNum: req.MaxNum,
		JoinNum: req.JoinNum,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除活动表
func (c *cActivity) Delete(ctx context.Context, req *v1.ActivityDeleteReq) (res *v1.ActivityDeleteRes, err error) {
	err = service.Activity().Delete(ctx, req.ID)
	return
}

// Detail 获取活动表详情
func (c *cActivity) Detail(ctx context.Context, req *v1.ActivityDetailReq) (res *v1.ActivityDetailRes, err error) {
	res = &v1.ActivityDetailRes{}
	res.ActivityDetailOutput, err = service.Activity().Detail(ctx, req.ID)
	return
}

// List 获取活动表列表
func (c *cActivity) List(ctx context.Context, req *v1.ActivityListReq) (res *v1.ActivityListRes, err error) {
	res = &v1.ActivityListRes{}
	res.List, res.Total, err = service.Activity().List(ctx, &model.ActivityListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Type: req.Type,
		ConditionType: req.ConditionType,
		IsAutoReward: req.IsAutoReward,
		Status: req.Status,
	})
	return
}

