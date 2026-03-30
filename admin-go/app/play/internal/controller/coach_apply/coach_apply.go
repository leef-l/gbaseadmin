package coach_apply

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var CoachApply = cCoachApply{}

type cCoachApply struct{}

// Create 创建陪玩师申请表
func (c *cCoachApply) Create(ctx context.Context, req *v1.CoachApplyCreateReq) (res *v1.CoachApplyCreateRes, err error) {
	err = service.CoachApply().Create(ctx, &model.CoachApplyCreateInput{
		MemberID: req.MemberID,
		RealName: req.RealName,
		IDCard: req.IDCard,
		IDCardFrontImage: req.IDCardFrontImage,
		IDCardBackImage: req.IDCardBackImage,
		SkillDesc: req.SkillDesc,
		AuditStatus: req.AuditStatus,
		AuditRemark: req.AuditRemark,
		AuditAt: req.AuditAt,
	})
	return
}

// Update 更新陪玩师申请表
func (c *cCoachApply) Update(ctx context.Context, req *v1.CoachApplyUpdateReq) (res *v1.CoachApplyUpdateRes, err error) {
	err = service.CoachApply().Update(ctx, &model.CoachApplyUpdateInput{
		ID: req.ID,
		MemberID: req.MemberID,
		RealName: req.RealName,
		IDCard: req.IDCard,
		IDCardFrontImage: req.IDCardFrontImage,
		IDCardBackImage: req.IDCardBackImage,
		SkillDesc: req.SkillDesc,
		AuditStatus: req.AuditStatus,
		AuditRemark: req.AuditRemark,
		AuditAt: req.AuditAt,
	})
	return
}

// Delete 删除陪玩师申请表
func (c *cCoachApply) Delete(ctx context.Context, req *v1.CoachApplyDeleteReq) (res *v1.CoachApplyDeleteRes, err error) {
	err = service.CoachApply().Delete(ctx, req.ID)
	return
}

// Detail 获取陪玩师申请表详情
func (c *cCoachApply) Detail(ctx context.Context, req *v1.CoachApplyDetailReq) (res *v1.CoachApplyDetailRes, err error) {
	res = &v1.CoachApplyDetailRes{}
	res.CoachApplyDetailOutput, err = service.CoachApply().Detail(ctx, req.ID)
	return
}

// List 获取陪玩师申请表列表
func (c *cCoachApply) List(ctx context.Context, req *v1.CoachApplyListReq) (res *v1.CoachApplyListRes, err error) {
	res = &v1.CoachApplyListRes{}
	res.List, res.Total, err = service.CoachApply().List(ctx, &model.CoachApplyListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		AuditStatus: req.AuditStatus,
	})
	return
}


