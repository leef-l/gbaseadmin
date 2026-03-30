package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Shop DTO 模型

// ShopCreateInput 创建店铺表输入
type ShopCreateInput struct {
	Title string `json:"title"`
	LogoImage string `json:"logoImage"`
	CoverImage string `json:"coverImage"`
	ContactName string `json:"contactName"`
	ContactPhone string `json:"contactPhone"`
	Intro string `json:"intro"`
	CommissionRate int `json:"commissionRate"`
	CoachNum int `json:"coachNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// ShopUpdateInput 更新店铺表输入
type ShopUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	LogoImage string `json:"logoImage"`
	CoverImage string `json:"coverImage"`
	ContactName string `json:"contactName"`
	ContactPhone string `json:"contactPhone"`
	Intro string `json:"intro"`
	CommissionRate int `json:"commissionRate"`
	CoachNum int `json:"coachNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// ShopDetailOutput 店铺表详情输出
type ShopDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	LogoImage string `json:"logoImage"`
	CoverImage string `json:"coverImage"`
	ContactName string `json:"contactName"`
	ContactPhone string `json:"contactPhone"`
	Intro string `json:"intro"`
	CommissionRate int `json:"commissionRate"`
	CoachNum int `json:"coachNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ShopListOutput 店铺表列表输出
type ShopListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Title string `json:"title"`
	LogoImage string `json:"logoImage"`
	CoverImage string `json:"coverImage"`
	ContactName string `json:"contactName"`
	ContactPhone string `json:"contactPhone"`
	Intro string `json:"intro"`
	CommissionRate int `json:"commissionRate"`
	CoachNum int `json:"coachNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ShopListInput 店铺表列表查询输入
type ShopListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Status int `json:"status"`
}

