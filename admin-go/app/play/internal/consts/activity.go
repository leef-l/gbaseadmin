package consts

// ActivityType 活动类型
const (
	ActivityType1 = 1 // 充值活动
	ActivityType2 = 2 // 下单活动
	ActivityType3 = 3 // 注册活动
	ActivityType4 = 4 // 图文步骤活动
	ActivityType5 = 5 // 自定义活动
)

// ActivityConditionType 参与条件
const (
	ActivityConditionType0 = 0 // 无条件
	ActivityConditionType1 = 1 // 需报名
	ActivityConditionType2 = 2 // 充值满额
	ActivityConditionType3 = 3 // 下单满额
	ActivityConditionType4 = 4 // 完成步骤
)

// ActivityIsAutoReward 是否自动发奖
const (
	ActivityIsAutoReward0 = 0 // 否（需审核）
	ActivityIsAutoReward1 = 1 // 是（用户完成即发）
)

// ActivityStatus 状态
const (
	ActivityStatus0 = 0 // 关闭
	ActivityStatus1 = 1 // 开启
)

