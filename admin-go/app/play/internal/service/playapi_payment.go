package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiPayment interface {
	Pay(ctx context.Context, req *v1.PaymentPayReq) (res *v1.PaymentPayRes, err error)
	WxCallback(ctx context.Context, req *v1.PaymentWxCallbackReq) error
	AlipayCallback(ctx context.Context, req *v1.PaymentAlipayCallbackReq) error
}

var localPlayapiPayment IPlayapiPayment

func PlayapiPayment() IPlayapiPayment {
	return localPlayapiPayment
}

func RegisterPlayapiPayment(i IPlayapiPayment) {
	localPlayapiPayment = i
}
