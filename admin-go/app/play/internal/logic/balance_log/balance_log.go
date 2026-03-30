package balance_log

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
	service.RegisterBalanceLog(New())
}

func New() *sBalanceLog {
	return &sBalanceLog{}
}

type sBalanceLog struct{}

// Create 创建余额流水表
func (s *sBalanceLog) Create(ctx context.Context, in *model.BalanceLogCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayBalanceLog.Ctx(ctx).Data(g.Map{
		dao.PlayBalanceLog.Columns().Id:        id,
		dao.PlayBalanceLog.Columns().MemberId: in.MemberID,
		dao.PlayBalanceLog.Columns().BizType: in.BizType,
		dao.PlayBalanceLog.Columns().BizId: in.BizID,
		dao.PlayBalanceLog.Columns().ChangeAmount: in.ChangeAmount,
		dao.PlayBalanceLog.Columns().BeforeBalance: in.BeforeBalance,
		dao.PlayBalanceLog.Columns().AfterBalance: in.AfterBalance,
		dao.PlayBalanceLog.Columns().Remark: in.Remark,
		dao.PlayBalanceLog.Columns().CreatedAt: gtime.Now(),
		dao.PlayBalanceLog.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新余额流水表
func (s *sBalanceLog) Update(ctx context.Context, in *model.BalanceLogUpdateInput) error {
	data := g.Map{
		dao.PlayBalanceLog.Columns().MemberId: in.MemberID,
		dao.PlayBalanceLog.Columns().BizType: in.BizType,
		dao.PlayBalanceLog.Columns().BizId: in.BizID,
		dao.PlayBalanceLog.Columns().ChangeAmount: in.ChangeAmount,
		dao.PlayBalanceLog.Columns().BeforeBalance: in.BeforeBalance,
		dao.PlayBalanceLog.Columns().AfterBalance: in.AfterBalance,
		dao.PlayBalanceLog.Columns().Remark: in.Remark,
		dao.PlayBalanceLog.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayBalanceLog.Ctx(ctx).Where(dao.PlayBalanceLog.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除余额流水表
func (s *sBalanceLog) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayBalanceLog.Ctx(ctx).Where(dao.PlayBalanceLog.Columns().Id, id).Data(g.Map{
		dao.PlayBalanceLog.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取余额流水表详情
func (s *sBalanceLog) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.BalanceLogDetailOutput, err error) {
	out = &model.BalanceLogDetailOutput{}
	err = dao.PlayBalanceLog.Ctx(ctx).Where(dao.PlayBalanceLog.Columns().Id, id).Where(dao.PlayBalanceLog.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取余额流水表列表
func (s *sBalanceLog) List(ctx context.Context, in *model.BalanceLogListInput) (list []*model.BalanceLogListOutput, total int, err error) {
	m := dao.PlayBalanceLog.Ctx(ctx).Where(dao.PlayBalanceLog.Columns().DeletedAt, nil)
	if in.BizType > 0 {
		m = m.Where(dao.PlayBalanceLog.Columns().BizType, in.BizType)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayBalanceLog.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

