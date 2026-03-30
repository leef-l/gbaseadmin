package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// ActivityJoin API

// ActivityJoinCreateReq 创建活动参与记录表请求
type ActivityJoinCreateReq struct {
	g.Meta `path:"/activity_join/create" method:"post" tags:"活动参与记录表" summary:"创建活动参与记录表"`
	ActivityID snowflake.JsonInt64 `json:"activityID" v:"required#活动ID不能为空" dc:"活动ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#会员ID不能为空" dc:"会员ID"`
	JoinStatus int `json:"joinStatus"  dc:"参与状态"`
	CurrentStep int `json:"currentStep"  dc:"当前完成到第几步（步骤活动用）"`
	FinishAt *gtime.Time `json:"finishAt"  dc:"完成时间"`
	RewardAt *gtime.Time `json:"rewardAt"  dc:"领奖时间"`
	Remark string `json:"remark"  dc:"备注"`
}

// ActivityJoinCreateRes 创建活动参与记录表响应
type ActivityJoinCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityJoinUpdateReq 更新活动参与记录表请求
type ActivityJoinUpdateReq struct {
	g.Meta `path:"/activity_join/update" method:"put" tags:"活动参与记录表" summary:"更新活动参与记录表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动参与记录表ID"`
	ActivityID snowflake.JsonInt64 `json:"activityID" dc:"活动ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"会员ID"`
	JoinStatus int `json:"joinStatus" dc:"参与状态"`
	CurrentStep int `json:"currentStep" dc:"当前完成到第几步（步骤活动用）"`
	FinishAt *gtime.Time `json:"finishAt" dc:"完成时间"`
	RewardAt *gtime.Time `json:"rewardAt" dc:"领奖时间"`
	Remark string `json:"remark" dc:"备注"`
}

// ActivityJoinUpdateRes 更新活动参与记录表响应
type ActivityJoinUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityJoinDeleteReq 删除活动参与记录表请求
type ActivityJoinDeleteReq struct {
	g.Meta `path:"/activity_join/delete" method:"delete" tags:"活动参与记录表" summary:"删除活动参与记录表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动参与记录表ID"`
}

// ActivityJoinDeleteRes 删除活动参与记录表响应
type ActivityJoinDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityJoinDetailReq 获取活动参与记录表详情请求
type ActivityJoinDetailReq struct {
	g.Meta `path:"/activity_join/detail" method:"get" tags:"活动参与记录表" summary:"获取活动参与记录表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动参与记录表ID"`
}

// ActivityJoinDetailRes 获取活动参与记录表详情响应
type ActivityJoinDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ActivityJoinDetailOutput
}

// ActivityJoinListReq 获取活动参与记录表列表请求
type ActivityJoinListReq struct {
	g.Meta   `path:"/activity_join/list" method:"get" tags:"活动参与记录表" summary:"获取活动参与记录表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	JoinStatus int `json:"joinStatus" dc:"参与状态"`
}

// ActivityJoinListRes 获取活动参与记录表列表响应
type ActivityJoinListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ActivityJoinListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

