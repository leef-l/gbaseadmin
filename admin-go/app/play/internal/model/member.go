package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Member DTO 模型

// MemberCreateInput 创建会员表输入
type MemberCreateInput struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Gender int `json:"gender"`
	MemberLevelID snowflake.JsonInt64 `json:"memberLevelID"`
	Exp int `json:"exp"`
	Balance int64 `json:"balance"`
	IsCoach int `json:"isCoach"`
	Status int `json:"status"`
	LastLoginAt *gtime.Time `json:"lastLoginAt"`
}

// MemberUpdateInput 更新会员表输入
type MemberUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Gender int `json:"gender"`
	MemberLevelID snowflake.JsonInt64 `json:"memberLevelID"`
	Exp int `json:"exp"`
	Balance int64 `json:"balance"`
	IsCoach int `json:"isCoach"`
	Status int `json:"status"`
	LastLoginAt *gtime.Time `json:"lastLoginAt"`
}

// MemberDetailOutput 会员表详情输出
type MemberDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Phone string `json:"phone"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Gender int `json:"gender"`
	MemberLevelID snowflake.JsonInt64 `json:"memberLevelID"`
	MemberLevelTitle string `json:"memberLevelTitle"`
	Exp int `json:"exp"`
	Balance int64 `json:"balance"`
	IsCoach int `json:"isCoach"`
	Status int `json:"status"`
	LastLoginAt *gtime.Time `json:"lastLoginAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// MemberListOutput 会员表列表输出
type MemberListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	Phone string `json:"phone"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Gender int `json:"gender"`
	MemberLevelID snowflake.JsonInt64 `json:"memberLevelID"`
	MemberLevelTitle string `json:"memberLevelTitle"`
	Exp int `json:"exp"`
	Balance int64 `json:"balance"`
	IsCoach int `json:"isCoach"`
	Status int `json:"status"`
	LastLoginAt *gtime.Time `json:"lastLoginAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// MemberListInput 会员表列表查询输入
type MemberListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Gender int `json:"gender"`
	IsCoach int `json:"isCoach"`
	Status int `json:"status"`
}

