package coach

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Coach = cCoach{}

type cCoach struct{}

// Create 创建陪玩师表
func (c *cCoach) Create(ctx context.Context, req *v1.CoachCreateReq) (res *v1.CoachCreateRes, err error) {
	err = service.Coach().Create(ctx, &model.CoachCreateInput{
		MemberID: req.MemberID,
		CoachLevelID: req.CoachLevelID,
		ShopID: req.ShopID,
		RealName: req.RealName,
		Intro: req.Intro,
		CoverImage: req.CoverImage,
		TotalOrders: req.TotalOrders,
		TotalScore: req.TotalScore,
		ScoreNum: req.ScoreNum,
		IncomeTotal: req.IncomeTotal,
		IncomeBalance: req.IncomeBalance,
		IsOnline: req.IsOnline,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新陪玩师表
func (c *cCoach) Update(ctx context.Context, req *v1.CoachUpdateReq) (res *v1.CoachUpdateRes, err error) {
	err = service.Coach().Update(ctx, &model.CoachUpdateInput{
		ID: req.ID,
		MemberID: req.MemberID,
		CoachLevelID: req.CoachLevelID,
		ShopID: req.ShopID,
		RealName: req.RealName,
		Intro: req.Intro,
		CoverImage: req.CoverImage,
		TotalOrders: req.TotalOrders,
		TotalScore: req.TotalScore,
		ScoreNum: req.ScoreNum,
		IncomeTotal: req.IncomeTotal,
		IncomeBalance: req.IncomeBalance,
		IsOnline: req.IsOnline,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除陪玩师表
func (c *cCoach) Delete(ctx context.Context, req *v1.CoachDeleteReq) (res *v1.CoachDeleteRes, err error) {
	err = service.Coach().Delete(ctx, req.ID)
	return
}

// Detail 获取陪玩师表详情
func (c *cCoach) Detail(ctx context.Context, req *v1.CoachDetailReq) (res *v1.CoachDetailRes, err error) {
	res = &v1.CoachDetailRes{}
	res.CoachDetailOutput, err = service.Coach().Detail(ctx, req.ID)
	return
}

// List 获取陪玩师表列表
func (c *cCoach) List(ctx context.Context, req *v1.CoachListReq) (res *v1.CoachListRes, err error) {
	res = &v1.CoachListRes{}
	res.List, res.Total, err = service.Coach().List(ctx, &model.CoachListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsOnline: req.IsOnline,
		Status: req.Status,
	})
	return
}

