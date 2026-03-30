package oauth

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Oauth = cOauth{}

type cOauth struct{}

// Create 创建第三方登录绑定表
func (c *cOauth) Create(ctx context.Context, req *v1.OauthCreateReq) (res *v1.OauthCreateRes, err error) {
	err = service.Oauth().Create(ctx, &model.OauthCreateInput{
		MemberID: req.MemberID,
		Provider: req.Provider,
		OpenID: req.OpenID,
		UnionID: req.UnionID,
		Nickname: req.Nickname,
		Avatar: req.Avatar,
		AccessToken: req.AccessToken,
		RefreshToken: req.RefreshToken,
		ExpireAt: req.ExpireAt,
	})
	return
}

// Update 更新第三方登录绑定表
func (c *cOauth) Update(ctx context.Context, req *v1.OauthUpdateReq) (res *v1.OauthUpdateRes, err error) {
	err = service.Oauth().Update(ctx, &model.OauthUpdateInput{
		ID: req.ID,
		MemberID: req.MemberID,
		Provider: req.Provider,
		OpenID: req.OpenID,
		UnionID: req.UnionID,
		Nickname: req.Nickname,
		Avatar: req.Avatar,
		AccessToken: req.AccessToken,
		RefreshToken: req.RefreshToken,
		ExpireAt: req.ExpireAt,
	})
	return
}

// Delete 删除第三方登录绑定表
func (c *cOauth) Delete(ctx context.Context, req *v1.OauthDeleteReq) (res *v1.OauthDeleteRes, err error) {
	err = service.Oauth().Delete(ctx, req.ID)
	return
}

// Detail 获取第三方登录绑定表详情
func (c *cOauth) Detail(ctx context.Context, req *v1.OauthDetailReq) (res *v1.OauthDetailRes, err error) {
	res = &v1.OauthDetailRes{}
	res.OauthDetailOutput, err = service.Oauth().Detail(ctx, req.ID)
	return
}

// List 获取第三方登录绑定表列表
func (c *cOauth) List(ctx context.Context, req *v1.OauthListReq) (res *v1.OauthListRes, err error) {
	res = &v1.OauthListRes{}
	res.List, res.Total, err = service.Oauth().List(ctx, &model.OauthListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Provider: req.Provider,
	})
	return
}

