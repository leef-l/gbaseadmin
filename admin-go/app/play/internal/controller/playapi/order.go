package playapi

import (
	"context"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Order = &cOrder{}

type cOrder struct{}

func (c *cOrder) Create(ctx context.Context, req *v1.OrderCreateReq) (res *v1.OrderCreateRes, err error) {
	return service.PlayapiOrder().Create(ctx, req)
}

func (c *cOrder) List(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	res = &v1.OrderListRes{}
	res.List, res.Total, err = service.PlayapiOrder().List(ctx, req)
	return
}

func (c *cOrder) Detail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error) {
	return service.PlayapiOrder().Detail(ctx, req)
}

func (c *cOrder) Cancel(ctx context.Context, req *v1.OrderCancelReq) (res *v1.OrderCancelRes, err error) {
	err = service.PlayapiOrder().Cancel(ctx, req)
	return
}

func (c *cOrder) Refund(ctx context.Context, req *v1.OrderRefundReq) (res *v1.OrderRefundRes, err error) {
	err = service.PlayapiOrder().Refund(ctx, req)
	return
}

var OrderCoach = &cOrderCoach{}

type cOrderCoach struct{}

func (c *cOrderCoach) Accept(ctx context.Context, req *v1.OrderAcceptReq) (res *v1.OrderAcceptRes, err error) {
	err = service.PlayapiOrder().Accept(ctx, req)
	return
}

func (c *cOrderCoach) Complete(ctx context.Context, req *v1.OrderCompleteReq) (res *v1.OrderCompleteRes, err error) {
	err = service.PlayapiOrder().Complete(ctx, req)
	return
}
