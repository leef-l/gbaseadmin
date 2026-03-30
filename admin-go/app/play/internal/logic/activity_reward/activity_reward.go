package activity_reward

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
	service.RegisterActivityReward(New())
}

func New() *sActivityReward {
	return &sActivityReward{}
}

type sActivityReward struct{}

// Create 创建æ´»åŠ¨å¥–åŠ±è¡¨
func (s *sActivityReward) Create(ctx context.Context, in *model.ActivityRewardCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayActivityReward.Ctx(ctx).Data(g.Map{
		dao.PlayActivityReward.Columns().Id:        id,
		dao.PlayActivityReward.Columns().ActivityId: in.ActivityID,
		dao.PlayActivityReward.Columns().RewardType: in.RewardType,
		dao.PlayActivityReward.Columns().RewardValue: in.RewardValue,
		dao.PlayActivityReward.Columns().RewardName: in.RewardName,
		dao.PlayActivityReward.Columns().Sort: in.Sort,
		dao.PlayActivityReward.Columns().CreatedAt: gtime.Now(),
		dao.PlayActivityReward.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新æ´»åŠ¨å¥–åŠ±è¡¨
func (s *sActivityReward) Update(ctx context.Context, in *model.ActivityRewardUpdateInput) error {
	data := g.Map{
		dao.PlayActivityReward.Columns().ActivityId: in.ActivityID,
		dao.PlayActivityReward.Columns().RewardType: in.RewardType,
		dao.PlayActivityReward.Columns().RewardValue: in.RewardValue,
		dao.PlayActivityReward.Columns().RewardName: in.RewardName,
		dao.PlayActivityReward.Columns().Sort: in.Sort,
		dao.PlayActivityReward.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayActivityReward.Ctx(ctx).Where(dao.PlayActivityReward.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除æ´»åŠ¨å¥–åŠ±è¡¨
func (s *sActivityReward) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayActivityReward.Ctx(ctx).Where(dao.PlayActivityReward.Columns().Id, id).Data(g.Map{
		dao.PlayActivityReward.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取æ´»åŠ¨å¥–åŠ±è¡¨详情
func (s *sActivityReward) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityRewardDetailOutput, err error) {
	out = &model.ActivityRewardDetailOutput{}
	err = dao.PlayActivityReward.Ctx(ctx).Where(dao.PlayActivityReward.Columns().Id, id).Where(dao.PlayActivityReward.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询æ´»åŠ¨ID关联显示
	if out.ActivityID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("play_activity").Where("id", out.ActivityID).Where("deleted_at", nil).Value("title")
		out.ActivityTitle = val.String()
	}
	return
}

// List 获取æ´»åŠ¨å¥–åŠ±è¡¨列表
func (s *sActivityReward) List(ctx context.Context, in *model.ActivityRewardListInput) (list []*model.ActivityRewardListOutput, total int, err error) {
	m := dao.PlayActivityReward.Ctx(ctx).Where(dao.PlayActivityReward.Columns().DeletedAt, nil)
	if in.RewardType > 0 {
		m = m.Where(dao.PlayActivityReward.Columns().RewardType, in.RewardType)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayActivityReward.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.ActivityID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("play_activity").Where("id", item.ActivityID).Where("deleted_at", nil).Value("title")
			item.ActivityTitle = val.String()
		}
	}
	return
}

