package service

import (
	"context"
	"gbaseadmin/app/{{.AppName}}/internal/model"
	"gbaseadmin/utility/snowflake"
)

type I{{.ModelName}} interface {
	Create(ctx context.Context, in *model.{{.ModelName}}CreateInput) error
	Update(ctx context.Context, in *model.{{.ModelName}}UpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.{{.ModelName}}DetailOutput, err error)
	List(ctx context.Context, in *model.{{.ModelName}}ListInput) (list []*model.{{.ModelName}}ListOutput, total int, err error)
{{- if .HasParentID}}
	Tree(ctx context.Context) (tree []*model.{{.ModelName}}TreeOutput, err error)
{{- end}}
}

var local{{.ModelName}} I{{.ModelName}}

func {{.ModelName}}() I{{.ModelName}} {
	return local{{.ModelName}}
}

func Register{{.ModelName}}(i I{{.ModelName}}) {
	local{{.ModelName}} = i
}
