package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiMember interface {
	Info(ctx context.Context, req *v1.MemberInfoReq) (*v1.MemberInfoRes, error)
	Update(ctx context.Context, req *v1.MemberUpdateReq) (*v1.MemberUpdateRes, error)
	SwitchRole(ctx context.Context, req *v1.MemberSwitchRoleReq) (*v1.MemberSwitchRoleRes, error)
	BalanceLog(ctx context.Context, req *v1.MemberBalanceLogReq) (*v1.MemberBalanceLogRes, error)
}

var localPlayapiMember IPlayapiMember

func PlayapiMember() IPlayapiMember { return localPlayapiMember }

func RegisterPlayapiMember(s IPlayapiMember) { localPlayapiMember = s }
