package balance_log

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// AddLog 统一余额变动入口
func (s *sBalanceLog) AddLog(ctx context.Context, in *model.AddBalanceLogInput) error {
	// 查询当前余额
	val, err := dao.PlayMember.Ctx(ctx).Where(dao.PlayMember.Columns().Id, in.MemberID).Value(dao.PlayMember.Columns().Balance)
	if err != nil {
		return err
	}
	beforeBalance := val.Int64()
	afterBalance := beforeBalance + in.ChangeAmount

	if afterBalance < 0 {
		return gerror.New("余额不足")
	}

	// 写入余额流水
	logID := snowflake.Generate()
	_, err = dao.PlayBalanceLog.Ctx(ctx).Data(g.Map{
		dao.PlayBalanceLog.Columns().Id:            logID,
		dao.PlayBalanceLog.Columns().MemberId:      in.MemberID,
		dao.PlayBalanceLog.Columns().BizType:       in.BizType,
		dao.PlayBalanceLog.Columns().BizId:         in.BizID,
		dao.PlayBalanceLog.Columns().ChangeAmount:  in.ChangeAmount,
		dao.PlayBalanceLog.Columns().BeforeBalance: beforeBalance,
		dao.PlayBalanceLog.Columns().AfterBalance:  afterBalance,
		dao.PlayBalanceLog.Columns().Remark:        in.Remark,
		dao.PlayBalanceLog.Columns().CreatedAt:     gtime.Now(),
		dao.PlayBalanceLog.Columns().UpdatedAt:     gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}

	// 更新会员余额
	_, err = dao.PlayMember.Ctx(ctx).Where(dao.PlayMember.Columns().Id, in.MemberID).Data(g.Map{
		dao.PlayMember.Columns().Balance:   afterBalance,
		dao.PlayMember.Columns().UpdatedAt: gtime.Now(),
	}).Update()
	return err
}
