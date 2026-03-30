package member

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Member = cMember{}

type cMember struct{}

// Create 创建会员表
func (c *cMember) Create(ctx context.Context, req *v1.MemberCreateReq) (res *v1.MemberCreateRes, err error) {
	err = service.Member().Create(ctx, &model.MemberCreateInput{
		Phone: req.Phone,
		Password: req.Password,
		Nickname: req.Nickname,
		Avatar: req.Avatar,
		Gender: req.Gender,
		MemberLevelID: req.MemberLevelID,
		Exp: req.Exp,
		Balance: req.Balance,
		IsCoach: req.IsCoach,
		Status: req.Status,
		LastLoginAt: req.LastLoginAt,
	})
	return
}

// Update 更新会员表
func (c *cMember) Update(ctx context.Context, req *v1.MemberUpdateReq) (res *v1.MemberUpdateRes, err error) {
	err = service.Member().Update(ctx, &model.MemberUpdateInput{
		ID: req.ID,
		Phone: req.Phone,
		Password: req.Password,
		Nickname: req.Nickname,
		Avatar: req.Avatar,
		Gender: req.Gender,
		MemberLevelID: req.MemberLevelID,
		Exp: req.Exp,
		Balance: req.Balance,
		IsCoach: req.IsCoach,
		Status: req.Status,
		LastLoginAt: req.LastLoginAt,
	})
	return
}

// Delete 删除会员表
func (c *cMember) Delete(ctx context.Context, req *v1.MemberDeleteReq) (res *v1.MemberDeleteRes, err error) {
	err = service.Member().Delete(ctx, req.ID)
	return
}

// Detail 获取会员表详情
func (c *cMember) Detail(ctx context.Context, req *v1.MemberDetailReq) (res *v1.MemberDetailRes, err error) {
	res = &v1.MemberDetailRes{}
	res.MemberDetailOutput, err = service.Member().Detail(ctx, req.ID)
	return
}

// List 获取会员表列表
func (c *cMember) List(ctx context.Context, req *v1.MemberListReq) (res *v1.MemberListRes, err error) {
	res = &v1.MemberListRes{}
	res.List, res.Total, err = service.Member().List(ctx, &model.MemberListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Gender: req.Gender,
		IsCoach: req.IsCoach,
		Status: req.Status,
	})
	return
}

