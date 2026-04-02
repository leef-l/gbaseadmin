package activity_step

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var ActivityStep = cActivityStep{}

type cActivityStep struct{}

// Create 创建活动步骤表
func (c *cActivityStep) Create(ctx context.Context, req *v1.ActivityStepCreateReq) (res *v1.ActivityStepCreateRes, err error) {
	err = service.ActivityStep().Create(ctx, &model.ActivityStepCreateInput{
		ActivityID:  req.ActivityID,
		StepNum:     req.StepNum,
		Title:       req.Title,
		StepType:    req.StepType,
		ExampleText: req.ExampleText,
		DescContent: req.DescContent,
		StepImage:   req.StepImage,
		IsRequired:  req.IsRequired,
		Sort:        req.Sort,
	})
	return
}

// Update 更新活动步骤表
func (c *cActivityStep) Update(ctx context.Context, req *v1.ActivityStepUpdateReq) (res *v1.ActivityStepUpdateRes, err error) {
	err = service.ActivityStep().Update(ctx, &model.ActivityStepUpdateInput{
		ID:          req.ID,
		ActivityID:  req.ActivityID,
		StepNum:     req.StepNum,
		Title:       req.Title,
		StepType:    req.StepType,
		ExampleText: req.ExampleText,
		DescContent: req.DescContent,
		StepImage:   req.StepImage,
		IsRequired:  req.IsRequired,
		Sort:        req.Sort,
	})
	return
}

// Delete 删除活动步骤表
func (c *cActivityStep) Delete(ctx context.Context, req *v1.ActivityStepDeleteReq) (res *v1.ActivityStepDeleteRes, err error) {
	err = service.ActivityStep().Delete(ctx, req.ID)
	return
}

// Detail 获取活动步骤表详情
func (c *cActivityStep) Detail(ctx context.Context, req *v1.ActivityStepDetailReq) (res *v1.ActivityStepDetailRes, err error) {
	res = &v1.ActivityStepDetailRes{}
	res.ActivityStepDetailOutput, err = service.ActivityStep().Detail(ctx, req.ID)
	return
}

// List 获取活动步骤表列表
func (c *cActivityStep) List(ctx context.Context, req *v1.ActivityStepListReq) (res *v1.ActivityStepListRes, err error) {
	res = &v1.ActivityStepListRes{}
	res.List, res.Total, err = service.ActivityStep().List(ctx, &model.ActivityStepListInput{
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
		ActivityID: req.ActivityID,
	})
	return
}

