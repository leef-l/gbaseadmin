package role

import (
	"context"

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
		dao.Role.Columns().ParentID: in.ParentID,
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
	_, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.Id).Data(g.Map{
		dao.Role.Columns().ParentID: in.ParentID,
		dao.Role.Columns().Title: in.Title,
		dao.Role.Columns().DataScope: in.DataScope,
		dao.Role.Columns().Sort: in.Sort,
		dao.Role.Columns().Status: in.Status,
		dao.Role.Columns().UpdatedAt: gtime.Now(),
	}).Update()
	return err
}

// Delete 删除角色表
func (s *sRole) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, id).Delete()
	return err
}

// Detail 获取角色表详情
func (s *sRole) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.RoleDetailOutput, err error) {
	out = &model.RoleDetailOutput{}
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, id).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取角色表列表
func (s *sRole) List(ctx context.Context, in *model.RoleListInput) (list []*model.RoleListOutput, total int, err error) {
	m := dao.Role.Ctx(ctx)
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.Role.Columns().Id).Scan(&list)
	return
}

// Tree 获取角色表树形结构
func (s *sRole) Tree(ctx context.Context) (tree []*model.RoleTreeOutput, err error) {
	var list []*model.RoleTreeOutput
	err = dao.Role.Ctx(ctx).OrderAsc(dao.Role.Columns().Sort).Scan(&list)
	if err != nil {
		return
	}

	// 使用 map 迭代方式组装树
	nodeMap := make(map[int64]*model.RoleTreeOutput, len(list))
	for _, item := range list {
		nodeMap[int64(item.Id)] = item
	}

	tree = make([]*model.RoleTreeOutput, 0)
	for _, item := range list {
		if int64(item.ParentId) == 0 {
			tree = append(tree, item)
		} else if parent, ok := nodeMap[int64(item.ParentId)]; ok {
			parent.Children = append(parent.Children, item)
		}
	}
	return
}

