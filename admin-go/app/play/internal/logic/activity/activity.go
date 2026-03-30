package activity

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
	service.RegisterActivity(New())
}

func New() *sActivity {
	return &sActivity{}
}

type sActivity struct{}

// Create 创建活动表
func (s *sActivity) Create(ctx context.Context, in *model.ActivityCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayActivity.Ctx(ctx).Data(g.Map{
		dao.PlayActivity.Columns().Id:        id,
		dao.PlayActivity.Columns().Title: in.Title,
		dao.PlayActivity.Columns().CoverImage: in.CoverImage,
		dao.PlayActivity.Columns().DescContent: in.DescContent,
		dao.PlayActivity.Columns().Type: in.Type,
		dao.PlayActivity.Columns().ConditionType: in.ConditionType,
		dao.PlayActivity.Columns().ConditionValue: in.ConditionValue,
		dao.PlayActivity.Columns().IsAutoReward: in.IsAutoReward,
		dao.PlayActivity.Columns().StartAt: in.StartAt,
		dao.PlayActivity.Columns().EndAt: in.EndAt,
		dao.PlayActivity.Columns().MaxNum: in.MaxNum,
		dao.PlayActivity.Columns().JoinNum: in.JoinNum,
		dao.PlayActivity.Columns().Sort: in.Sort,
		dao.PlayActivity.Columns().Status: in.Status,
		dao.PlayActivity.Columns().CreatedAt: gtime.Now(),
		dao.PlayActivity.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新活动表
func (s *sActivity) Update(ctx context.Context, in *model.ActivityUpdateInput) error {
	data := g.Map{
		dao.PlayActivity.Columns().Title: in.Title,
		dao.PlayActivity.Columns().CoverImage: in.CoverImage,
		dao.PlayActivity.Columns().DescContent: in.DescContent,
		dao.PlayActivity.Columns().Type: in.Type,
		dao.PlayActivity.Columns().ConditionType: in.ConditionType,
		dao.PlayActivity.Columns().ConditionValue: in.ConditionValue,
		dao.PlayActivity.Columns().IsAutoReward: in.IsAutoReward,
		dao.PlayActivity.Columns().StartAt: in.StartAt,
		dao.PlayActivity.Columns().EndAt: in.EndAt,
		dao.PlayActivity.Columns().MaxNum: in.MaxNum,
		dao.PlayActivity.Columns().JoinNum: in.JoinNum,
		dao.PlayActivity.Columns().Sort: in.Sort,
		dao.PlayActivity.Columns().Status: in.Status,
		dao.PlayActivity.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayActivity.Ctx(ctx).Where(dao.PlayActivity.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除活动表
func (s *sActivity) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayActivity.Ctx(ctx).Where(dao.PlayActivity.Columns().Id, id).Data(g.Map{
		dao.PlayActivity.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取活动表详情
func (s *sActivity) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityDetailOutput, err error) {
	out = &model.ActivityDetailOutput{}
	err = dao.PlayActivity.Ctx(ctx).Where(dao.PlayActivity.Columns().Id, id).Where(dao.PlayActivity.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取活动表列表
func (s *sActivity) List(ctx context.Context, in *model.ActivityListInput) (list []*model.ActivityListOutput, total int, err error) {
	m := dao.PlayActivity.Ctx(ctx).Where(dao.PlayActivity.Columns().DeletedAt, nil)
	if in.Type > 0 {
		m = m.Where(dao.PlayActivity.Columns().Type, in.Type)
	}
	if in.ConditionType > 0 {
		m = m.Where(dao.PlayActivity.Columns().ConditionType, in.ConditionType)
	}
	if in.IsAutoReward > 0 {
		m = m.Where(dao.PlayActivity.Columns().IsAutoReward, in.IsAutoReward)
	}
	if in.Status > 0 {
		m = m.Where(dao.PlayActivity.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayActivity.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

