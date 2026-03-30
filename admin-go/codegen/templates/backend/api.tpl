package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/{{.AppName}}/internal/model"
	"gbaseadmin/utility/snowflake"
)

// {{.ModelName}} API

// {{.ModelName}}CreateReq 创建{{.Comment}}请求
type {{.ModelName}}CreateReq struct {
	g.Meta `path:"/{{.ModuleName}}/create" method:"post" tags:"{{.Comment}}" summary:"创建{{.Comment}}"`
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden)}}
	{{.NameCamel}} {{if .IsForeignKey}}snowflake.JsonInt64{{else}}{{.GoType}}{{end}} `json:"{{.NameLower}}" {{if .IsRequired}}v:"required#{{.Label}}不能为空"{{end}} dc:"{{.Label}}"`
{{- end}}
{{- end}}
}

// {{.ModelName}}CreateRes 创建{{.Comment}}响应
type {{.ModelName}}CreateRes struct {
	g.Meta `mime:"application/json"`
}

// {{.ModelName}}UpdateReq 更新{{.Comment}}请求
type {{.ModelName}}UpdateReq struct {
	g.Meta `path:"/{{.ModuleName}}/update" method:"put" tags:"{{.Comment}}" summary:"更新{{.Comment}}"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"{{.Comment}}ID"`
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden)}}
	{{.NameCamel}} {{if .IsForeignKey}}snowflake.JsonInt64{{else}}{{.GoType}}{{end}} `json:"{{.NameLower}}" dc:"{{.Label}}"`
{{- end}}
{{- end}}
}

// {{.ModelName}}UpdateRes 更新{{.Comment}}响应
type {{.ModelName}}UpdateRes struct {
	g.Meta `mime:"application/json"`
}

// {{.ModelName}}DeleteReq 删除{{.Comment}}请求
type {{.ModelName}}DeleteReq struct {
	g.Meta `path:"/{{.ModuleName}}/delete" method:"delete" tags:"{{.Comment}}" summary:"删除{{.Comment}}"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"{{.Comment}}ID"`
}

// {{.ModelName}}DeleteRes 删除{{.Comment}}响应
type {{.ModelName}}DeleteRes struct {
	g.Meta `mime:"application/json"`
}

// {{.ModelName}}DetailReq 获取{{.Comment}}详情请求
type {{.ModelName}}DetailReq struct {
	g.Meta `path:"/{{.ModuleName}}/detail" method:"get" tags:"{{.Comment}}" summary:"获取{{.Comment}}详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"{{.Comment}}ID"`
}

// {{.ModelName}}DetailRes 获取{{.Comment}}详情响应
type {{.ModelName}}DetailRes struct {
	g.Meta `mime:"application/json"`
	*model.{{.ModelName}}DetailOutput
}

// {{.ModelName}}ListReq 获取{{.Comment}}列表请求
type {{.ModelName}}ListReq struct {
	g.Meta   `path:"/{{.ModuleName}}/list" method:"get" tags:"{{.Comment}}" summary:"获取{{.Comment}}列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
	{{.NameCamel}} int `json:"{{.NameLower}}" dc:"{{.Label}}"`
{{- end}}
{{- end}}
}

// {{.ModelName}}ListRes 获取{{.Comment}}列表响应
type {{.ModelName}}ListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.{{.ModelName}}ListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}
{{if .HasParentID}}
// {{.ModelName}}TreeReq 获取{{.Comment}}树形结构请求
type {{.ModelName}}TreeReq struct {
	g.Meta `path:"/{{.ModuleName}}/tree" method:"get" tags:"{{.Comment}}" summary:"获取{{.Comment}}树形结构"`
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (not .IsParentID) (.IsEnum)}}
	{{.NameCamel}} int `json:"{{.NameLower}}" dc:"{{.Label}}"`
{{- end}}
{{- end}}
}

// {{.ModelName}}TreeRes 获取{{.Comment}}树形结构响应
type {{.ModelName}}TreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.{{.ModelName}}TreeOutput `json:"list" dc:"树形数据"`
}
{{end}}
