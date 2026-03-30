package coach_apply

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
)

// Audit 审核陪玩师申请
func (s *sCoachApply) Audit(ctx context.Context, in *model.CoachApplyAuditInput) error {
	_, err := dao.PlayCoachApply.Ctx(ctx).Where(dao.PlayCoachApply.Columns().Id, in.ID).Data(g.Map{
		dao.PlayCoachApply.Columns().AuditStatus: in.AuditStatus,
		dao.PlayCoachApply.Columns().AuditRemark: in.AuditRemark,
		dao.PlayCoachApply.Columns().AuditAt:     gtime.Now(),
		dao.PlayCoachApply.Columns().UpdatedAt:   gtime.Now(),
	}).Update()
	return err
}
