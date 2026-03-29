package {{.ModuleName}}

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/system/internal/dao"
	"gbaseadmin/app/system/internal/model"
	"gbaseadmin/app/system/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.Register{{.ModelName}}(New())
}

func New() *s{{.ModelName}} {
	return &s{{.ModelName}}{}
}

type s{{.ModelName}} struct{}

// Create 创建{{.Comment}}
func (s *s{{.ModelName}}) Create(ctx context.Context, in *model.{{.ModelName}}CreateInput) error {
	id := snowflake.Generate()
	_, err := dao.{{.ModelName}}.Ctx(ctx).Data(g.Map{
		dao.{{.ModelName}}.Columns().Id:        id,
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden)}}
		dao.{{.ModelName}}.Columns().{{.NameCamel}}: in.{{.NameCamel}},
{{- end}}
{{- end}}
		dao.{{.ModelName}}.Columns().CreatedAt: gtime.Now(),
		dao.{{.ModelName}}.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新{{.Comment}}
func (s *s{{.ModelName}}) Update(ctx context.Context, in *model.{{.ModelName}}UpdateInput) error {
	_, err := dao.{{.ModelName}}.Ctx(ctx).Where(dao.{{.ModelName}}.Columns().Id, in.Id).Data(g.Map{
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden)}}
		dao.{{.ModelName}}.Columns().{{.NameCamel}}: in.{{.NameCamel}},
{{- end}}
{{- end}}
		dao.{{.ModelName}}.Columns().UpdatedAt: gtime.Now(),
	}).Update()
	return err
}

// Delete 删除{{.Comment}}
func (s *s{{.ModelName}}) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.{{.ModelName}}.Ctx(ctx).Where(dao.{{.ModelName}}.Columns().Id, id).Delete()
	return err
}

// Detail 获取{{.Comment}}详情
func (s *s{{.ModelName}}) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.{{.ModelName}}DetailOutput, err error) {
	out = &model.{{.ModelName}}DetailOutput{}
	err = dao.{{.ModelName}}.Ctx(ctx).Where(dao.{{.ModelName}}.Columns().Id, id).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取{{.Comment}}列表
func (s *s{{.ModelName}}) List(ctx context.Context, in *model.{{.ModelName}}ListInput) (list []*model.{{.ModelName}}ListOutput, total int, err error) {
	m := dao.{{.ModelName}}.Ctx(ctx)
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.{{.ModelName}}.Columns().Id).Scan(&list)
	return
}
{{if .HasParentID}}
// Tree 获取{{.Comment}}树形结构
func (s *s{{.ModelName}}) Tree(ctx context.Context) (tree []*model.{{.ModelName}}TreeOutput, err error) {
	var list []*model.{{.ModelName}}TreeOutput
	err = dao.{{.ModelName}}.Ctx(ctx).{{if .HasSort}}OrderAsc(dao.{{.ModelName}}.Columns().Sort).{{end}}Scan(&list)
	if err != nil {
		return
	}

	// 使用 map 迭代方式组装树
	nodeMap := make(map[int64]*model.{{.ModelName}}TreeOutput, len(list))
	for _, item := range list {
		nodeMap[int64(item.Id)] = item
	}

	tree = make([]*model.{{.ModelName}}TreeOutput, 0)
	for _, item := range list {
		if int64(item.ParentId) == 0 {
			tree = append(tree, item)
		} else if parent, ok := nodeMap[int64(item.ParentId)]; ok {
			parent.Children = append(parent.Children, item)
		}
	}
	return
}
{{end}}
