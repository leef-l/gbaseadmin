package activity_join

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
	service.RegisterActivityJoin(New())
}

func New() *sActivityJoin {
	return &sActivityJoin{}
}

type sActivityJoin struct{}

// Create 创建æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨
func (s *sActivityJoin) Create(ctx context.Context, in *model.ActivityJoinCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayActivityJoin.Ctx(ctx).Data(g.Map{
		dao.PlayActivityJoin.Columns().Id:        id,
		dao.PlayActivityJoin.Columns().ActivityId: in.ActivityID,
		dao.PlayActivityJoin.Columns().MemberId: in.MemberID,
		dao.PlayActivityJoin.Columns().JoinStatus: in.JoinStatus,
		dao.PlayActivityJoin.Columns().CurrentStep: in.CurrentStep,
		dao.PlayActivityJoin.Columns().FinishAt: in.FinishAt,
		dao.PlayActivityJoin.Columns().RewardAt: in.RewardAt,
		dao.PlayActivityJoin.Columns().Remark: in.Remark,
		dao.PlayActivityJoin.Columns().CreatedAt: gtime.Now(),
		dao.PlayActivityJoin.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨
func (s *sActivityJoin) Update(ctx context.Context, in *model.ActivityJoinUpdateInput) error {
	data := g.Map{
		dao.PlayActivityJoin.Columns().ActivityId: in.ActivityID,
		dao.PlayActivityJoin.Columns().MemberId: in.MemberID,
		dao.PlayActivityJoin.Columns().JoinStatus: in.JoinStatus,
		dao.PlayActivityJoin.Columns().CurrentStep: in.CurrentStep,
		dao.PlayActivityJoin.Columns().FinishAt: in.FinishAt,
		dao.PlayActivityJoin.Columns().RewardAt: in.RewardAt,
		dao.PlayActivityJoin.Columns().Remark: in.Remark,
		dao.PlayActivityJoin.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayActivityJoin.Ctx(ctx).Where(dao.PlayActivityJoin.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨
func (s *sActivityJoin) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayActivityJoin.Ctx(ctx).Where(dao.PlayActivityJoin.Columns().Id, id).Data(g.Map{
		dao.PlayActivityJoin.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨详情
func (s *sActivityJoin) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityJoinDetailOutput, err error) {
	out = &model.ActivityJoinDetailOutput{}
	err = dao.PlayActivityJoin.Ctx(ctx).Where(dao.PlayActivityJoin.Columns().Id, id).Where(dao.PlayActivityJoin.Columns().DeletedAt, nil).Scan(out)
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

// List 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨列表
func (s *sActivityJoin) List(ctx context.Context, in *model.ActivityJoinListInput) (list []*model.ActivityJoinListOutput, total int, err error) {
	m := dao.PlayActivityJoin.Ctx(ctx).Where(dao.PlayActivityJoin.Columns().DeletedAt, nil)
	if in.JoinStatus > 0 {
		m = m.Where(dao.PlayActivityJoin.Columns().JoinStatus, in.JoinStatus)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayActivityJoin.Columns().Id).Scan(&list)
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

