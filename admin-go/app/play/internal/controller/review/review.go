package review

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Review = cReview{}

type cReview struct{}

// Create 创建评价表
func (c *cReview) Create(ctx context.Context, req *v1.ReviewCreateReq) (res *v1.ReviewCreateRes, err error) {
	err = service.Review().Create(ctx, &model.ReviewCreateInput{
		OrderID: req.OrderID,
		MemberID: req.MemberID,
		CoachID: req.CoachID,
		Score: req.Score,
		ReviewContent: req.ReviewContent,
		ReviewImage: req.ReviewImage,
		ReplyContent: req.ReplyContent,
		ReplyAt: req.ReplyAt,
		IsAnonymous: req.IsAnonymous,
		Status: req.Status,
	})
	return
}

// Update 更新评价表
func (c *cReview) Update(ctx context.Context, req *v1.ReviewUpdateReq) (res *v1.ReviewUpdateRes, err error) {
	err = service.Review().Update(ctx, &model.ReviewUpdateInput{
		ID: req.ID,
		OrderID: req.OrderID,
		MemberID: req.MemberID,
		CoachID: req.CoachID,
		Score: req.Score,
		ReviewContent: req.ReviewContent,
		ReviewImage: req.ReviewImage,
		ReplyContent: req.ReplyContent,
		ReplyAt: req.ReplyAt,
		IsAnonymous: req.IsAnonymous,
		Status: req.Status,
	})
	return
}

// Delete 删除评价表
func (c *cReview) Delete(ctx context.Context, req *v1.ReviewDeleteReq) (res *v1.ReviewDeleteRes, err error) {
	err = service.Review().Delete(ctx, req.ID)
	return
}

// Detail 获取评价表详情
func (c *cReview) Detail(ctx context.Context, req *v1.ReviewDetailReq) (res *v1.ReviewDetailRes, err error) {
	res = &v1.ReviewDetailRes{}
	res.ReviewDetailOutput, err = service.Review().Detail(ctx, req.ID)
	return
}

// List 获取评价表列表
func (c *cReview) List(ctx context.Context, req *v1.ReviewListReq) (res *v1.ReviewListRes, err error) {
	res = &v1.ReviewListRes{}
	res.List, res.Total, err = service.Review().List(ctx, &model.ReviewListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		IsAnonymous: req.IsAnonymous,
		Status: req.Status,
	})
	return
}

