package playapi

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var Message = &cMessage{}

type cMessage struct{}

// List 我的消息列表
func (c *cMessage) List(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error) {
	res = &v1.MessageListRes{}
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	res.List, res.Total, err = service.PlayapiMessage().List(ctx, memberID, req.Type, req.Page, req.PageSize)
	return
}

// Read 标记单条消息已读
func (c *cMessage) Read(ctx context.Context, req *v1.MessageReadReq) (res *v1.MessageReadRes, err error) {
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	err = service.PlayapiMessage().Read(ctx, memberID, req.Id)
	return
}

// ReadAll 标记全部消息已读
func (c *cMessage) ReadAll(ctx context.Context, req *v1.MessageReadAllReq) (res *v1.MessageReadAllRes, err error) {
	memberID := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	err = service.PlayapiMessage().ReadAll(ctx, memberID)
	return
}
