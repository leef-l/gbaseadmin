package goods

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Goods = cGoods{}

type cGoods struct{}

// Create 创建å•†å“è¡¨
func (c *cGoods) Create(ctx context.Context, req *v1.GoodsCreateReq) (res *v1.GoodsCreateRes, err error) {
	err = service.Goods().Create(ctx, &model.GoodsCreateInput{
		CategoryID: req.CategoryID,
		CoachID: req.CoachID,
		Title: req.Title,
		CoverImage: req.CoverImage,
		DescContent: req.DescContent,
		Price: req.Price,
		Unit: req.Unit,
		SalesNum: req.SalesNum,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新å•†å“è¡¨
func (c *cGoods) Update(ctx context.Context, req *v1.GoodsUpdateReq) (res *v1.GoodsUpdateRes, err error) {
	err = service.Goods().Update(ctx, &model.GoodsUpdateInput{
		ID: req.ID,
		CategoryID: req.CategoryID,
		CoachID: req.CoachID,
		Title: req.Title,
		CoverImage: req.CoverImage,
		DescContent: req.DescContent,
		Price: req.Price,
		Unit: req.Unit,
		SalesNum: req.SalesNum,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除å•†å“è¡¨
func (c *cGoods) Delete(ctx context.Context, req *v1.GoodsDeleteReq) (res *v1.GoodsDeleteRes, err error) {
	err = service.Goods().Delete(ctx, req.ID)
	return
}

// Detail 获取å•†å“è¡¨详情
func (c *cGoods) Detail(ctx context.Context, req *v1.GoodsDetailReq) (res *v1.GoodsDetailRes, err error) {
	res = &v1.GoodsDetailRes{}
	res.GoodsDetailOutput, err = service.Goods().Detail(ctx, req.ID)
	return
}

// List 获取å•†å“è¡¨列表
func (c *cGoods) List(ctx context.Context, req *v1.GoodsListReq) (res *v1.GoodsListRes, err error) {
	res = &v1.GoodsListRes{}
	res.List, res.Total, err = service.Goods().List(ctx, &model.GoodsListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Status: req.Status,
	})
	return
}

