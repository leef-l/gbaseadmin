package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IOrder interface {
	Create(ctx context.Context, in *model.OrderCreateInput) error
	Update(ctx context.Context, in *model.OrderUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.OrderDetailOutput, err error)
	List(ctx context.Context, in *model.OrderListInput) (list []*model.OrderListOutput, total int, err error)
}

var localOrder IOrder

func Order() IOrder {
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
