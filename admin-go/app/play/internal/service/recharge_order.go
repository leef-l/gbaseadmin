package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IRechargeOrder interface {
	Create(ctx context.Context, in *model.RechargeOrderCreateInput) error
	Update(ctx context.Context, in *model.RechargeOrderUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.RechargeOrderDetailOutput, err error)
	List(ctx context.Context, in *model.RechargeOrderListInput) (list []*model.RechargeOrderListOutput, total int, err error)
}

var localRechargeOrder IRechargeOrder

func RechargeOrder() IRechargeOrder {
	return localRechargeOrder
}

func RegisterRechargeOrder(i IRechargeOrder) {
	localRechargeOrder = i
}
