package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Dir DTO 模型

// DirCreateInput 创建æ–‡ä»¶ç›®å½•输入
type DirCreateInput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Name string `json:"name"`
	Path string `json:"path"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// DirUpdateInput 更新æ–‡ä»¶ç›®å½•输入
type DirUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Name string `json:"name"`
	Path string `json:"path"`
	Sort int `json:"sort"`
	Status int `json:"status"`
}

// DirDetailOutput æ–‡ä»¶ç›®å½•详情输出
type DirDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	DirName string `json:"dirName"`
	Name string `json:"name"`
	Path string `json:"path"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// DirListOutput æ–‡ä»¶ç›®å½•列表输出
type DirListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	DirName string `json:"dirName"`
	Name string `json:"name"`
	Path string `json:"path"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// DirListInput æ–‡ä»¶ç›®å½•列表查询输入
type DirListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Status int `json:"status"`
}

// DirTreeOutput æ–‡ä»¶ç›®å½•树形输出
type DirTreeOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	DirName string `json:"dirName"`
	Name string `json:"name"`
	Path string `json:"path"`
	Sort int `json:"sort"`
	Status int `json:"status"`
	Children []*DirTreeOutput `json:"children"`
}

