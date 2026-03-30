package playapi

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Member = &cMember{}

type cMember struct{}

func (c *cMember) Info(ctx context.Context, req *v1.MemberInfoReq) (res *v1.MemberInfoRes, err error) {
	return service.PlayapiMember().Info(ctx, req)
}

func (c *cMember) Update(ctx context.Context, req *v1.MemberUpdateReq) (res *v1.MemberUpdateRes, err error) {
	return service.PlayapiMember().Update(ctx, req)
}

func (c *cMember) SwitchRole(ctx context.Context, req *v1.MemberSwitchRoleReq) (res *v1.MemberSwitchRoleRes, err error) {
	return service.PlayapiMember().SwitchRole(ctx, req)
}

func (c *cMember) BalanceLog(ctx context.Context, req *v1.MemberBalanceLogReq) (res *v1.MemberBalanceLogRes, err error) {
	return service.PlayapiMember().BalanceLog(ctx, req)
}
