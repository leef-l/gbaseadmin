package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IGoods interface {
	Create(ctx context.Context, in *model.GoodsCreateInput) error
	Update(ctx context.Context, in *model.GoodsUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.GoodsDetailOutput, err error)
	List(ctx context.Context, in *model.GoodsListInput) (list []*model.GoodsListOutput, total int, err error)
}

var localGoods IGoods

func Goods() IGoods {
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
