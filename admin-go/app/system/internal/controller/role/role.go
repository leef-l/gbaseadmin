package role

import (
	"context"

	v1 "gbaseadmin/app/system/api/system/v1"
	"gbaseadmin/app/system/internal/model"
	"gbaseadmin/app/system/internal/service"
)

var Role = cRole{}

type cRole struct{}

// Create 创建角色表
func (c *cRole) Create(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error) {
	err = service.Role().Create(ctx, &model.RoleCreateInput{
		ParentID: req.ParentID,
		Title: req.Title,
		DataScope: req.DataScope,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新角色表
func (c *cRole) Update(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	err = service.Role().Update(ctx, &model.RoleUpdateInput{
		ID: req.ID,
		ParentID: req.ParentID,
		Title: req.Title,
		DataScope: req.DataScope,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除角色表
func (c *cRole) Delete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.ID)
	return
}

// Detail 获取角色表详情
func (c *cRole) Detail(ctx context.Context, req *v1.RoleDetailReq) (res *v1.RoleDetailRes, err error) {
	res = &v1.RoleDetailRes{}
	res.RoleDetailOutput, err = service.Role().Detail(ctx, req.ID)
	return
}

// List 获取角色表列表
func (c *cRole) List(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	res = &v1.RoleListRes{}
	res.List, res.Total, err = service.Role().List(ctx, &model.RoleListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	return
}

// Tree 获取角色表树形结构
func (c *cRole) Tree(ctx context.Context, req *v1.RoleTreeReq) (res *v1.RoleTreeRes, err error) {
	res = &v1.RoleTreeRes{}
	res.List, err = service.Role().Tree(ctx)
	return
}

