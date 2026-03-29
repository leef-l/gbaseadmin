package model

import (
	"gbaseadmin/utility/snowflake"
)

// Role DTO 模型

// RoleCreateInput 创建角色表输入
type RoleCreateInput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	DataScope int `json:"dataScope"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// RoleUpdateInput 更新角色表输入
type RoleUpdateInput struct {
	Id snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	DataScope int `json:"dataScope"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// RoleDetailOutput 角色表详情输出
type RoleDetailOutput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	DataScope int `json:"dataScope"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// RoleListOutput 角色表列表输出
type RoleListOutput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	DataScope int `json:"dataScope"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// RoleListInput 角色表列表查询输入
type RoleListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

// RoleTreeOutput 角色表树形输出
type RoleTreeOutput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	DataScope int `json:"dataScope"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	Children []*RoleTreeOutput `json:"children"`
}

