package sys_test

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
	service.RegisterSysTest(New())
}

func New() *sSysTest {
	return &sSysTest{}
}

type sSysTest struct{}

// Create 创建æµ‹è¯•è¡¨
func (s *sSysTest) Create(ctx context.Context, in *model.SysTestCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.SysTest.Ctx(ctx).Data(g.Map{
		dao.SysTest.Columns().Id:        id,
		dao.SysTest.Columns().ParentId: in.ParentID,
		dao.SysTest.Columns().Title: in.Title,
		dao.SysTest.Columns().Code: in.Code,
		dao.SysTest.Columns().Type: in.Type,
		dao.SysTest.Columns().Status: in.Status,
		dao.SysTest.Columns().Sort: in.Sort,
		dao.SysTest.Columns().Remark: in.Remark,
		dao.SysTest.Columns().CreatedAt: gtime.Now(),
		dao.SysTest.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新æµ‹è¯•è¡¨
func (s *sSysTest) Update(ctx context.Context, in *model.SysTestUpdateInput) error {
	data := g.Map{
		dao.SysTest.Columns().ParentId: in.ParentID,
		dao.SysTest.Columns().Title: in.Title,
		dao.SysTest.Columns().Code: in.Code,
		dao.SysTest.Columns().Type: in.Type,
		dao.SysTest.Columns().Status: in.Status,
		dao.SysTest.Columns().Sort: in.Sort,
		dao.SysTest.Columns().Remark: in.Remark,
		dao.SysTest.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.SysTest.Ctx(ctx).Where(dao.SysTest.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除æµ‹è¯•è¡¨
func (s *sSysTest) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.SysTest.Ctx(ctx).Where(dao.SysTest.Columns().Id, id).Data(g.Map{
		dao.SysTest.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取æµ‹è¯•è¡¨详情
func (s *sSysTest) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.SysTestDetailOutput, err error) {
	out = &model.SysTestDetailOutput{}
	err = dao.SysTest.Ctx(ctx).Where(dao.SysTest.Columns().Id, id).Where(dao.SysTest.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询ä¸Šçº§IDï¼Œ0è¡¨ç¤ºé¡¶çº§关联显示
	if out.ParentID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("sys_test").Where("id", out.ParentID).Where("deleted_at", nil).Value("title")
		out.SysTestTitle = val.String()
	}
	return
}

// List 获取æµ‹è¯•è¡¨列表
func (s *sSysTest) List(ctx context.Context, in *model.SysTestListInput) (list []*model.SysTestListOutput, total int, err error) {
	m := dao.SysTest.Ctx(ctx).Where(dao.SysTest.Columns().DeletedAt, nil)
	if in.Type > 0 {
		m = m.Where(dao.SysTest.Columns().Type, in.Type)
	}
	if in.Status > 0 {
		m = m.Where(dao.SysTest.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.SysTest.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.ParentID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("sys_test").Where("id", item.ParentID).Where("deleted_at", nil).Value("title")
			item.SysTestTitle = val.String()
		}
	}
	return
}

// Tree 获取æµ‹è¯•è¡¨树形结构
func (s *sSysTest) Tree(ctx context.Context) (tree []*model.SysTestTreeOutput, err error) {
	var list []*model.SysTestTreeOutput
	err = dao.SysTest.Ctx(ctx).Where(dao.SysTest.Columns().DeletedAt, nil).OrderAsc(dao.SysTest.Columns().Sort).Scan(&list)
	if err != nil {
		return
	}

	// 使用 map 迭代方式组装树
	nodeMap := make(map[int64]*model.SysTestTreeOutput, len(list))
	for _, item := range list {
		item.Children = make([]*model.SysTestTreeOutput, 0)
		nodeMap[int64(item.ID)] = item
	}

	tree = make([]*model.SysTestTreeOutput, 0)
	for _, item := range list {
		if int64(item.ParentID) == 0 {
			tree = append(tree, item)
		} else if parent, ok := nodeMap[int64(item.ParentID)]; ok {
			parent.Children = append(parent.Children, item)
		}
	}
	return
}

