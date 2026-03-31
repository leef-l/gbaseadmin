package playapi

import (
	"context"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Review = &cReview{}

type cReview struct{}

func (c *cReview) Create(ctx context.Context, req *v1.ReviewCreateReq) (res *v1.ReviewCreateRes, err error) {
	err = service.PlayapiReview().Create(ctx, req)
	return
}

var ReviewPublic = &cReviewPublic{}

type cReviewPublic struct{}

func (c *cReviewPublic) List(ctx context.Context, req *v1.ReviewListReq) (res *v1.ReviewListRes, err error) {
	res = &v1.ReviewListRes{}
	res.List, res.Total, err = service.PlayapiReview().List(ctx, req)
	return
}

var ReviewCoach = &cReviewCoach{}

type cReviewCoach struct{}

func (c *cReviewCoach) Reply(ctx context.Context, req *v1.ReviewReplyReq) (res *v1.ReviewReplyRes, err error) {
	err = service.PlayapiReview().Reply(ctx, req)
	return
}
