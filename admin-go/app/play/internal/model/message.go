package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Message DTO 模型

// MessageCreateInput 创建会员消息输入
type MessageCreateInput struct {
	MemberID snowflake.JsonInt64 `json:"memberID"`
	Title string `json:"title"`
	Content string `json:"content"`
	MsgType int `json:"msgType"`
	BizID string `json:"bizID"`
	IsRead int `json:"isRead"`
	Status int `json:"status"`
}

// MessageUpdateInput 更新会员消息输入
type MessageUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	Title string `json:"title"`
	Content string `json:"content"`
	MsgType int `json:"msgType"`
	BizID string `json:"bizID"`
	IsRead int `json:"isRead"`
	Status int `json:"status"`
}

// MessageDetailOutput 会员消息详情输出
type MessageDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	MemberNickname string `json:"memberNickname"`
	Title string `json:"title"`
	Content string `json:"content"`
	MsgType int `json:"msgType"`
	BizID string `json:"bizID"`
	IsRead int `json:"isRead"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// MessageListOutput 会员消息列表输出
type MessageListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	MemberNickname string `json:"memberNickname"`
	Title string `json:"title"`
	Content string `json:"content"`
	MsgType int `json:"msgType"`
	BizID string `json:"bizID"`
	IsRead int `json:"isRead"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// MessageListInput 会员消息列表查询输入
type MessageListInput struct {
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
	OrderBy   string `json:"orderBy"`
	OrderDir  string `json:"orderDir"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Title string `json:"title"`
}

