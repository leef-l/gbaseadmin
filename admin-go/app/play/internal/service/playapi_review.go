package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiReview interface {
	Create(ctx context.Context, req *v1.ReviewCreateReq) error
	List(ctx context.Context, req *v1.ReviewListReq) (list []v1.ReviewListItem, total int, err error)
	Reply(ctx context.Context, req *v1.ReviewReplyReq) error
}

var localPlayapiReview IPlayapiReview

func PlayapiReview() IPlayapiReview {
	return localPlayapiReview
}

func RegisterPlayapiReview(i IPlayapiReview) {
	localPlayapiReview = i
}
