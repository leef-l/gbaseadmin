package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Goods DTO 模型

// GoodsCreateInput 创建商品表输入
type GoodsCreateInput struct {
	CategoryID snowflake.JsonInt64 `json:"categoryID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	Title string `json:"title"`
	CoverImage string `json:"coverImage"`
	DescContent string `json:"descContent"`
	Price int64 `json:"price"`
	Unit string `json:"unit"`
	SalesNum int `json:"salesNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// GoodsUpdateInput 更新商品表输入
type GoodsUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	CategoryID snowflake.JsonInt64 `json:"categoryID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	Title string `json:"title"`
	CoverImage string `json:"coverImage"`
	DescContent string `json:"descContent"`
	Price int64 `json:"price"`
	Unit string `json:"unit"`
	SalesNum int `json:"salesNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// GoodsDetailOutput 商品表详情输出
type GoodsDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	CategoryID snowflake.JsonInt64 `json:"categoryID"`
	CategoryTitle string `json:"categoryTitle"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	Title string `json:"title"`
	CoverImage string `json:"coverImage"`
	DescContent string `json:"descContent"`
	Price int64 `json:"price"`
	Unit string `json:"unit"`
	SalesNum int `json:"salesNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// GoodsListOutput 商品表列表输出
type GoodsListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	CategoryID snowflake.JsonInt64 `json:"categoryID"`
	CategoryTitle string `json:"categoryTitle"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	Title string `json:"title"`
	CoverImage string `json:"coverImage"`
	DescContent string `json:"descContent"`
	Price int64 `json:"price"`
	Unit string `json:"unit"`
	SalesNum int `json:"salesNum"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// GoodsListInput 商品表列表查询输入
type GoodsListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Status int `json:"status"`
}

