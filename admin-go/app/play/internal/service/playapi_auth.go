package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiAuth interface {
	Login(ctx context.Context, req *v1.AuthLoginReq) (*v1.AuthLoginRes, error)
	SendCode(ctx context.Context, req *v1.AuthSendCodeReq) (*v1.AuthSendCodeRes, error)
	RefreshToken(ctx context.Context, req *v1.AuthRefreshTokenReq) (*v1.AuthRefreshTokenRes, error)
	WxLogin(ctx context.Context, req *v1.AuthWxLoginReq) (*v1.AuthWxLoginRes, error)
	AlipayLogin(ctx context.Context, req *v1.AuthAlipayLoginReq) (*v1.AuthAlipayLoginRes, error)
}

var localPlayapiAuth IPlayapiAuth

func PlayapiAuth() IPlayapiAuth { return localPlayapiAuth }

func RegisterPlayapiAuth(s IPlayapiAuth) { localPlayapiAuth = s }
