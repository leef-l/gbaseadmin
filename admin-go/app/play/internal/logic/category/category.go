package category

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

type sCategory struct{}

// Create 创建商品分类表
func (s *sCategory) Create(ctx context.Context, in *model.CategoryCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayCategory.Ctx(ctx).Data(g.Map{
		dao.PlayCategory.Columns().Id:        id,
		dao.PlayCategory.Columns().ParentId: in.ParentID,
		dao.PlayCategory.Columns().Title: in.Title,
		dao.PlayCategory.Columns().Icon: in.Icon,
		dao.PlayCategory.Columns().CoverImage: in.CoverImage,
		dao.PlayCategory.Columns().Sort: in.Sort,
		dao.PlayCategory.Columns().Status: in.Status,
		dao.PlayCategory.Columns().CreatedAt: gtime.Now(),
		dao.PlayCategory.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新商品分类表
func (s *sCategory) Update(ctx context.Context, in *model.CategoryUpdateInput) error {
	data := g.Map{
		dao.PlayCategory.Columns().ParentId: in.ParentID,
		dao.PlayCategory.Columns().Title: in.Title,
		dao.PlayCategory.Columns().Icon: in.Icon,
		dao.PlayCategory.Columns().CoverImage: in.CoverImage,
		dao.PlayCategory.Columns().Sort: in.Sort,
		dao.PlayCategory.Columns().Status: in.Status,
		dao.PlayCategory.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayCategory.Ctx(ctx).Where(dao.PlayCategory.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除商品分类表
func (s *sCategory) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayCategory.Ctx(ctx).Where(dao.PlayCategory.Columns().Id, id).Data(g.Map{
		dao.PlayCategory.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取商品分类表详情
func (s *sCategory) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CategoryDetailOutput, err error) {
	out = &model.CategoryDetailOutput{}
	err = dao.PlayCategory.Ctx(ctx).Where(dao.PlayCategory.Columns().Id, id).Where(dao.PlayCategory.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询上级分类ID，0 表示顶级分类关联显示
	if out.ParentID != 0 {
		val, err := g.DB().Ctx(ctx).Model("play_category").Where("id", out.ParentID).Where("deleted_at", nil).Value("title")
		if err == nil {
			out.CategoryTitle = val.String()
		}
	}
	return
}

// List 获取商品分类表列表
func (s *sCategory) List(ctx context.Context, in *model.CategoryListInput) (list []*model.CategoryListOutput, total int, err error) {
	m := dao.PlayCategory.Ctx(ctx).Where(dao.PlayCategory.Columns().DeletedAt, nil)
	if in.Status > 0 {
		m = m.Where(dao.PlayCategory.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayCategory.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.ParentID != 0 {
			val, err := g.DB().Ctx(ctx).Model("play_category").Where("id", item.ParentID).Where("deleted_at", nil).Value("title")
			if err == nil {
				item.CategoryTitle = val.String()
			}
		}
	}
	return
}

// Tree 获取商品分类表树形结构
func (s *sCategory) Tree(ctx context.Context) (tree []*model.CategoryTreeOutput, err error) {
	var list []*model.CategoryTreeOutput
	err = dao.PlayCategory.Ctx(ctx).Where(dao.PlayCategory.Columns().DeletedAt, nil).OrderAsc(dao.PlayCategory.Columns().Sort).Scan(&list)
	if err != nil {
		return
	}

	// 使用 map 迭代方式组装树
	nodeMap := make(map[int64]*model.CategoryTreeOutput, len(list))
	for _, item := range list {
		item.Children = make([]*model.CategoryTreeOutput, 0)
		nodeMap[int64(item.ID)] = item
	}

	tree = make([]*model.CategoryTreeOutput, 0)
	for _, item := range list {
		if int64(item.ParentID) == 0 {
			tree = append(tree, item)
		} else if parent, ok := nodeMap[int64(item.ParentID)]; ok {
			parent.Children = append(parent.Children, item)
		}
	}
	return
}

