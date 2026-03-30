package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Category DTO 模型

// CategoryCreateInput 创建商品分类表输入
type CategoryCreateInput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	Icon string `json:"icon"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// CategoryUpdateInput 更新商品分类表输入
type CategoryUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	Icon string `json:"icon"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// CategoryDetailOutput 商品分类表详情输出
type CategoryDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	CategoryTitle string `json:"categoryTitle"`
	Title string `json:"title"`
	Icon string `json:"icon"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CategoryListOutput 商品分类表列表输出
type CategoryListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	CategoryTitle string `json:"categoryTitle"`
	Title string `json:"title"`
	Icon string `json:"icon"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CategoryListInput 商品分类表列表查询输入
type CategoryListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Status int `json:"status"`
}

// CategoryTreeOutput 商品分类表树形输出
type CategoryTreeOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	CategoryTitle string `json:"categoryTitle"`
	Title string `json:"title"`
	Icon string `json:"icon"`
	CoverImage string `json:"coverImage"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	Children []*CategoryTreeOutput `json:"children"`
}

