package dept

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/system/internal/dao"
	"gbaseadmin/app/system/internal/model"
	"gbaseadmin/app/system/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterDept(New())
}

func New() *sDept {
	return &sDept{}
}

type sDept struct{}

// Create 创建部门表
func (s *sDept) Create(ctx context.Context, in *model.DeptCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.Dept.Ctx(ctx).Data(g.Map{
		dao.Dept.Columns().Id:        id,
		dao.Dept.Columns().ParentID: in.ParentID,
		dao.Dept.Columns().Title: in.Title,
		dao.Dept.Columns().Username: in.Username,
		dao.Dept.Columns().Email: in.Email,
		dao.Dept.Columns().Sort: in.Sort,
		dao.Dept.Columns().Status: in.Status,
		dao.Dept.Columns().CreatedAt: gtime.Now(),
		dao.Dept.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新部门表
func (s *sDept) Update(ctx context.Context, in *model.DeptUpdateInput) error {
	_, err := dao.Dept.Ctx(ctx).Where(dao.Dept.Columns().Id, in.Id).Data(g.Map{
		dao.Dept.Columns().ParentID: in.ParentID,
		dao.Dept.Columns().Title: in.Title,
		dao.Dept.Columns().Username: in.Username,
		dao.Dept.Columns().Email: in.Email,
		dao.Dept.Columns().Sort: in.Sort,
		dao.Dept.Columns().Status: in.Status,
		dao.Dept.Columns().UpdatedAt: gtime.Now(),
	}).Update()
	return err
}

// Delete 删除部门表
func (s *sDept) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.Dept.Ctx(ctx).Where(dao.Dept.Columns().Id, id).Delete()
	return err
}

// Detail 获取部门表详情
func (s *sDept) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.DeptDetailOutput, err error) {
	out = &model.DeptDetailOutput{}
	err = dao.Dept.Ctx(ctx).Where(dao.Dept.Columns().Id, id).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取部门表列表
func (s *sDept) List(ctx context.Context, in *model.DeptListInput) (list []*model.DeptListOutput, total int, err error) {
	m := dao.Dept.Ctx(ctx)
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.Dept.Columns().Id).Scan(&list)
	return
}

// Tree 获取部门表树形结构
func (s *sDept) Tree(ctx context.Context) (tree []*model.DeptTreeOutput, err error) {
	var list []*model.DeptTreeOutput
	err = dao.Dept.Ctx(ctx).OrderAsc(dao.Dept.Columns().Sort).Scan(&list)
	if err != nil {
		return
	}

	// 使用 map 迭代方式组装树
	nodeMap := make(map[int64]*model.DeptTreeOutput, len(list))
	for _, item := range list {
		nodeMap[int64(item.Id)] = item
	}

	tree = make([]*model.DeptTreeOutput, 0)
	for _, item := range list {
		if int64(item.ParentId) == 0 {
			tree = append(tree, item)
		} else if parent, ok := nodeMap[int64(item.ParentId)]; ok {
			parent.Children = append(parent.Children, item)
		}
	}
	return
}

