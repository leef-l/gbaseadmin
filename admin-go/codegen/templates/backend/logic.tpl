package {{.ModuleName}}

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
{{- if .HasPassword}}
	"golang.org/x/crypto/bcrypt"
{{- end}}

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
{{- range .Fields}}
{{- if .IsPassword}}
	hashed{{.NameCamel}}, err := bcrypt.GenerateFromPassword([]byte(in.{{.NameCamel}}), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
{{- end}}
{{- end}}
	_, err{{if not .HasPassword}} :={{else}} ={{end}} dao.{{.ModelName}}.Ctx(ctx).Data(g.Map{
		dao.{{.ModelName}}.Columns().Id:        id,
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden)}}
{{- if .IsPassword}}
		dao.{{$.ModelName}}.Columns().{{.NameDao}}: string(hashed{{.NameCamel}}),
{{- else}}
		dao.{{$.ModelName}}.Columns().{{.NameDao}}: in.{{.NameCamel}},
{{- end}}
{{- end}}
{{- end}}
		dao.{{.ModelName}}.Columns().CreatedAt: gtime.Now(),
		dao.{{.ModelName}}.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新{{.Comment}}
func (s *s{{.ModelName}}) Update(ctx context.Context, in *model.{{.ModelName}}UpdateInput) error {
	data := g.Map{
{{- range .Fields}}
{{- if and (not .IsID) (not .IsHidden) (not .IsPassword)}}
		dao.{{$.ModelName}}.Columns().{{.NameDao}}: in.{{.NameCamel}},
{{- end}}
{{- end}}
		dao.{{.ModelName}}.Columns().UpdatedAt: gtime.Now(),
	}
{{- range .Fields}}
{{- if .IsPassword}}
	if in.{{.NameCamel}} != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(in.{{.NameCamel}}), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		data[dao.{{$.ModelName}}.Columns().{{.NameDao}}] = string(hashed)
	}
{{- end}}
{{- end}}
	_, err := dao.{{.ModelName}}.Ctx(ctx).Where(dao.{{.ModelName}}.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除{{.Comment}}
func (s *s{{.ModelName}}) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.{{.ModelName}}.Ctx(ctx).Where(dao.{{.ModelName}}.Columns().Id, id).Data(g.Map{
		dao.{{.ModelName}}.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取{{.Comment}}详情
func (s *s{{.ModelName}}) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.{{.ModelName}}DetailOutput, err error) {
	out = &model.{{.ModelName}}DetailOutput{}
	err = dao.{{.ModelName}}.Ctx(ctx).Where(dao.{{.ModelName}}.Columns().Id, id).Where(dao.{{.ModelName}}.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
{{- range .Fields}}
{{- if and .RefFieldName (not .IsHidden)}}
	// 查询{{.Label}}关联显示
	if out.{{.NameCamel}} != 0 {
		val, _ := g.DB().Ctx(ctx).Model("{{.RefTable}}").Where("id", out.{{.NameCamel}}).Where("deleted_at", nil).Value("{{.RefDisplayField}}")
		out.{{.RefFieldName}} = val.String()
	}
{{- end}}
{{- end}}
	return
}

// List 获取{{.Comment}}列表
func (s *s{{.ModelName}}) List(ctx context.Context, in *model.{{.ModelName}}ListInput) (list []*model.{{.ModelName}}ListOutput, total int, err error) {
	m := dao.{{.ModelName}}.Ctx(ctx).Where(dao.{{.ModelName}}.Columns().DeletedAt, nil)
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
	if in.{{.NameCamel}} > 0 {
		m = m.Where(dao.{{$.ModelName}}.Columns().{{.NameDao}}, in.{{.NameCamel}})
	}
{{- end}}
{{- end}}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.{{.ModelName}}.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
{{- $hasRef := false}}
{{- range .Fields}}
{{- if and .RefFieldName (not .IsHidden)}}
{{- $hasRef = true}}
{{- end}}
{{- end}}
{{- if $hasRef}}
	// 填充关联显示字段
	for _, item := range list {
{{- range .Fields}}
{{- if and .RefFieldName (not .IsHidden)}}
		if item.{{.NameCamel}} != 0 {
			val, _ := g.DB().Ctx(ctx).Model("{{.RefTable}}").Where("id", item.{{.NameCamel}}).Where("deleted_at", nil).Value("{{.RefDisplayField}}")
			item.{{.RefFieldName}} = val.String()
		}
{{- end}}
{{- end}}
	}
{{- end}}
	return
}
{{if .HasParentID}}
// Tree 获取{{.Comment}}树形结构
func (s *s{{.ModelName}}) Tree(ctx context.Context) (tree []*model.{{.ModelName}}TreeOutput, err error) {
	var list []*model.{{.ModelName}}TreeOutput
	err = dao.{{.ModelName}}.Ctx(ctx).Where(dao.{{.ModelName}}.Columns().DeletedAt, nil).{{if .HasSort}}OrderAsc(dao.{{.ModelName}}.Columns().Sort).{{end}}Scan(&list)
	if err != nil {
		return
	}

	// 使用 map 迭代方式组装树
	nodeMap := make(map[int64]*model.{{.ModelName}}TreeOutput, len(list))
	for _, item := range list {
		item.Children = make([]*model.{{.ModelName}}TreeOutput, 0)
		nodeMap[int64(item.ID)] = item
	}

	tree = make([]*model.{{.ModelName}}TreeOutput, 0)
	for _, item := range list {
		if int64(item.ParentID) == 0 {
			tree = append(tree, item)
		} else if parent, ok := nodeMap[int64(item.ParentID)]; ok {
			parent.Children = append(parent.Children, item)
		}
	}
	return
}
{{end}}
