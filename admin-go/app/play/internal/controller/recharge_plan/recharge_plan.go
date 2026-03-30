package recharge_plan

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var RechargePlan = cRechargePlan{}

type cRechargePlan struct{}

// Create 创建充值方案表
func (c *cRechargePlan) Create(ctx context.Context, req *v1.RechargePlanCreateReq) (res *v1.RechargePlanCreateRes, err error) {
	err = service.RechargePlan().Create(ctx, &model.RechargePlanCreateInput{
		Title: req.Title,
		Amount: req.Amount,
		GiftAmount: req.GiftAmount,
		CoverImage: req.CoverImage,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新充值方案表
func (c *cRechargePlan) Update(ctx context.Context, req *v1.RechargePlanUpdateReq) (res *v1.RechargePlanUpdateRes, err error) {
	err = service.RechargePlan().Update(ctx, &model.RechargePlanUpdateInput{
		ID: req.ID,
		Title: req.Title,
		Amount: req.Amount,
		GiftAmount: req.GiftAmount,
		CoverImage: req.CoverImage,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除充值方案表
func (c *cRechargePlan) Delete(ctx context.Context, req *v1.RechargePlanDeleteReq) (res *v1.RechargePlanDeleteRes, err error) {
	err = service.RechargePlan().Delete(ctx, req.ID)
	return
}

// Detail 获取充值方案表详情
func (c *cRechargePlan) Detail(ctx context.Context, req *v1.RechargePlanDetailReq) (res *v1.RechargePlanDetailRes, err error) {
	res = &v1.RechargePlanDetailRes{}
	res.RechargePlanDetailOutput, err = service.RechargePlan().Detail(ctx, req.ID)
	return
}

// List 获取充值方案表列表
func (c *cRechargePlan) List(ctx context.Context, req *v1.RechargePlanListReq) (res *v1.RechargePlanListRes, err error) {
	res = &v1.RechargePlanListRes{}
	res.List, res.Total, err = service.RechargePlan().List(ctx, &model.RechargePlanListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Status: req.Status,
	})
	return
}

