package model

import (
	"gbaseadmin/utility/snowflake"
)

// Users DTO 模型

// UsersCreateInput 创建用户表输入
type UsersCreateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Status int `json:"status"`
}

// UsersUpdateInput 更新用户表输入
type UsersUpdateInput struct {
	Id snowflake.JsonInt64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Status int `json:"status"`
}

// UsersDetailOutput 用户表详情输出
type UsersDetailOutput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// UsersListOutput 用户表列表输出
type UsersListOutput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// UsersListInput 用户表列表查询输入
type UsersListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

