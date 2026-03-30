package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// SysTest DTO 模型

// SysTestCreateInput 创建æµ‹è¯•è¡¨输入
type SysTestCreateInput struct {
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	Code string `json:"code"`
	Type int `json:"type"`
	Status int `json:"status"`
	Sort int `json:"sort"`
	Remark string `json:"remark"`
}

// SysTestUpdateInput 更新æµ‹è¯•è¡¨输入
type SysTestUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	Title string `json:"title"`
	Code string `json:"code"`
	Type int `json:"type"`
	Status int `json:"status"`
	Sort int `json:"sort"`
	Remark string `json:"remark"`
}

// SysTestDetailOutput æµ‹è¯•è¡¨详情输出
type SysTestDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	SysTestTitle string `json:"sysTestTitle"`
	Title string `json:"title"`
	Code string `json:"code"`
	Type int `json:"type"`
	Status int `json:"status"`
	Sort int `json:"sort"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// SysTestListOutput æµ‹è¯•è¡¨列表输出
type SysTestListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	SysTestTitle string `json:"sysTestTitle"`
	Title string `json:"title"`
	Code string `json:"code"`
	Type int `json:"type"`
	Status int `json:"status"`
	Sort int `json:"sort"`
	Remark string `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// SysTestListInput æµ‹è¯•è¡¨列表查询输入
type SysTestListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Type int `json:"type"`
	Status int `json:"status"`
}

// SysTestTreeOutput æµ‹è¯•è¡¨树形输出
type SysTestTreeOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	ParentID snowflake.JsonInt64 `json:"parentID"`
	SysTestTitle string `json:"sysTestTitle"`
	Title string `json:"title"`
	Code string `json:"code"`
	Type int `json:"type"`
	Status int `json:"status"`
	Sort int `json:"sort"`
	Remark string `json:"remark"`
	Children []*SysTestTreeOutput `json:"children"`
}

