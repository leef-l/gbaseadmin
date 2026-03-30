package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiOrder interface {
	Create(ctx context.Context, req *v1.OrderCreateReq) (res *v1.OrderCreateRes, err error)
	List(ctx context.Context, req *v1.OrderListReq) (list []v1.OrderListItem, total int, err error)
	Detail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error)
	Cancel(ctx context.Context, req *v1.OrderCancelReq) error
	Refund(ctx context.Context, req *v1.OrderRefundReq) error
	Accept(ctx context.Context, req *v1.OrderAcceptReq) error
	Complete(ctx context.Context, req *v1.OrderCompleteReq) error
}

var localPlayapiOrder IPlayapiOrder

func PlayapiOrder() IPlayapiOrder {
	return localPlayapiOrder
}

func RegisterPlayapiOrder(i IPlayapiOrder) {
	localPlayapiOrder = i
}
