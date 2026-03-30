package balance_log

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var BalanceLog = cBalanceLog{}

type cBalanceLog struct{}

// Create 创建ä½™é¢æµæ°´è¡¨
func (c *cBalanceLog) Create(ctx context.Context, req *v1.BalanceLogCreateReq) (res *v1.BalanceLogCreateRes, err error) {
	err = service.BalanceLog().Create(ctx, &model.BalanceLogCreateInput{
		MemberID: req.MemberID,
		BizType: req.BizType,
		BizID: req.BizID,
		ChangeAmount: req.ChangeAmount,
		BeforeBalance: req.BeforeBalance,
		AfterBalance: req.AfterBalance,
		Remark: req.Remark,
	})
	return
}

// Update 更新ä½™é¢æµæ°´è¡¨
func (c *cBalanceLog) Update(ctx context.Context, req *v1.BalanceLogUpdateReq) (res *v1.BalanceLogUpdateRes, err error) {
	err = service.BalanceLog().Update(ctx, &model.BalanceLogUpdateInput{
		ID: req.ID,
		MemberID: req.MemberID,
		BizType: req.BizType,
		BizID: req.BizID,
		ChangeAmount: req.ChangeAmount,
		BeforeBalance: req.BeforeBalance,
		AfterBalance: req.AfterBalance,
		Remark: req.Remark,
	})
	return
}

// Delete 删除ä½™é¢æµæ°´è¡¨
func (c *cBalanceLog) Delete(ctx context.Context, req *v1.BalanceLogDeleteReq) (res *v1.BalanceLogDeleteRes, err error) {
	err = service.BalanceLog().Delete(ctx, req.ID)
	return
}

// Detail 获取ä½™é¢æµæ°´è¡¨详情
func (c *cBalanceLog) Detail(ctx context.Context, req *v1.BalanceLogDetailReq) (res *v1.BalanceLogDetailRes, err error) {
	res = &v1.BalanceLogDetailRes{}
	res.BalanceLogDetailOutput, err = service.BalanceLog().Detail(ctx, req.ID)
	return
}

// List 获取ä½™é¢æµæ°´è¡¨列表
func (c *cBalanceLog) List(ctx context.Context, req *v1.BalanceLogListReq) (res *v1.BalanceLogListRes, err error) {
	res = &v1.BalanceLogListRes{}
	res.List, res.Total, err = service.BalanceLog().List(ctx, &model.BalanceLogListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		BizType: req.BizType,
	})
	return
}

