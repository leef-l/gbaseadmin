package role

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/system/internal/dao"
	"gbaseadmin/app/system/internal/model"
	"gbaseadmin/app/system/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

type sRole struct{}

// Create 创建角色表
func (s *sRole) Create(ctx context.Context, in *model.RoleCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.Role.Ctx(ctx).Data(g.Map{
		dao.Role.Columns().Id:        id,
		dao.Role.Columns().ParentId: in.ParentID,
		dao.Role.Columns().Title: in.Title,
		dao.Role.Columns().DataScope: in.DataScope,
		dao.Role.Columns().Sort: in.Sort,
		dao.Role.Columns().Status: in.Status,
		dao.Role.Columns().CreatedAt: gtime.Now(),
		dao.Role.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新角色表
func (s *sRole) Update(ctx context.Context, in *model.RoleUpdateInput) error {
	data := g.Map{
		dao.Role.Columns().ParentId: in.ParentID,
		dao.Role.Columns().Title: in.Title,
		dao.Role.Columns().DataScope: in.DataScope,
		dao.Role.Columns().Sort: in.Sort,
		dao.Role.Columns().Status: in.Status,
		dao.Role.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除角色表
func (s *sRole) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, id).Data(g.Map{
		dao.Role.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取角色表详情
func (s *sRole) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.RoleDetailOutput, err error) {
	out = &model.RoleDetailOutput{}
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, id).Where(dao.Role.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询上级角色ID，0 表示顶级角色关联显示
	if out.ParentID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("role").Where("id", out.ParentID).Where("deleted_at", nil).Value("title")
		out.RoleTitle = val.String()
	}
	return
}

// List 获取角色表列表
func (s *sRole) List(ctx context.Context, in *model.RoleListInput) (list []*model.RoleListOutput, total int, err error) {
	m := dao.Role.Ctx(ctx).Where(dao.Role.Columns().DeletedAt, nil)
	if in.DataScope > 0 {
		m = m.Where(dao.Role.Columns().DataScope, in.DataScope)
	}
	if in.Status > 0 {
		m = m.Where(dao.Role.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.Role.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.ParentID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("role").Where("id", item.ParentID).Where("deleted_at", nil).Value("title")
			item.RoleTitle = val.String()
		}
	}
	return
}

// Tree 获取角色表树形结构
func (s *sRole) Tree(ctx context.Context) (tree []*model.RoleTreeOutput, err error) {
	var list []*model.RoleTreeOutput
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().DeletedAt, nil).OrderAsc(dao.Role.Columns().Sort).Scan(&list)
	if err != nil {
		return
	}

	// 使用 map 迭代方式组装树
	nodeMap := make(map[int64]*model.RoleTreeOutput, len(list))
	for _, item := range list {
		item.Children = make([]*model.RoleTreeOutput, 0)
		nodeMap[int64(item.ID)] = item
	}

	tree = make([]*model.RoleTreeOutput, 0)
	for _, item := range list {
		if int64(item.ParentID) == 0 {
			tree = append(tree, item)
		} else if parent, ok := nodeMap[int64(item.ParentID)]; ok {
			parent.Children = append(parent.Children, item)
		}
	}
	return
}

