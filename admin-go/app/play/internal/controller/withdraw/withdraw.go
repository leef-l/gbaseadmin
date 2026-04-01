package withdraw

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Withdraw = cWithdraw{}

type cWithdraw struct{}

// Create 创建陪玩师提现记录
func (c *cWithdraw) Create(ctx context.Context, req *v1.WithdrawCreateReq) (res *v1.WithdrawCreateRes, err error) {
	err = service.Withdraw().Create(ctx, &model.WithdrawCreateInput{
		CoachID: req.CoachID,
		MemberID: req.MemberID,
		Amount: req.Amount,
		Status: req.Status,
		Reason: req.Reason,
		AuditedAt: req.AuditedAt,
	})
	return
}

// Update 更新陪玩师提现记录
func (c *cWithdraw) Update(ctx context.Context, req *v1.WithdrawUpdateReq) (res *v1.WithdrawUpdateRes, err error) {
	err = service.Withdraw().Update(ctx, &model.WithdrawUpdateInput{
		ID: req.ID,
		CoachID: req.CoachID,
		MemberID: req.MemberID,
		Amount: req.Amount,
		Status: req.Status,
		Reason: req.Reason,
		AuditedAt: req.AuditedAt,
	})
	return
}

// Delete 删除陪玩师提现记录
func (c *cWithdraw) Delete(ctx context.Context, req *v1.WithdrawDeleteReq) (res *v1.WithdrawDeleteRes, err error) {
	err = service.Withdraw().Delete(ctx, req.ID)
	return
}

// BatchDelete 批量删除陪玩师提现记录
func (c *cWithdraw) BatchDelete(ctx context.Context, req *v1.WithdrawBatchDeleteReq) (res *v1.WithdrawBatchDeleteRes, err error) {
	err = service.Withdraw().BatchDelete(ctx, req.IDs)
	return
}

// Detail 获取陪玩师提现记录详情
func (c *cWithdraw) Detail(ctx context.Context, req *v1.WithdrawDetailReq) (res *v1.WithdrawDetailRes, err error) {
	res = &v1.WithdrawDetailRes{}
	res.WithdrawDetailOutput, err = service.Withdraw().Detail(ctx, req.ID)
	return
}

// List 获取陪玩师提现记录列表
func (c *cWithdraw) List(ctx context.Context, req *v1.WithdrawListReq) (res *v1.WithdrawListRes, err error) {
	res = &v1.WithdrawListRes{}
	res.List, res.Total, err = service.Withdraw().List(ctx, &model.WithdrawListInput{
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
		OrderBy:   req.OrderBy,
		OrderDir:  req.OrderDir,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return
}
// Export 导出陪玩师提现记录
func (c *cWithdraw) Export(ctx context.Context, req *v1.WithdrawExportReq) (res *v1.WithdrawExportRes, err error) {
	list, err := service.Withdraw().Export(ctx, &model.WithdrawListInput{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		return
	}
	// CSV 导出
	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", "text/csv; charset=utf-8")
	r.Response.Header().Set("Content-Disposition", `attachment; filename="withdraw.csv"`)
	r.Response.Write("\xEF\xBB\xBF") // UTF-8 BOM
	// 表头
	r.Response.Writeln("陪玩师ID,会员ID,提现金额,状态 0=待审核 1=已打款 2=已拒绝,拒绝原因,审核时间,创建时间")
	// 数据行
	for _, item := range list {
		r.Response.Writefln("%v,%v,%v,%v,%v,%v,%v",
			item.CoachRealName,
			 item.MemberNickname,
			 item.Amount,
			 item.Status,
			 item.Reason,
			 item.AuditedAt,
			item.CreatedAt,
		)
	}
	return
}

