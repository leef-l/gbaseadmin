package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Banner DTO 模型

// BannerCreateInput 创建首页Banner轮播输入
type BannerCreateInput struct {
	Title string `json:"title"`
	Image string `json:"image"`
	LinkType int `json:"linkType"`
	LinkValue string `json:"linkValue"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	StartTime *gtime.Time `json:"startTime"`
	EndTime *gtime.Time `json:"endTime"`
	Remark string `json:"remark"`
}

// BannerUpdateInput 更新首页Banner轮播输入
type BannerUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	LinkType int `json:"linkType"`
	LinkValue string `json:"linkValue"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	StartTime *gtime.Time `json:"startTime"`
	EndTime *gtime.Time `json:"endTime"`
	Remark string `json:"remark"`
}

// BannerDetailOutput 首页Banner轮播详情输出
type BannerDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	LinkType int `json:"linkType"`
	LinkValue string `json:"linkValue"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	StartTime *gtime.Time `json:"startTime"`
	EndTime *gtime.Time `json:"endTime"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// BannerListOutput 首页Banner轮播列表输出
type BannerListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	LinkType int `json:"linkType"`
	LinkValue string `json:"linkValue"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	StartTime *gtime.Time `json:"startTime"`
	EndTime *gtime.Time `json:"endTime"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// BannerListInput 首页Banner轮播列表查询输入
type BannerListInput struct {
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
	OrderBy   string `json:"orderBy"`
	OrderDir  string `json:"orderDir"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Title string `json:"title"`
	Remark string `json:"remark"`
}

