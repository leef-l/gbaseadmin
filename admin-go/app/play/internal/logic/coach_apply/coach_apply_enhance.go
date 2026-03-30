package coach_apply

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// Audit 审核陪玩师申请
func (s *sCoachApply) Audit(ctx context.Context, in *model.CoachApplyAuditInput) error {
	// 查询申请记录
	apply, err := dao.PlayCoachApply.Ctx(ctx).Where(dao.PlayCoachApply.Columns().Id, in.ID).One()
	if err != nil {
		return err
	}
	if apply.IsEmpty() {
		return gerror.New("申请记录不存在")
	}
	if apply[dao.PlayCoachApply.Columns().AuditStatus].Int() != 0 {
		return gerror.New("该申请已审核，不能重复操作")
	}

	// 更新审核状态
	_, err = dao.PlayCoachApply.Ctx(ctx).Where(dao.PlayCoachApply.Columns().Id, in.ID).Data(g.Map{
		dao.PlayCoachApply.Columns().AuditStatus: in.AuditStatus,
		dao.PlayCoachApply.Columns().AuditRemark: in.AuditRemark,
		dao.PlayCoachApply.Columns().AuditAt:     gtime.Now(),
		dao.PlayCoachApply.Columns().UpdatedAt:   gtime.Now(),
	}).Update()
	if err != nil {
		return err
	}

	// 审核通过，自动创建陪玩师记录
	if in.AuditStatus == 1 {
		coachID := snowflake.Generate()
		memberID := apply[dao.PlayCoachApply.Columns().MemberId].Int64()
		realName := apply[dao.PlayCoachApply.Columns().RealName].String()
		_, err = dao.PlayCoach.Ctx(ctx).Data(g.Map{
			dao.PlayCoach.Columns().Id:        coachID,
			dao.PlayCoach.Columns().MemberId:  memberID,
			dao.PlayCoach.Columns().RealName:  realName,
			dao.PlayCoach.Columns().Intro:     apply[dao.PlayCoachApply.Columns().SkillDesc].String(),
			dao.PlayCoach.Columns().Status:    1,
			dao.PlayCoach.Columns().IsOnline:  0,
			dao.PlayCoach.Columns().CreatedAt: gtime.Now(),
			dao.PlayCoach.Columns().UpdatedAt: gtime.Now(),
		}).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}
