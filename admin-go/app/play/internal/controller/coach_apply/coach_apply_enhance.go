package coach_apply

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

// Audit 审核陪玩师申请
func (c *cCoachApply) Audit(ctx context.Context, req *v1.CoachApplyAuditReq) (res *v1.CoachApplyAuditRes, err error) {
	err = service.CoachApplyEnhance().Audit(ctx, &model.CoachApplyAuditInput{
		ID:          req.ID,
		AuditStatus: req.AuditStatus,
		AuditRemark: req.AuditRemark,
	})
	return
}
