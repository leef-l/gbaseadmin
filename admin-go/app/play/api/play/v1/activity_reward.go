package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// ActivityReward API

// ActivityRewardCreateReq 创建活动奖励表请求
type ActivityRewardCreateReq struct {
	g.Meta `path:"/activity_reward/create" method:"post" tags:"活动奖励表" summary:"创建活动奖励表"`
	ActivityID snowflake.JsonInt64 `json:"activityID" v:"required#活动ID不能为空" dc:"活动ID"`
	RewardType int `json:"rewardType"  dc:"奖励类型"`
	RewardValue int64 `json:"rewardValue"  dc:"奖励数值（余额=分，优惠券=coupon_id，经验=值，等级天数=天）"`
	RewardName string `json:"rewardName" v:"required#奖励名称（展示用，如"送50元余额"）不能为空" dc:"奖励名称（展示用，如"送50元余额"）"`
	Sort int `json:"sort"  dc:"排序（升序）"`
}

// ActivityRewardCreateRes 创建活动奖励表响应
type ActivityRewardCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityRewardUpdateReq 更新活动奖励表请求
type ActivityRewardUpdateReq struct {
	g.Meta `path:"/activity_reward/update" method:"put" tags:"活动奖励表" summary:"更新活动奖励表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动奖励表ID"`
	ActivityID snowflake.JsonInt64 `json:"activityID" dc:"活动ID"`
	RewardType int `json:"rewardType" dc:"奖励类型"`
	RewardValue int64 `json:"rewardValue" dc:"奖励数值（余额=分，优惠券=coupon_id，经验=值，等级天数=天）"`
	RewardName string `json:"rewardName" dc:"奖励名称（展示用，如"送50元余额"）"`
	Sort int `json:"sort" dc:"排序（升序）"`
}

// ActivityRewardUpdateRes 更新活动奖励表响应
type ActivityRewardUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityRewardDeleteReq 删除活动奖励表请求
type ActivityRewardDeleteReq struct {
	g.Meta `path:"/activity_reward/delete" method:"delete" tags:"活动奖励表" summary:"删除活动奖励表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动奖励表ID"`
}

// ActivityRewardDeleteRes 删除活动奖励表响应
type ActivityRewardDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityRewardDetailReq 获取活动奖励表详情请求
type ActivityRewardDetailReq struct {
	g.Meta `path:"/activity_reward/detail" method:"get" tags:"活动奖励表" summary:"获取活动奖励表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动奖励表ID"`
}

// ActivityRewardDetailRes 获取活动奖励表详情响应
type ActivityRewardDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ActivityRewardDetailOutput
}

// ActivityRewardListReq 获取活动奖励表列表请求
type ActivityRewardListReq struct {
	g.Meta   `path:"/activity_reward/list" method:"get" tags:"活动奖励表" summary:"获取活动奖励表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	RewardType int `json:"rewardType" dc:"奖励类型"`
}

// ActivityRewardListRes 获取活动奖励表列表响应
type ActivityRewardListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ActivityRewardListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

