package consts

// ActivityType 活动类型
const (
	ActivityTypeRecharge = 1 // 充值活动
	ActivityTypeOrder    = 2 // 下单活动
	ActivityTypeRegister = 3 // 注册活动
	ActivityTypeStep     = 4 // 图文步骤活动
	ActivityTypeCustom   = 5 // 自定义活动
)

// ActivityConditionType 参与条件
const (
	ActivityConditionTypeNone       = 0 // 无条件
	ActivityConditionTypeSignUp     = 1 // 需报名
	ActivityConditionTypeRecharge   = 2 // 充值满额
	ActivityConditionTypeOrder      = 3 // 下单满额
	ActivityConditionTypeCompleteStep = 4 // 完成步骤
)

// ActivityIsAutoReward 是否自动发奖
const (
	ActivityIsAutoRewardNo  = 0
	ActivityIsAutoRewardYes = 1
)

// ActivityStatus 状态
const (
	ActivityStatusOff = 0
	ActivityStatusOn  = 1
)
