package category

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Category = cCategory{}

type cCategory struct{}

// Create 创建商品分类表
func (c *cCategory) Create(ctx context.Context, req *v1.CategoryCreateReq) (res *v1.CategoryCreateRes, err error) {
	err = service.Category().Create(ctx, &model.CategoryCreateInput{
		ParentID: req.ParentID,
		Title: req.Title,
		Icon: req.Icon,
		CoverImage: req.CoverImage,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新商品分类表
func (c *cCategory) Update(ctx context.Context, req *v1.CategoryUpdateReq) (res *v1.CategoryUpdateRes, err error) {
	err = service.Category().Update(ctx, &model.CategoryUpdateInput{
		ID: req.ID,
		ParentID: req.ParentID,
		Title: req.Title,
		Icon: req.Icon,
		CoverImage: req.CoverImage,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除商品分类表
func (c *cCategory) Delete(ctx context.Context, req *v1.CategoryDeleteReq) (res *v1.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, req.ID)
	return
}

// Detail 获取商品分类表详情
func (c *cCategory) Detail(ctx context.Context, req *v1.CategoryDetailReq) (res *v1.CategoryDetailRes, err error) {
	res = &v1.CategoryDetailRes{}
	res.CategoryDetailOutput, err = service.Category().Detail(ctx, req.ID)
	return
}

// List 获取商品分类表列表
func (c *cCategory) List(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error) {
	res = &v1.CategoryListRes{}
	res.List, res.Total, err = service.Category().List(ctx, &model.CategoryListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Status: req.Status,
	})
	return
}

// Tree 获取商品分类表树形结构
func (c *cCategory) Tree(ctx context.Context, req *v1.CategoryTreeReq) (res *v1.CategoryTreeRes, err error) {
	res = &v1.CategoryTreeRes{}
	res.List, err = service.Category().Tree(ctx)
	return
}

