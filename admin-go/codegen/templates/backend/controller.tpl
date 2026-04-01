package {{.PackageName}}

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

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

// BatchDelete 批量删除{{.Comment}}
func (c *c{{.ModelName}}) BatchDelete(ctx context.Context, req *v1.{{.ModelName}}BatchDeleteReq) (res *v1.{{.ModelName}}BatchDeleteRes, err error) {
	err = service.{{.ModelName}}().BatchDelete(ctx, req.IDs)
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
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
		OrderBy:   req.OrderBy,
		OrderDir:  req.OrderDir,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
		{{.NameCamel}}: req.{{.NameCamel}},
{{- end}}
{{- end}}
{{- range .Fields}}
{{- if .IsSearchable}}
		{{.NameCamel}}: req.{{.NameCamel}},
{{- end}}
{{- end}}
	})
	return
}
// Export 导出{{.Comment}}
func (c *c{{.ModelName}}) Export(ctx context.Context, req *v1.{{.ModelName}}ExportReq) (res *v1.{{.ModelName}}ExportRes, err error) {
	list, err := service.{{.ModelName}}().Export(ctx, &model.{{.ModelName}}ListInput{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
		{{.NameCamel}}: req.{{.NameCamel}},
{{- end}}
{{- end}}
{{- range .Fields}}
{{- if .IsSearchable}}
		{{.NameCamel}}: req.{{.NameCamel}},
{{- end}}
{{- end}}
	})
	if err != nil {
		return
	}
	// CSV 导出
	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", "text/csv; charset=utf-8")
	r.Response.Header().Set("Content-Disposition", `attachment; filename="{{.ModuleName}}.csv"`)
	r.Response.Write("\xEF\xBB\xBF") // UTF-8 BOM
	// 表头
	r.Response.Writeln("{{- $first := true}}{{- range .Fields}}{{- if and (not .IsHidden) (not .IsID) (not .IsPassword)}}{{if not $first}},{{end}}{{if .RefFieldJSON}}{{.ShortLabel}}{{else}}{{.ShortLabel}}{{end}}{{$first = false}}{{- end}}{{- end}},创建时间")
	// 数据行
	for _, item := range list {
		r.Response.Writefln("{{- $first := true}}{{- range .Fields}}{{- if and (not .IsHidden) (not .IsID) (not .IsPassword)}}{{if not $first}},{{end}}%v{{$first = false}}{{- end}}{{- end}},%v",
{{- $firstArg := true}}
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (not .IsPassword)}}
{{- if .RefFieldJSON}}
			{{if not $firstArg}} {{end}}item.{{.RefFieldName}},
{{- else}}
			{{if not $firstArg}} {{end}}item.{{.NameCamel}},
{{- end}}
{{- $firstArg = false}}
{{- end}}
{{- end}}
			item.CreatedAt,
		)
	}
	return
}
{{if .HasParentID}}
// Tree 获取{{.Comment}}树形结构
func (c *c{{.ModelName}}) Tree(ctx context.Context, req *v1.{{.ModelName}}TreeReq) (res *v1.{{.ModelName}}TreeRes, err error) {
	res = &v1.{{.ModelName}}TreeRes{}
	res.List, err = service.{{.ModelName}}().Tree(ctx, &model.{{.ModelName}}TreeInput{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (not .IsParentID) (.IsEnum)}}
		{{.NameCamel}}: req.{{.NameCamel}},
{{- end}}
{{- end}}
{{- range .Fields}}
{{- if .IsSearchable}}
		{{.NameCamel}}: req.{{.NameCamel}},
{{- end}}
{{- end}}
	})
	return
}
{{end}}
