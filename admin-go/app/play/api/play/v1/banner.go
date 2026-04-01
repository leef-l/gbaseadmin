package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Banner API

// BannerCreateReq 创建首页Banner轮播请求
type BannerCreateReq struct {
	g.Meta `path:"/banner/create" method:"post" tags:"首页Banner轮播" summary:"创建首页Banner轮播"`
	Title string `json:"title"  dc:"Banner标题"`
	Image string `json:"image"  dc:"图片URL"`
	LinkType int `json:"linkType"  dc:"跳转类型"`
	LinkValue string `json:"linkValue"  dc:"跳转值(页面路径/URL/业务ID/App Scheme)"`
	Sort int `json:"sort"  dc:"排序(越大越前)"`
	Status int `json:"status"  dc:"状态"`
	StartTime *gtime.Time `json:"startTime"  dc:"生效开始时间"`
	EndTime *gtime.Time `json:"endTime"  dc:"生效结束时间"`
	Remark string `json:"remark"  dc:"备注"`
}

// BannerCreateRes 创建首页Banner轮播响应
type BannerCreateRes struct {
	g.Meta `mime:"application/json"`
}

// BannerUpdateReq 更新首页Banner轮播请求
type BannerUpdateReq struct {
	g.Meta `path:"/banner/update" method:"put" tags:"首页Banner轮播" summary:"更新首页Banner轮播"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"首页Banner轮播ID"`
	Title string `json:"title" dc:"Banner标题"`
	Image string `json:"image" dc:"图片URL"`
	LinkType int `json:"linkType" dc:"跳转类型"`
	LinkValue string `json:"linkValue" dc:"跳转值(页面路径/URL/业务ID/App Scheme)"`
	Sort int `json:"sort" dc:"排序(越大越前)"`
	Status int `json:"status" dc:"状态"`
	StartTime *gtime.Time `json:"startTime" dc:"生效开始时间"`
	EndTime *gtime.Time `json:"endTime" dc:"生效结束时间"`
	Remark string `json:"remark" dc:"备注"`
}

// BannerUpdateRes 更新首页Banner轮播响应
type BannerUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// BannerDeleteReq 删除首页Banner轮播请求
type BannerDeleteReq struct {
	g.Meta `path:"/banner/delete" method:"delete" tags:"首页Banner轮播" summary:"删除首页Banner轮播"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"首页Banner轮播ID"`
}

// BannerDeleteRes 删除首页Banner轮播响应
type BannerDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// BannerBatchDeleteReq 批量删除首页Banner轮播请求
type BannerBatchDeleteReq struct {
	g.Meta `path:"/banner/batch-delete" method:"delete" tags:"首页Banner轮播" summary:"批量删除首页Banner轮播"`
	IDs    []snowflake.JsonInt64 `json:"ids" v:"required#ID列表不能为空" dc:"首页Banner轮播ID列表"`
}

// BannerBatchDeleteRes 批量删除首页Banner轮播响应
type BannerBatchDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// BannerDetailReq 获取首页Banner轮播详情请求
type BannerDetailReq struct {
	g.Meta `path:"/banner/detail" method:"get" tags:"首页Banner轮播" summary:"获取首页Banner轮播详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"首页Banner轮播ID"`
}

// BannerDetailRes 获取首页Banner轮播详情响应
type BannerDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.BannerDetailOutput
}

// BannerListReq 获取首页Banner轮播列表请求
type BannerListReq struct {
	g.Meta    `path:"/banner/list" method:"get" tags:"首页Banner轮播" summary:"获取首页Banner轮播列表"`
	PageNum   int    `json:"pageNum" d:"1" dc:"页码"`
	PageSize  int    `json:"pageSize" d:"10" dc:"每页数量"`
	OrderBy   string `json:"orderBy" dc:"排序字段"`
	OrderDir  string `json:"orderDir" d:"asc" dc:"排序方向:asc/desc"`
	StartTime string `json:"startTime" dc:"开始时间"`
	EndTime   string `json:"endTime" dc:"结束时间"`
	Title string `json:"title" dc:"Banner标题"`
	Remark string `json:"remark" dc:"备注"`
}

// BannerListRes 获取首页Banner轮播列表响应
type BannerListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.BannerListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}
// BannerExportReq 导出首页Banner轮播请求
type BannerExportReq struct {
	g.Meta    `path:"/banner/export" method:"get" tags:"首页Banner轮播" summary:"导出首页Banner轮播"`
	StartTime string `json:"startTime" dc:"开始时间"`
	EndTime   string `json:"endTime" dc:"结束时间"`
	Title string `json:"title" dc:"Banner标题"`
	Remark string `json:"remark" dc:"备注"`
}

// BannerExportRes 导出首页Banner轮播响应
type BannerExportRes struct {
	g.Meta `mime:"text/csv"`
}


