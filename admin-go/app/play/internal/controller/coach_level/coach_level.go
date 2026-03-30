package coach_level

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var CoachLevel = cCoachLevel{}

type cCoachLevel struct{}

// Create 创建陪玩师等级表
func (c *cCoachLevel) Create(ctx context.Context, req *v1.CoachLevelCreateReq) (res *v1.CoachLevelCreateRes, err error) {
	err = service.CoachLevel().Create(ctx, &model.CoachLevelCreateInput{
		Title: req.Title,
		Level: req.Level,
		Icon: req.Icon,
		MinOrders: req.MinOrders,
		MinScore: req.MinScore,
		CommissionRate: req.CommissionRate,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新陪玩师等级表
func (c *cCoachLevel) Update(ctx context.Context, req *v1.CoachLevelUpdateReq) (res *v1.CoachLevelUpdateRes, err error) {
	err = service.CoachLevel().Update(ctx, &model.CoachLevelUpdateInput{
		ID: req.ID,
		Title: req.Title,
		Level: req.Level,
		Icon: req.Icon,
		MinOrders: req.MinOrders,
		MinScore: req.MinScore,
		CommissionRate: req.CommissionRate,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除陪玩师等级表
func (c *cCoachLevel) Delete(ctx context.Context, req *v1.CoachLevelDeleteReq) (res *v1.CoachLevelDeleteRes, err error) {
	err = service.CoachLevel().Delete(ctx, req.ID)
	return
}

// Detail 获取陪玩师等级表详情
func (c *cCoachLevel) Detail(ctx context.Context, req *v1.CoachLevelDetailReq) (res *v1.CoachLevelDetailRes, err error) {
	res = &v1.CoachLevelDetailRes{}
	res.CoachLevelDetailOutput, err = service.CoachLevel().Detail(ctx, req.ID)
	return
}

// List 获取陪玩师等级表列表
func (c *cCoachLevel) List(ctx context.Context, req *v1.CoachLevelListReq) (res *v1.CoachLevelListRes, err error) {
	res = &v1.CoachLevelListRes{}
	res.List, res.Total, err = service.CoachLevel().List(ctx, &model.CoachLevelListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Level: req.Level,
		Status: req.Status,
	})
	return
}

