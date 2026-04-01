package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 活动列表（公开） ====================

type ActivityListApiReq struct {
	g.Meta   `path:"/activity/list" method:"get" tags:"C端活动" summary:"活动列表"`
	Page     int `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type ActivityListApiRes struct {
	g.Meta `mime:"application/json"`
	Total  int                    `json:"total" dc:"总数"`
	List   []ActivityListApiItem  `json:"list" dc:"活动列表"`
}

type ActivityListApiItem struct {
	ActivityID  string `json:"activityId" dc:"活动ID"`
	Title       string `json:"title" dc:"活动标题"`
	Cover       string `json:"cover" dc:"封面图"`
	Description string `json:"description" dc:"活动简介"`
	Type        int    `json:"type" dc:"类型:1=充值活动,2=下单活动,3=注册活动,4=图文步骤活动,5=自定义活动"`
	StartTime   string `json:"startTime" dc:"开始时间"`
	EndTime     string `json:"endTime" dc:"结束时间"`
	JoinCount   int    `json:"joinCount" dc:"参与人数"`
}

// ==================== 活动详情（公开） ====================

type ActivityDetailApiReq struct {
	g.Meta     `path:"/activity/detail" method:"get" tags:"C端活动" summary:"活动详情"`
	ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
}

type ActivityDetailApiRes struct {
	g.Meta         `mime:"application/json"`
	ActivityID     string                  `json:"activityId"      dc:"活动ID"`
	Title          string                  `json:"title"           dc:"活动标题"`
	Cover          string                  `json:"cover"           dc:"封面图"`
	Content        string                  `json:"content"         dc:"活动详情(富文本)"`
	Type           int                     `json:"type"            dc:"类型"`
	StartTime      string                  `json:"startTime"       dc:"开始时间"`
	EndTime        string                  `json:"endTime"         dc:"结束时间"`
	JoinCount      int                     `json:"joinCount"       dc:"参与人数"`
	Steps          []ActivityStepApiItem   `json:"steps"           dc:"活动步骤列表"`
	Rewards        []ActivityRewardApiItem `json:"rewards"         dc:"奖励列表"`
	Joined         bool                    `json:"joined"          dc:"是否已报名"`
	CompletedSteps []string                `json:"completedSteps"  dc:"已完成的步骤ID列表"`
}

type ActivityStepApiItem struct {
	StepID      string `json:"stepId"      dc:"步骤ID"`
	StepNo      int    `json:"stepNo"      dc:"步骤序号"`
	Title       string `json:"title"       dc:"步骤标题"`
	Description string `json:"description" dc:"步骤描述"`
	StepType    int    `json:"stepType"    dc:"步骤类型:1=文字,2=链接,3=图片"`
	ExampleText string `json:"exampleText" dc:"示例文字或链接URL"`
	StepImage   string `json:"stepImage"   dc:"步骤示例图片"`
	IsRequired  int    `json:"isRequired"  dc:"是否需要填写:0=不需要,1=需要"`
}

type ActivityRewardApiItem struct {
	RewardID    string `json:"rewardId" dc:"奖励ID"`
	RewardName  string `json:"rewardName" dc:"奖励名称"`
	RewardType  int    `json:"rewardType" dc:"奖励类型:1=余额,2=优惠券,3=经验值"`
	RewardValue int64  `json:"rewardValue" dc:"奖励值"`
}

// ==================== 报名参与活动（MemberAuth） ====================

type ActivityJoinApiReq struct {
	g.Meta     `path:"/activity/join" method:"post" tags:"C端活动" summary:"报名参与活动"`
	ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
}

type ActivityJoinApiRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 完成活动步骤（MemberAuth） ====================

type ActivityStepApiReq struct {
	g.Meta     `path:"/activity/complete_step" method:"post" tags:"C端活动" summary:"完成活动步骤"`
	ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
	StepID     string `json:"stepId" v:"required#步骤ID不能为空" dc:"步骤ID"`
	ImageUrl   string `json:"imageUrl" dc:"用户上传的图片URL（图片步骤时传入）"`
}

type ActivityStepApiRes struct {
	g.Meta      `mime:"application/json"`
	CurrentStep int  `json:"currentStep" dc:"当前步骤序号"`
	IsCompleted bool `json:"isCompleted" dc:"是否已完成全部步骤"`
}

// ==================== 取消报名（MemberAuth） ====================

type ActivityQuitApiReq struct {
	g.Meta     `path:"/activity/quit" method:"post" tags:"C端活动" summary:"取消报名"`
	ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
}

type ActivityQuitApiRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 领取奖励（MemberAuth） ====================

type ActivityClaimApiReq struct {
	g.Meta     `path:"/activity/claim_reward" method:"post" tags:"C端活动" summary:"领取奖励"`
	ActivityID string `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
	RewardID   string `json:"rewardId" v:"required#奖励ID不能为空" dc:"奖励ID"`
}

type ActivityClaimApiRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 我参与的活动列表（MemberAuth） ====================

type ActivityMyJoinsReq struct {
	g.Meta   `path:"/activity/my_list" method:"get" tags:"C端活动" summary:"我参与的活动列表"`
	Page     int `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type ActivityMyJoinsRes struct {
	g.Meta `mime:"application/json"`
	Total  int                      `json:"total" dc:"总数"`
	List   []ActivityMyJoinsItem    `json:"list" dc:"我参与的活动列表"`
}

type ActivityMyJoinsItem struct {
	ActivityID  string `json:"activityId" dc:"活动ID"`
	Title       string `json:"title" dc:"活动标题"`
	Cover       string `json:"cover" dc:"封面图"`
	Type        int    `json:"type" dc:"类型"`
	StartTime   string `json:"startTime" dc:"开始时间"`
	EndTime     string `json:"endTime" dc:"结束时间"`
	CurrentStep int    `json:"currentStep" dc:"当前步骤"`
	TotalSteps  int    `json:"totalSteps" dc:"总步骤数"`
	JoinStatus  int    `json:"joinStatus" dc:"参与状态:0=已报名,1=进行中,2=已完成,3=已领奖"`
	JoinedAt    string `json:"joinedAt" dc:"参与时间"`
}
