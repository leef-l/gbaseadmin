package consts

// ActivityStepLogStepType 步骤类型
const (
	ActivityStepLogStepTypeText  = 1 // 文字
	ActivityStepLogStepTypeLink  = 2 // 链接
	ActivityStepLogStepTypeImage = 3 // 图片
)

// ActivityStepLogAuditStatus 审核状态
const (
	ActivityStepLogAuditStatusPending  = 0 // 待审核
	ActivityStepLogAuditStatusApproved = 1 // 通过
	ActivityStepLogAuditStatusRejected = 2 // 驳回
)
