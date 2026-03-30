package member_level

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var MemberLevel = cMemberLevel{}

type cMemberLevel struct{}

// Create 创建会员等级表
func (c *cMemberLevel) Create(ctx context.Context, req *v1.MemberLevelCreateReq) (res *v1.MemberLevelCreateRes, err error) {
	err = service.MemberLevel().Create(ctx, &model.MemberLevelCreateInput{
		Title: req.Title,
		Level: req.Level,
		Icon: req.Icon,
		MinExp: req.MinExp,
		Discount: req.Discount,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新会员等级表
func (c *cMemberLevel) Update(ctx context.Context, req *v1.MemberLevelUpdateReq) (res *v1.MemberLevelUpdateRes, err error) {
	err = service.MemberLevel().Update(ctx, &model.MemberLevelUpdateInput{
		ID: req.ID,
		Title: req.Title,
		Level: req.Level,
		Icon: req.Icon,
		MinExp: req.MinExp,
		Discount: req.Discount,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除会员等级表
func (c *cMemberLevel) Delete(ctx context.Context, req *v1.MemberLevelDeleteReq) (res *v1.MemberLevelDeleteRes, err error) {
	err = service.MemberLevel().Delete(ctx, req.ID)
	return
}

// Detail 获取会员等级表详情
func (c *cMemberLevel) Detail(ctx context.Context, req *v1.MemberLevelDetailReq) (res *v1.MemberLevelDetailRes, err error) {
	res = &v1.MemberLevelDetailRes{}
	res.MemberLevelDetailOutput, err = service.MemberLevel().Detail(ctx, req.ID)
	return
}

// List 获取会员等级表列表
func (c *cMemberLevel) List(ctx context.Context, req *v1.MemberLevelListReq) (res *v1.MemberLevelListRes, err error) {
	res = &v1.MemberLevelListRes{}
	res.List, res.Total, err = service.MemberLevel().List(ctx, &model.MemberLevelListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Level: req.Level,
		Status: req.Status,
	})
	return
}

