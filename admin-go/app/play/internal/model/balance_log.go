package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// BalanceLog DTO 模型

// BalanceLogCreateInput 创建ä½™é¢æµæ°´è¡¨输入
type BalanceLogCreateInput struct {
	MemberID snowflake.JsonInt64 `json:"memberID"`
	BizType int `json:"bizType"`
	BizID snowflake.JsonInt64 `json:"bizID"`
	ChangeAmount int64 `json:"changeAmount"`
	BeforeBalance int64 `json:"beforeBalance"`
	AfterBalance int64 `json:"afterBalance"`
	Remark string `json:"remark"`
}

// BalanceLogUpdateInput 更新ä½™é¢æµæ°´è¡¨输入
type BalanceLogUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	BizType int `json:"bizType"`
	BizID snowflake.JsonInt64 `json:"bizID"`
	ChangeAmount int64 `json:"changeAmount"`
	BeforeBalance int64 `json:"beforeBalance"`
	AfterBalance int64 `json:"afterBalance"`
	Remark string `json:"remark"`
}

// BalanceLogDetailOutput ä½™é¢æµæ°´è¡¨详情输出
type BalanceLogDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	BizType int `json:"bizType"`
	BizID snowflake.JsonInt64 `json:"bizID"`
	ChangeAmount int64 `json:"changeAmount"`
	BeforeBalance int64 `json:"beforeBalance"`
	AfterBalance int64 `json:"afterBalance"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// BalanceLogListOutput ä½™é¢æµæ°´è¡¨列表输出
type BalanceLogListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	BizType int `json:"bizType"`
	BizID snowflake.JsonInt64 `json:"bizID"`
	ChangeAmount int64 `json:"changeAmount"`
	BeforeBalance int64 `json:"beforeBalance"`
	AfterBalance int64 `json:"afterBalance"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// BalanceLogListInput ä½™é¢æµæ°´è¡¨列表查询输入
type BalanceLogListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	BizType int `json:"bizType"`
}

