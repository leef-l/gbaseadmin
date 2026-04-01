package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// ProfitLog DTO 模型

// ProfitLogCreateInput 创建利润分成流水表输入
type ProfitLogCreateInput struct {
	OrderID snowflake.JsonInt64 `json:"orderID"`
	OrderNo string `json:"orderNo"`
	PayAmount int64 `json:"payAmount"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	PlatformRate int `json:"platformRate"`
	PlatformAmount int64 `json:"platformAmount"`
	ShopRate int `json:"shopRate"`
	ShopAmount int64 `json:"shopAmount"`
	CoachAmount int64 `json:"coachAmount"`
	SettleStatus int `json:"settleStatus"`
	SettleAt *gtime.Time `json:"settleAt"`
}

// ProfitLogUpdateInput 更新利润分成流水表输入
type ProfitLogUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	OrderNo string `json:"orderNo"`
	PayAmount int64 `json:"payAmount"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	PlatformRate int `json:"platformRate"`
	PlatformAmount int64 `json:"platformAmount"`
	ShopRate int `json:"shopRate"`
	ShopAmount int64 `json:"shopAmount"`
	CoachAmount int64 `json:"coachAmount"`
	SettleStatus int `json:"settleStatus"`
	SettleAt *gtime.Time `json:"settleAt"`
}

// ProfitLogDetailOutput 利润分成流水表详情输出
type ProfitLogDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	OrderNo string `json:"orderNo"`
	PayAmount int64 `json:"payAmount"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	CoachRealName string `json:"coachRealName"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	ShopTitle string `json:"shopTitle"`
	PlatformRate int `json:"platformRate"`
	PlatformAmount int64 `json:"platformAmount"`
	ShopRate int `json:"shopRate"`
	ShopAmount int64 `json:"shopAmount"`
	CoachAmount int64 `json:"coachAmount"`
	SettleStatus int `json:"settleStatus"`
	SettleAt *gtime.Time `json:"settleAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ProfitLogListOutput 利润分成流水表列表输出
type ProfitLogListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	OrderNo string `json:"orderNo"`
	PayAmount int64 `json:"payAmount"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	CoachRealName string `json:"coachRealName"`
	ShopID snowflake.JsonInt64 `json:"shopID"`
	ShopTitle string `json:"shopTitle"`
	PlatformRate int `json:"platformRate"`
	PlatformAmount int64 `json:"platformAmount"`
	ShopRate int `json:"shopRate"`
	ShopAmount int64 `json:"shopAmount"`
	CoachAmount int64 `json:"coachAmount"`
	SettleStatus int `json:"settleStatus"`
	SettleAt *gtime.Time `json:"settleAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ProfitLogListInput 利润分成流水表列表查询输入
type ProfitLogListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	SettleStatus int `json:"settleStatus"`
}

