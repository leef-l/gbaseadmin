package playapi

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var GoodsPublic = &cGoodsPublic{}

type cGoodsPublic struct{}

func (c *cGoodsPublic) GoodsList(ctx context.Context, req *v1.GoodsListReq) (res *v1.GoodsListRes, err error) {
	res = &v1.GoodsListRes{}
	res.List, res.Total, err = service.PlayapiGoods().List(ctx, req)
	return
}

func (c *cGoodsPublic) GoodsDetail(ctx context.Context, req *v1.GoodsDetailReq) (res *v1.GoodsDetailRes, err error) {
	return service.PlayapiGoods().Detail(ctx, req)
}

func (c *cGoodsPublic) CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error) {
	res = &v1.CategoryListRes{}
	res.List, err = service.PlayapiGoods().CategoryList(ctx, req)
	return
}
