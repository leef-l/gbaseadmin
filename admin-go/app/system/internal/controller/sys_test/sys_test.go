package sys_test

import (
	"context"

	v1 "gbaseadmin/app/system/api/system/v1"
	"gbaseadmin/app/system/internal/model"
	"gbaseadmin/app/system/internal/service"
)

var SysTest = cSysTest{}

type cSysTest struct{}

// Create 创建æµ‹è¯•è¡¨
func (c *cSysTest) Create(ctx context.Context, req *v1.SysTestCreateReq) (res *v1.SysTestCreateRes, err error) {
	err = service.SysTest().Create(ctx, &model.SysTestCreateInput{
		ParentID: req.ParentID,
		Title: req.Title,
		Code: req.Code,
		Type: req.Type,
		Status: req.Status,
		Sort: req.Sort,
		Remark: req.Remark,
	})
	return
}

// Update 更新æµ‹è¯•è¡¨
func (c *cSysTest) Update(ctx context.Context, req *v1.SysTestUpdateReq) (res *v1.SysTestUpdateRes, err error) {
	err = service.SysTest().Update(ctx, &model.SysTestUpdateInput{
		ID: req.ID,
		ParentID: req.ParentID,
		Title: req.Title,
		Code: req.Code,
		Type: req.Type,
		Status: req.Status,
		Sort: req.Sort,
		Remark: req.Remark,
	})
	return
}

// Delete 删除æµ‹è¯•è¡¨
func (c *cSysTest) Delete(ctx context.Context, req *v1.SysTestDeleteReq) (res *v1.SysTestDeleteRes, err error) {
	err = service.SysTest().Delete(ctx, req.ID)
	return
}

// Detail 获取æµ‹è¯•è¡¨详情
func (c *cSysTest) Detail(ctx context.Context, req *v1.SysTestDetailReq) (res *v1.SysTestDetailRes, err error) {
	res = &v1.SysTestDetailRes{}
	res.SysTestDetailOutput, err = service.SysTest().Detail(ctx, req.ID)
	return
}

// List 获取æµ‹è¯•è¡¨列表
func (c *cSysTest) List(ctx context.Context, req *v1.SysTestListReq) (res *v1.SysTestListRes, err error) {
	res = &v1.SysTestListRes{}
	res.List, res.Total, err = service.SysTest().List(ctx, &model.SysTestListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Type: req.Type,
		Status: req.Status,
	})
	return
}

// Tree 获取æµ‹è¯•è¡¨树形结构
func (c *cSysTest) Tree(ctx context.Context, req *v1.SysTestTreeReq) (res *v1.SysTestTreeRes, err error) {
	res = &v1.SysTestTreeRes{}
	res.List, err = service.SysTest().Tree(ctx)
	return
}

