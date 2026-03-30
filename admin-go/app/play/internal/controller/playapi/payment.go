package playapi

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Payment = &cPayment{}

type cPayment struct{}

func (c *cPayment) Pay(ctx context.Context, req *v1.PaymentPayReq) (res *v1.PaymentPayRes, err error) {
	return service.PlayapiPayment().Pay(ctx, req)
}

var PaymentNotify = &cPaymentNotify{}

type cPaymentNotify struct{}

func (c *cPaymentNotify) WxCallback(ctx context.Context, req *v1.PaymentWxCallbackReq) (res *v1.PaymentWxCallbackRes, err error) {
	err = service.PlayapiPayment().WxCallback(ctx, req)
	return
}

func (c *cPaymentNotify) AlipayCallback(ctx context.Context, req *v1.PaymentAlipayCallbackReq) (res *v1.PaymentAlipayCallbackRes, err error) {
	err = service.PlayapiPayment().AlipayCallback(ctx, req)
	return
}
