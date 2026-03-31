package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// File DTO 模型

// FileCreateInput 创建æ–‡ä»¶è®°å½•输入
type FileCreateInput struct {
	DirID snowflake.JsonInt64 `json:"dirID"`
	Name string `json:"name"`
	URL string `json:"url"`
	Ext string `json:"ext"`
	Size int64 `json:"size"`
	Mime string `json:"mime"`
	Storage int `json:"storage"`
	IsImage int `json:"isImage"`
}

// FileUpdateInput 更新æ–‡ä»¶è®°å½•输入
type FileUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	DirID snowflake.JsonInt64 `json:"dirID"`
	Name string `json:"name"`
	URL string `json:"url"`
	Ext string `json:"ext"`
	Size int64 `json:"size"`
	Mime string `json:"mime"`
	Storage int `json:"storage"`
	IsImage int `json:"isImage"`
}

// FileDetailOutput æ–‡ä»¶è®°å½•详情输出
type FileDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	DirID snowflake.JsonInt64 `json:"dirID"`
	DirName string `json:"dirName"`
	Name string `json:"name"`
	URL string `json:"url"`
	Ext string `json:"ext"`
	Size int64 `json:"size"`
	Mime string `json:"mime"`
	Storage int `json:"storage"`
	IsImage int `json:"isImage"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// FileListOutput æ–‡ä»¶è®°å½•列表输出
type FileListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	DirID snowflake.JsonInt64 `json:"dirID"`
	DirName string `json:"dirName"`
	Name string `json:"name"`
	URL string `json:"url"`
	Ext string `json:"ext"`
	Size int64 `json:"size"`
	Mime string `json:"mime"`
	Storage int `json:"storage"`
	IsImage int `json:"isImage"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// FileListInput æ–‡ä»¶è®°å½•列表查询输入
type FileListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Storage int `json:"storage"`
	IsImage int `json:"isImage"`
}

