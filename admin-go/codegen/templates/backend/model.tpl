package model

import (
	"gbaseadmin/utility/snowflake"
)

// {{.ModelName}} DTO 模型

// {{.ModelName}}CreateInput 创建{{.Comment}}输入
type {{.ModelName}}CreateInput struct {
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden)}}
	{{.NameCamel}} {{if .IsForeignKey}}snowflake.JsonInt64{{else}}{{.GoType}}{{end}} `json:"{{.NameLower}}"`
{{- end}}
{{- end}}
}

// {{.ModelName}}UpdateInput 更新{{.Comment}}输入
type {{.ModelName}}UpdateInput struct {
	Id snowflake.JsonInt64 `json:"id"`
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden)}}
	{{.NameCamel}} {{if .IsForeignKey}}snowflake.JsonInt64{{else}}{{.GoType}}{{end}} `json:"{{.NameLower}}"`
{{- end}}
{{- end}}
}

// {{.ModelName}}DetailOutput {{.Comment}}详情输出
type {{.ModelName}}DetailOutput struct {
{{- range .Fields}}
{{- if not .IsHidden}}
	{{.NameCamel}} {{if or .IsID .IsForeignKey}}snowflake.JsonInt64{{else}}{{.GoType}}{{end}} `json:"{{.NameLower}}"`
{{- end}}
{{- end}}
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// {{.ModelName}}ListOutput {{.Comment}}列表输出
type {{.ModelName}}ListOutput struct {
{{- range .Fields}}
{{- if not .IsHidden}}
	{{.NameCamel}} {{if or .IsID .IsForeignKey}}snowflake.JsonInt64{{else}}{{.GoType}}{{end}} `json:"{{.NameLower}}"`
{{- end}}
{{- end}}
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// {{.ModelName}}ListInput {{.Comment}}列表查询输入
type {{.ModelName}}ListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}
{{if .HasParentID}}
// {{.ModelName}}TreeOutput {{.Comment}}树形输出
type {{.ModelName}}TreeOutput struct {
{{- range .Fields}}
{{- if not .IsHidden}}
	{{.NameCamel}} {{if or .IsID .IsForeignKey}}snowflake.JsonInt64{{else}}{{.GoType}}{{end}} `json:"{{.NameLower}}"`
{{- end}}
{{- end}}
	Children []*{{.ModelName}}TreeOutput `json:"children"`
}
{{end}}
