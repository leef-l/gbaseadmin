package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Review DTO 模型

// ReviewCreateInput 创建评价表输入
type ReviewCreateInput struct {
	OrderID snowflake.JsonInt64 `json:"orderID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	Score int `json:"score"`
	ReviewContent string `json:"reviewContent"`
	ReviewImage string `json:"reviewImage"`
	ReplyContent string `json:"replyContent"`
	ReplyAt *gtime.Time `json:"replyAt"`
	IsAnonymous int `json:"isAnonymous"`
	Status int `json:"status"`
}

// ReviewUpdateInput 更新评价表输入
type ReviewUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	Score int `json:"score"`
	ReviewContent string `json:"reviewContent"`
	ReviewImage string `json:"reviewImage"`
	ReplyContent string `json:"replyContent"`
	ReplyAt *gtime.Time `json:"replyAt"`
	IsAnonymous int `json:"isAnonymous"`
	Status int `json:"status"`
}

// ReviewDetailOutput 评价表详情输出
type ReviewDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	Score int `json:"score"`
	ReviewContent string `json:"reviewContent"`
	ReviewImage string `json:"reviewImage"`
	ReplyContent string `json:"replyContent"`
	ReplyAt *gtime.Time `json:"replyAt"`
	IsAnonymous int `json:"isAnonymous"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ReviewListOutput 评价表列表输出
type ReviewListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	OrderID snowflake.JsonInt64 `json:"orderID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	CoachID snowflake.JsonInt64 `json:"coachID"`
	Score int `json:"score"`
	ReviewContent string `json:"reviewContent"`
	ReviewImage string `json:"reviewImage"`
	ReplyContent string `json:"replyContent"`
	ReplyAt *gtime.Time `json:"replyAt"`
	IsAnonymous int `json:"isAnonymous"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// ReviewListInput 评价表列表查询输入
type ReviewListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	IsAnonymous int `json:"isAnonymous"`
	Status int `json:"status"`
}

