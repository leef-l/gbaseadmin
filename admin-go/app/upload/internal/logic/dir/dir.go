package dir

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/upload/internal/dao"
	"gbaseadmin/app/upload/internal/model"
	"gbaseadmin/app/upload/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterDir(New())
}

func New() *sDir {
	return &sDir{}
}

type sDir struct{}

// Create 创建æ–‡ä»¶ç›®å½•
func (s *sDir) Create(ctx context.Context, in *model.DirCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.UploadDir.Ctx(ctx).Data(g.Map{
		dao.UploadDir.Columns().Id:        id,
		dao.UploadDir.Columns().ParentId: in.ParentID,
		dao.UploadDir.Columns().Name: in.Name,
		dao.UploadDir.Columns().Path: in.Path,
		dao.UploadDir.Columns().Sort: in.Sort,
		dao.UploadDir.Columns().Status: in.Status,
		dao.UploadDir.Columns().CreatedAt: gtime.Now(),
		dao.UploadDir.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新æ–‡ä»¶ç›®å½•
func (s *sDir) Update(ctx context.Context, in *model.DirUpdateInput) error {
	data := g.Map{
		dao.UploadDir.Columns().ParentId: in.ParentID,
		dao.UploadDir.Columns().Name: in.Name,
		dao.UploadDir.Columns().Path: in.Path,
		dao.UploadDir.Columns().Sort: in.Sort,
		dao.UploadDir.Columns().Status: in.Status,
		dao.UploadDir.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.UploadDir.Ctx(ctx).Where(dao.UploadDir.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除æ–‡ä»¶ç›®å½•
func (s *sDir) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.UploadDir.Ctx(ctx).Where(dao.UploadDir.Columns().Id, id).Data(g.Map{
		dao.UploadDir.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取æ–‡ä»¶ç›®å½•详情
func (s *sDir) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.DirDetailOutput, err error) {
	out = &model.DirDetailOutput{}
	err = dao.UploadDir.Ctx(ctx).Where(dao.UploadDir.Columns().Id, id).Where(dao.UploadDir.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询ä¸Šçº§ç›®å½•关联显示
	if out.ParentID != 0 {
		val, err := g.DB().Ctx(ctx).Model("upload_dir").Where("id", out.ParentID).Where("deleted_at", nil).Value("name")
		if err == nil {
			out.DirName = val.String()
		}
	}
	return
}

// List 获取æ–‡ä»¶ç›®å½•列表
func (s *sDir) List(ctx context.Context, in *model.DirListInput) (list []*model.DirListOutput, total int, err error) {
	m := dao.UploadDir.Ctx(ctx).Where(dao.UploadDir.Columns().DeletedAt, nil)
	if in.Status > 0 {
		m = m.Where(dao.UploadDir.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.UploadDir.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.ParentID != 0 {
			val, err := g.DB().Ctx(ctx).Model("upload_dir").Where("id", item.ParentID).Where("deleted_at", nil).Value("name")
			if err == nil {
				item.DirName = val.String()
			}
		}
	}
	return
}

// Tree 获取æ–‡ä»¶ç›®å½•树形结构
func (s *sDir) Tree(ctx context.Context) (tree []*model.DirTreeOutput, err error) {
	var list []*model.DirTreeOutput
	err = dao.UploadDir.Ctx(ctx).Where(dao.UploadDir.Columns().DeletedAt, nil).OrderAsc(dao.UploadDir.Columns().Sort).Scan(&list)
	if err != nil {
		return
	}

	// 使用 map 迭代方式组装树
	nodeMap := make(map[int64]*model.DirTreeOutput, len(list))
	for _, item := range list {
		item.Children = make([]*model.DirTreeOutput, 0)
		nodeMap[int64(item.ID)] = item
	}

	tree = make([]*model.DirTreeOutput, 0)
	for _, item := range list {
		if int64(item.ParentID) == 0 {
			tree = append(tree, item)
		} else if parent, ok := nodeMap[int64(item.ParentID)]; ok {
			parent.Children = append(parent.Children, item)
		}
	}
	return
}

