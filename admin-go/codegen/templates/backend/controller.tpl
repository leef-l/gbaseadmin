package {{.PackageName}}

import (
	"context"

	v1 "gbaseadmin/app/{{.AppName}}/api/{{.AppName}}/v1"
	"gbaseadmin/app/{{.AppName}}/internal/model"
	"gbaseadmin/app/{{.AppName}}/internal/service"
)

var {{.ModelName}} = c{{.ModelName}}{}

type c{{.ModelName}} struct{}

// Create 创建{{.Comment}}
func (c *c{{.ModelName}}) Create(ctx context.Context, req *v1.{{.ModelName}}CreateReq) (res *v1.{{.ModelName}}CreateRes, err error) {
	err = service.{{.ModelName}}().Create(ctx, &model.{{.ModelName}}CreateInput{
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden)}}
		{{.NameCamel}}: req.{{.NameCamel}},
{{- end}}
{{- end}}
	})
	return
}

// Update 更新{{.Comment}}
func (c *c{{.ModelName}}) Update(ctx context.Context, req *v1.{{.ModelName}}UpdateReq) (res *v1.{{.ModelName}}UpdateRes, err error) {
	err = service.{{.ModelName}}().Update(ctx, &model.{{.ModelName}}UpdateInput{
		ID: req.ID,
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden)}}
		{{.NameCamel}}: req.{{.NameCamel}},
{{- end}}
{{- end}}
	})
	return
}

// Delete 删除{{.Comment}}
func (c *c{{.ModelName}}) Delete(ctx context.Context, req *v1.{{.ModelName}}DeleteReq) (res *v1.{{.ModelName}}DeleteRes, err error) {
	err = service.{{.ModelName}}().Delete(ctx, req.ID)
	return
}

// Detail 获取{{.Comment}}详情
func (c *c{{.ModelName}}) Detail(ctx context.Context, req *v1.{{.ModelName}}DetailReq) (res *v1.{{.ModelName}}DetailRes, err error) {
	res = &v1.{{.ModelName}}DetailRes{}
	res.{{.ModelName}}DetailOutput, err = service.{{.ModelName}}().Detail(ctx, req.ID)
	return
}

// List 获取{{.Comment}}列表
func (c *c{{.ModelName}}) List(ctx context.Context, req *v1.{{.ModelName}}ListReq) (res *v1.{{.ModelName}}ListRes, err error) {
	res = &v1.{{.ModelName}}ListRes{}
	res.List, res.Total, err = service.{{.ModelName}}().List(ctx, &model.{{.ModelName}}ListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
		{{.NameCamel}}: req.{{.NameCamel}},
{{- end}}
{{- end}}
	})
	return
}
{{if .HasParentID}}
// Tree 获取{{.Comment}}树形结构
func (c *c{{.ModelName}}) Tree(ctx context.Context, req *v1.{{.ModelName}}TreeReq) (res *v1.{{.ModelName}}TreeRes, err error) {
	res = &v1.{{.ModelName}}TreeRes{}
	res.List, err = service.{{.ModelName}}().Tree(ctx)
	return
}
{{end}}
