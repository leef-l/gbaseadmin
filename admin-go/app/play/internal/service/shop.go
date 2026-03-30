package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IShop interface {
	Create(ctx context.Context, in *model.ShopCreateInput) error
	Update(ctx context.Context, in *model.ShopUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ShopDetailOutput, err error)
	List(ctx context.Context, in *model.ShopListInput) (list []*model.ShopListOutput, total int, err error)
}

var localShop IShop

func Shop() IShop {
	return localShop
}

func RegisterShop(i IShop) {
	localShop = i
}
