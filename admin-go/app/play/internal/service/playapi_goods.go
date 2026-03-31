package service

import (
	"context"

	v1 "gbaseadmin/app/play/api/playapi/v1"
)

type IPlayapiGoods interface {
	List(ctx context.Context, req *v1.GoodsListReq) (list []v1.GoodsListItem, total int, err error)
	Detail(ctx context.Context, req *v1.GoodsDetailReq) (res *v1.GoodsDetailRes, err error)
	CategoryList(ctx context.Context, req *v1.CategoryListReq) (list []v1.CategoryTreeItem, err error)
}

var localPlayapiGoods IPlayapiGoods

func PlayapiGoods() IPlayapiGoods {
	return localPlayapiGoods
}

func RegisterPlayapiGoods(i IPlayapiGoods) {
	localPlayapiGoods = i
}
