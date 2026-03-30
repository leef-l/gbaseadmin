package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Coach DTO 模型

// CoachCreateInput 创建陪玩师表输入
type CoachCreateInput struct {
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachLevelID snowflake.JsonInt64 `json:"coachLevelID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	RealName string `json:"realName"`
	Intro string `json:"intro"`
	CoverImage string `json:"coverImage"`
	TotalOrders int `json:"totalOrders"`
	TotalScore int `json:"totalScore"`
	ScoreNum int `json:"scoreNum"`
	IncomeTotal int64 `json:"incomeTotal"`
	IncomeBalance int64 `json:"incomeBalance"`
	IsOnline int `json:"isOnline"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// CoachUpdateInput 更新陪玩师表输入
type CoachUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachLevelID snowflake.JsonInt64 `json:"coachLevelID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	RealName string `json:"realName"`
	Intro string `json:"intro"`
	CoverImage string `json:"coverImage"`
	TotalOrders int `json:"totalOrders"`
	TotalScore int `json:"totalScore"`
	ScoreNum int `json:"scoreNum"`
	IncomeTotal int64 `json:"incomeTotal"`
	IncomeBalance int64 `json:"incomeBalance"`
	IsOnline int `json:"isOnline"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// CoachDetailOutput 陪玩师表详情输出
type CoachDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachLevelID snowflake.JsonInt64 `json:"coachLevelID"`
	CoachLevelTitle string `json:"coachLevelTitle"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	ShopTitle string `json:"shopTitle"`
	RealName string `json:"realName"`
	Intro string `json:"intro"`
	CoverImage string `json:"coverImage"`
	TotalOrders int `json:"totalOrders"`
	TotalScore int `json:"totalScore"`
	ScoreNum int `json:"scoreNum"`
	IncomeTotal int64 `json:"incomeTotal"`
	IncomeBalance int64 `json:"incomeBalance"`
	IsOnline int `json:"isOnline"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CoachListOutput 陪玩师表列表输出
type CoachListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachLevelID snowflake.JsonInt64 `json:"coachLevelID"`
	CoachLevelTitle string `json:"coachLevelTitle"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	ShopTitle string `json:"shopTitle"`
	RealName string `json:"realName"`
	Intro string `json:"intro"`
	CoverImage string `json:"coverImage"`
	TotalOrders int `json:"totalOrders"`
	TotalScore int `json:"totalScore"`
	ScoreNum int `json:"scoreNum"`
	IncomeTotal int64 `json:"incomeTotal"`
	IncomeBalance int64 `json:"incomeBalance"`
	IsOnline int `json:"isOnline"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CoachListInput 陪玩师表列表查询输入
type CoachListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	IsOnline int `json:"isOnline"`
	Status int `json:"status"`
}

