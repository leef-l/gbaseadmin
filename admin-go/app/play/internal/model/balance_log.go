package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// BalanceLog DTO 模型

// BalanceLogCreateInput 创建余额流水表输入
type BalanceLogCreateInput struct {
	MemberID snowflake.JsonInt64 `json:"memberID"`
	BizType int `json:"bizType"`
	BizID snowflake.JsonInt64 `json:"bizID"`
	ChangeAmount int64 `json:"changeAmount"`
	BeforeBalance int64 `json:"beforeBalance"`
	AfterBalance int64 `json:"afterBalance"`
	Remark string `json:"remark"`
}

// BalanceLogUpdateInput 更新余额流水表输入
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

// BalanceLogDetailOutput 余额流水表详情输出
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

// BalanceLogListOutput 余额流水表列表输出
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

// BalanceLogListInput 余额流水表列表查询输入
type BalanceLogListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	BizType int `json:"bizType"`
}

