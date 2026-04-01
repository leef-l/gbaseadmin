package banner

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Banner = cBanner{}

type cBanner struct{}

// Create 创建首页Banner轮播
func (c *cBanner) Create(ctx context.Context, req *v1.BannerCreateReq) (res *v1.BannerCreateRes, err error) {
	err = service.Banner().Create(ctx, &model.BannerCreateInput{
		Title: req.Title,
		Image: req.Image,
		LinkType: req.LinkType,
		LinkValue: req.LinkValue,
		Sort: req.Sort,
		Status: req.Status,
		StartTime: req.StartTime,
		EndTime: req.EndTime,
		Remark: req.Remark,
	})
	return
}

// Update 更新首页Banner轮播
func (c *cBanner) Update(ctx context.Context, req *v1.BannerUpdateReq) (res *v1.BannerUpdateRes, err error) {
	err = service.Banner().Update(ctx, &model.BannerUpdateInput{
		ID: req.ID,
		Title: req.Title,
		Image: req.Image,
		LinkType: req.LinkType,
		LinkValue: req.LinkValue,
		Sort: req.Sort,
		Status: req.Status,
		StartTime: req.StartTime,
		EndTime: req.EndTime,
		Remark: req.Remark,
	})
	return
}

// Delete 删除首页Banner轮播
func (c *cBanner) Delete(ctx context.Context, req *v1.BannerDeleteReq) (res *v1.BannerDeleteRes, err error) {
	err = service.Banner().Delete(ctx, req.ID)
	return
}

// BatchDelete 批量删除首页Banner轮播
func (c *cBanner) BatchDelete(ctx context.Context, req *v1.BannerBatchDeleteReq) (res *v1.BannerBatchDeleteRes, err error) {
	err = service.Banner().BatchDelete(ctx, req.IDs)
	return
}

// Detail 获取首页Banner轮播详情
func (c *cBanner) Detail(ctx context.Context, req *v1.BannerDetailReq) (res *v1.BannerDetailRes, err error) {
	res = &v1.BannerDetailRes{}
	res.BannerDetailOutput, err = service.Banner().Detail(ctx, req.ID)
	return
}

// List 获取首页Banner轮播列表
func (c *cBanner) List(ctx context.Context, req *v1.BannerListReq) (res *v1.BannerListRes, err error) {
	res = &v1.BannerListRes{}
	res.List, res.Total, err = service.Banner().List(ctx, &model.BannerListInput{
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
		OrderBy:   req.OrderBy,
		OrderDir:  req.OrderDir,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Title: req.Title,
		Remark: req.Remark,
	})
	return
}
// Export 导出首页Banner轮播
func (c *cBanner) Export(ctx context.Context, req *v1.BannerExportReq) (res *v1.BannerExportRes, err error) {
	list, err := service.Banner().Export(ctx, &model.BannerListInput{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Title: req.Title,
		Remark: req.Remark,
	})
	if err != nil {
		return
	}
	// CSV 导出
	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", "text/csv; charset=utf-8")
	r.Response.Header().Set("Content-Disposition", `attachment; filename="banner.csv"`)
	r.Response.Write("\xEF\xBB\xBF") // UTF-8 BOM
	// 表头
	r.Response.Writeln("Banner标题,图片URL,跳转类型,跳转值,排序,状态,生效开始时间,生效结束时间,备注,创建时间")
	// 数据行
	for _, item := range list {
		r.Response.Writefln("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v",
			item.Title,
			 item.Image,
			 item.LinkType,
			 item.LinkValue,
			 item.Sort,
			 item.Status,
			 item.StartTime,
			 item.EndTime,
			 item.Remark,
			item.CreatedAt,
		)
	}
	return
}

