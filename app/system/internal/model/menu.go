package model

import (
	"gbaseadmin/utility/snowflake"
)

// Menu DTO 模型

// MenuCreateInput 创建菜单表输入
type MenuCreateInput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	Type int `json:"type"`
	Path string `json:"path"`
	Component string `json:"component"`
	Permission string `json:"permission"`
	Icon string `json:"icon"`
	Sort int `json:"sort"`
	IsShow int `json:"isShow"`
	IsCache int `json:"isCache"`
	LinkURL string `json:"linkURL"`
	Status int `json:"status"`
}

// MenuUpdateInput 更新菜单表输入
type MenuUpdateInput struct {
	Id snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	Type int `json:"type"`
	Path string `json:"path"`
	Component string `json:"component"`
	Permission string `json:"permission"`
	Icon string `json:"icon"`
	Sort int `json:"sort"`
	IsShow int `json:"isShow"`
	IsCache int `json:"isCache"`
	LinkURL string `json:"linkURL"`
	Status int `json:"status"`
}

// MenuDetailOutput 菜单表详情输出
type MenuDetailOutput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	Type int `json:"type"`
	Path string `json:"path"`
	Component string `json:"component"`
	Permission string `json:"permission"`
	Icon string `json:"icon"`
	Sort int `json:"sort"`
	IsShow int `json:"isShow"`
	IsCache int `json:"isCache"`
	LinkURL string `json:"linkURL"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// MenuListOutput 菜单表列表输出
type MenuListOutput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	Type int `json:"type"`
	Path string `json:"path"`
	Component string `json:"component"`
	Permission string `json:"permission"`
	Icon string `json:"icon"`
	Sort int `json:"sort"`
	IsShow int `json:"isShow"`
	IsCache int `json:"isCache"`
	LinkURL string `json:"linkURL"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// MenuListInput 菜单表列表查询输入
type MenuListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

// MenuTreeOutput 菜单表树形输出
type MenuTreeOutput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	Type int `json:"type"`
	Path string `json:"path"`
	Component string `json:"component"`
	Permission string `json:"permission"`
	Icon string `json:"icon"`
	Sort int `json:"sort"`
	IsShow int `json:"isShow"`
	IsCache int `json:"isCache"`
	LinkURL string `json:"linkURL"`
	Status int `json:"status"`
	Children []*MenuTreeOutput `json:"children"`
}

