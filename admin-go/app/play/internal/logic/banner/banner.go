package banner

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
	service.RegisterBanner(New())
}

func New() *sBanner {
	return &sBanner{}
}

type sBanner struct{}

// Create 创建首页Banner轮播
func (s *sBanner) Create(ctx context.Context, in *model.BannerCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayBanner.Ctx(ctx).Data(g.Map{
		dao.PlayBanner.Columns().Id:        id,
		dao.PlayBanner.Columns().Title: in.Title,
		dao.PlayBanner.Columns().Image: in.Image,
		dao.PlayBanner.Columns().LinkType: in.LinkType,
		dao.PlayBanner.Columns().LinkValue: in.LinkValue,
		dao.PlayBanner.Columns().Sort: in.Sort,
		dao.PlayBanner.Columns().Status: in.Status,
		dao.PlayBanner.Columns().StartTime: in.StartTime,
		dao.PlayBanner.Columns().EndTime: in.EndTime,
		dao.PlayBanner.Columns().Remark: in.Remark,
		dao.PlayBanner.Columns().CreatedAt: gtime.Now(),
		dao.PlayBanner.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新首页Banner轮播
func (s *sBanner) Update(ctx context.Context, in *model.BannerUpdateInput) error {
	data := g.Map{
		dao.PlayBanner.Columns().Title: in.Title,
		dao.PlayBanner.Columns().Image: in.Image,
		dao.PlayBanner.Columns().LinkType: in.LinkType,
		dao.PlayBanner.Columns().LinkValue: in.LinkValue,
		dao.PlayBanner.Columns().Sort: in.Sort,
		dao.PlayBanner.Columns().Status: in.Status,
		dao.PlayBanner.Columns().StartTime: in.StartTime,
		dao.PlayBanner.Columns().EndTime: in.EndTime,
		dao.PlayBanner.Columns().Remark: in.Remark,
		dao.PlayBanner.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayBanner.Ctx(ctx).Where(dao.PlayBanner.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除首页Banner轮播
func (s *sBanner) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayBanner.Ctx(ctx).Where(dao.PlayBanner.Columns().Id, id).Data(g.Map{
		dao.PlayBanner.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// BatchDelete 批量软删除首页Banner轮播
func (s *sBanner) BatchDelete(ctx context.Context, ids []snowflake.JsonInt64) error {
	_, err := dao.PlayBanner.Ctx(ctx).WhereIn(dao.PlayBanner.Columns().Id, ids).Data(g.Map{
		dao.PlayBanner.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取首页Banner轮播详情
func (s *sBanner) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.BannerDetailOutput, err error) {
	out = &model.BannerDetailOutput{}
	err = dao.PlayBanner.Ctx(ctx).Where(dao.PlayBanner.Columns().Id, id).Where(dao.PlayBanner.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取首页Banner轮播列表
func (s *sBanner) List(ctx context.Context, in *model.BannerListInput) (list []*model.BannerListOutput, total int, err error) {
	m := dao.PlayBanner.Ctx(ctx).Where(dao.PlayBanner.Columns().DeletedAt, nil)
	if in.Title != "" {
		m = m.WhereLike(dao.PlayBanner.Columns().Title, "%"+in.Title+"%")
	}
	if in.Remark != "" {
		m = m.WhereLike(dao.PlayBanner.Columns().Remark, "%"+in.Remark+"%")
	}
	// 时间范围筛选
	if in.StartTime != "" {
		m = m.WhereGTE(dao.PlayBanner.Columns().CreatedAt, in.StartTime)
	}
	if in.EndTime != "" {
		m = m.WhereLTE(dao.PlayBanner.Columns().CreatedAt, in.EndTime)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	// 动态排序
	if in.OrderBy != "" {
		if in.OrderDir == "desc" {
			m = m.OrderDesc(in.OrderBy)
		} else {
			m = m.OrderAsc(in.OrderBy)
		}
	} else {
		m = m.OrderAsc(dao.PlayBanner.Columns().Id)
	}
	err = m.Page(in.PageNum, in.PageSize).Scan(&list)
	if err != nil {
		return
	}
	return
}
// Export 导出首页Banner轮播（不分页）
func (s *sBanner) Export(ctx context.Context, in *model.BannerListInput) (list []*model.BannerListOutput, err error) {
	m := dao.PlayBanner.Ctx(ctx).Where(dao.PlayBanner.Columns().DeletedAt, nil)
	if in.Title != "" {
		m = m.WhereLike(dao.PlayBanner.Columns().Title, "%"+in.Title+"%")
	}
	if in.Remark != "" {
		m = m.WhereLike(dao.PlayBanner.Columns().Remark, "%"+in.Remark+"%")
	}
	if in.StartTime != "" {
		m = m.WhereGTE(dao.PlayBanner.Columns().CreatedAt, in.StartTime)
	}
	if in.EndTime != "" {
		m = m.WhereLTE(dao.PlayBanner.Columns().CreatedAt, in.EndTime)
	}
	err = m.OrderAsc(dao.PlayBanner.Columns().Id).Limit(10000).Scan(&list)
	return
}


