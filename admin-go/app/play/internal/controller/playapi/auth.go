package playapi

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Auth = &cAuth{}

type cAuth struct{}

func (c *cAuth) Login(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	return service.PlayapiAuth().Login(ctx, req)
}

func (c *cAuth) SendCode(ctx context.Context, req *v1.AuthSendCodeReq) (res *v1.AuthSendCodeRes, err error) {
	return service.PlayapiAuth().SendCode(ctx, req)
}

func (c *cAuth) RefreshToken(ctx context.Context, req *v1.AuthRefreshTokenReq) (res *v1.AuthRefreshTokenRes, err error) {
	return service.PlayapiAuth().RefreshToken(ctx, req)
}

func (c *cAuth) WxLogin(ctx context.Context, req *v1.AuthWxLoginReq) (res *v1.AuthWxLoginRes, err error) {
	return service.PlayapiAuth().WxLogin(ctx, req)
}

func (c *cAuth) AlipayLogin(ctx context.Context, req *v1.AuthAlipayLoginReq) (res *v1.AuthAlipayLoginRes, err error) {
	return service.PlayapiAuth().AlipayLogin(ctx, req)
}
