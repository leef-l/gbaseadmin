package activity_step_log

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var ActivityStepLog = cActivityStepLog{}

type cActivityStepLog struct{}

// Create 创建活动步骤提交记录
func (c *cActivityStepLog) Create(ctx context.Context, req *v1.ActivityStepLogCreateReq) (res *v1.ActivityStepLogCreateRes, err error) {
	err = service.ActivityStepLog().Create(ctx, &model.ActivityStepLogCreateInput{
		ActivityID: req.ActivityID,
		StepID: req.StepID,
		JoinID: req.JoinID,
		MemberID: req.MemberID,
		StepType: req.StepType,
		SubmitText: req.SubmitText,
		SubmitImage: req.SubmitImage,
		AuditStatus: req.AuditStatus,
		AuditRemark: req.AuditRemark,
		AuditBy: req.AuditBy,
		AuditAt: req.AuditAt,
	})
	return
}

// Update 更新活动步骤提交记录
func (c *cActivityStepLog) Update(ctx context.Context, req *v1.ActivityStepLogUpdateReq) (res *v1.ActivityStepLogUpdateRes, err error) {
	err = service.ActivityStepLog().Update(ctx, &model.ActivityStepLogUpdateInput{
		ID: req.ID,
		ActivityID: req.ActivityID,
		StepID: req.StepID,
		JoinID: req.JoinID,
		MemberID: req.MemberID,
		StepType: req.StepType,
		SubmitText: req.SubmitText,
		SubmitImage: req.SubmitImage,
		AuditStatus: req.AuditStatus,
		AuditRemark: req.AuditRemark,
		AuditBy: req.AuditBy,
		AuditAt: req.AuditAt,
	})
	return
}

// Delete 删除活动步骤提交记录
func (c *cActivityStepLog) Delete(ctx context.Context, req *v1.ActivityStepLogDeleteReq) (res *v1.ActivityStepLogDeleteRes, err error) {
	err = service.ActivityStepLog().Delete(ctx, req.ID)
	return
}

// Detail 获取活动步骤提交记录详情
func (c *cActivityStepLog) Detail(ctx context.Context, req *v1.ActivityStepLogDetailReq) (res *v1.ActivityStepLogDetailRes, err error) {
	res = &v1.ActivityStepLogDetailRes{}
	res.ActivityStepLogDetailOutput, err = service.ActivityStepLog().Detail(ctx, req.ID)
	return
}

// List 获取活动步骤提交记录列表
func (c *cActivityStepLog) List(ctx context.Context, req *v1.ActivityStepLogListReq) (res *v1.ActivityStepLogListRes, err error) {
	res = &v1.ActivityStepLogListRes{}
	res.List, res.Total, err = service.ActivityStepLog().List(ctx, &model.ActivityStepLogListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		StepType: req.StepType,
		AuditStatus: req.AuditStatus,
	})
	return
}

