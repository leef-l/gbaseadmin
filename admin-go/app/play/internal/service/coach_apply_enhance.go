package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
)

type ICoachApplyEnhance interface {
	Audit(ctx context.Context, in *model.CoachApplyAuditInput) error
}

func CoachApplyEnhance() ICoachApplyEnhance {
	return localCoachApply.(ICoachApplyEnhance)
}
