package shop

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Shop = cShop{}

type cShop struct{}

// Create 创建店铺表
func (c *cShop) Create(ctx context.Context, req *v1.ShopCreateReq) (res *v1.ShopCreateRes, err error) {
	err = service.Shop().Create(ctx, &model.ShopCreateInput{
		Title: req.Title,
		LogoImage: req.LogoImage,
		CoverImage: req.CoverImage,
		ContactName: req.ContactName,
		ContactPhone: req.ContactPhone,
		Intro: req.Intro,
		CommissionRate: req.CommissionRate,
		CoachNum: req.CoachNum,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Update 更新店铺表
func (c *cShop) Update(ctx context.Context, req *v1.ShopUpdateReq) (res *v1.ShopUpdateRes, err error) {
	err = service.Shop().Update(ctx, &model.ShopUpdateInput{
		ID: req.ID,
		Title: req.Title,
		LogoImage: req.LogoImage,
		CoverImage: req.CoverImage,
		ContactName: req.ContactName,
		ContactPhone: req.ContactPhone,
		Intro: req.Intro,
		CommissionRate: req.CommissionRate,
		CoachNum: req.CoachNum,
		Sort: req.Sort,
		Status: req.Status,
	})
	return
}

// Delete 删除店铺表
func (c *cShop) Delete(ctx context.Context, req *v1.ShopDeleteReq) (res *v1.ShopDeleteRes, err error) {
	err = service.Shop().Delete(ctx, req.ID)
	return
}

// Detail 获取店铺表详情
func (c *cShop) Detail(ctx context.Context, req *v1.ShopDetailReq) (res *v1.ShopDetailRes, err error) {
	res = &v1.ShopDetailRes{}
	res.ShopDetailOutput, err = service.Shop().Detail(ctx, req.ID)
	return
}

// List 获取店铺表列表
func (c *cShop) List(ctx context.Context, req *v1.ShopListReq) (res *v1.ShopListRes, err error) {
	res = &v1.ShopListRes{}
	res.List, res.Total, err = service.Shop().List(ctx, &model.ShopListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Status: req.Status,
	})
	return
}

