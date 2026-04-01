package playapi

import (
	"context"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var BannerPublic = &cBannerPublic{}

type cBannerPublic struct{}

func (c *cBannerPublic) BannerList(ctx context.Context, req *v1.BannerListReq) (res *v1.BannerListRes, err error) {
	res = &v1.BannerListRes{}
	res.List, err = service.PlayapiBanner().List(ctx, req)
	return
}
