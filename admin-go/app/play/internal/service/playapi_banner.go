package service

import (
	"context"

	v1 "gbaseadmin/app/play/api/playapi/v1"
)

type IPlayapiBanner interface {
	List(ctx context.Context, req *v1.BannerListReq) (list []v1.BannerListItem, err error)
}

var localPlayapiBanner IPlayapiBanner

func PlayapiBanner() IPlayapiBanner {
	return localPlayapiBanner
}

func RegisterPlayapiBanner(i IPlayapiBanner) {
	localPlayapiBanner = i
}
