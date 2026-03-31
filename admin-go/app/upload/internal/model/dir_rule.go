package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// DirRule DTO 模型

// DirRuleCreateInput 创建æ–‡ä»¶ç›®å½•è§„åˆ™输入
type DirRuleCreateInput struct {
	DirID snowflake.JsonInt64 `json:"dirID"`
	Category int `json:"category"`
	SavePath string `json:"savePath"`
	Status int `json:"status"`
}

// DirRuleUpdateInput 更新æ–‡ä»¶ç›®å½•è§„åˆ™输入
type DirRuleUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	DirID snowflake.JsonInt64 `json:"dirID"`
	Category int `json:"category"`
	SavePath string `json:"savePath"`
	Status int `json:"status"`
}

// DirRuleDetailOutput æ–‡ä»¶ç›®å½•è§„åˆ™详情输出
type DirRuleDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	DirID snowflake.JsonInt64 `json:"dirID"`
	DirName string `json:"dirName"`
	Category int `json:"category"`
	SavePath string `json:"savePath"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// DirRuleListOutput æ–‡ä»¶ç›®å½•è§„åˆ™列表输出
type DirRuleListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	DirID snowflake.JsonInt64 `json:"dirID"`
	DirName string `json:"dirName"`
	Category int `json:"category"`
	SavePath string `json:"savePath"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// DirRuleListInput æ–‡ä»¶ç›®å½•è§„åˆ™列表查询输入
type DirRuleListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Category int `json:"category"`
	Status int `json:"status"`
}

