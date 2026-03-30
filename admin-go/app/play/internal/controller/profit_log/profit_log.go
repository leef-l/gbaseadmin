package profit_log

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var ProfitLog = cProfitLog{}

type cProfitLog struct{}

// Create 创建利润分成流水表
func (c *cProfitLog) Create(ctx context.Context, req *v1.ProfitLogCreateReq) (res *v1.ProfitLogCreateRes, err error) {
	err = service.ProfitLog().Create(ctx, &model.ProfitLogCreateInput{
		OrderID: req.OrderID,
		OrderNo: req.OrderNo,
		PayAmount: req.PayAmount,
		CoachID: req.CoachID,
		ShopID: req.ShopID,
		PlatformRate: req.PlatformRate,
		PlatformAmount: req.PlatformAmount,
		ShopRate: req.ShopRate,
		ShopAmount: req.ShopAmount,
		CoachAmount: req.CoachAmount,
		SettleStatus: req.SettleStatus,
		SettleAt: req.SettleAt,
	})
	return
}

// Update 更新利润分成流水表
func (c *cProfitLog) Update(ctx context.Context, req *v1.ProfitLogUpdateReq) (res *v1.ProfitLogUpdateRes, err error) {
	err = service.ProfitLog().Update(ctx, &model.ProfitLogUpdateInput{
		ID: req.ID,
		OrderID: req.OrderID,
		OrderNo: req.OrderNo,
		PayAmount: req.PayAmount,
		CoachID: req.CoachID,
		ShopID: req.ShopID,
		PlatformRate: req.PlatformRate,
		PlatformAmount: req.PlatformAmount,
		ShopRate: req.ShopRate,
		ShopAmount: req.ShopAmount,
		CoachAmount: req.CoachAmount,
		SettleStatus: req.SettleStatus,
		SettleAt: req.SettleAt,
	})
	return
}

// Delete 删除利润分成流水表
func (c *cProfitLog) Delete(ctx context.Context, req *v1.ProfitLogDeleteReq) (res *v1.ProfitLogDeleteRes, err error) {
	err = service.ProfitLog().Delete(ctx, req.ID)
	return
}

// Detail 获取利润分成流水表详情
func (c *cProfitLog) Detail(ctx context.Context, req *v1.ProfitLogDetailReq) (res *v1.ProfitLogDetailRes, err error) {
	res = &v1.ProfitLogDetailRes{}
	res.ProfitLogDetailOutput, err = service.ProfitLog().Detail(ctx, req.ID)
	return
}

// List 获取利润分成流水表列表
func (c *cProfitLog) List(ctx context.Context, req *v1.ProfitLogListReq) (res *v1.ProfitLogListRes, err error) {
	res = &v1.ProfitLogListRes{}
	res.List, res.Total, err = service.ProfitLog().List(ctx, &model.ProfitLogListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		SettleStatus: req.SettleStatus,
	})
	return
}

