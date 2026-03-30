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

// ActivityRewardCreateReq 创建æ´»åŠ¨å¥–åŠ±è¡¨请求
type ActivityRewardCreateReq struct {
	g.Meta `path:"/activity_reward/create" method:"post" tags:"æ´»åŠ¨å¥–åŠ±è¡¨" summary:"创建æ´»åŠ¨å¥–åŠ±è¡¨"`
	ActivityID snowflake.JsonInt64 `json:"activityID" v:"required#æ´»åŠ¨ID不能为空" dc:"æ´»åŠ¨ID"`
	RewardType int `json:"rewardType"  dc:"å¥–åŠ±ç±»åž‹"`
	RewardValue int64 `json:"rewardValue"  dc:"å¥–åŠ±æ•°å€¼"`
	RewardName string `json:"rewardName" v:"required#å¥–åŠ±åç§°不能为空" dc:"å¥–åŠ±åç§°"`
	Sort int `json:"sort"  dc:"æŽ’åº"`
}

// ActivityRewardCreateRes 创建æ´»åŠ¨å¥–åŠ±è¡¨响应
type ActivityRewardCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityRewardUpdateReq 更新æ´»åŠ¨å¥–åŠ±è¡¨请求
type ActivityRewardUpdateReq struct {
	g.Meta `path:"/activity_reward/update" method:"put" tags:"æ´»åŠ¨å¥–åŠ±è¡¨" summary:"更新æ´»åŠ¨å¥–åŠ±è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨å¥–åŠ±è¡¨ID"`
	ActivityID snowflake.JsonInt64 `json:"activityID" dc:"æ´»åŠ¨ID"`
	RewardType int `json:"rewardType" dc:"å¥–åŠ±ç±»åž‹"`
	RewardValue int64 `json:"rewardValue" dc:"å¥–åŠ±æ•°å€¼"`
	RewardName string `json:"rewardName" dc:"å¥–åŠ±åç§°"`
	Sort int `json:"sort" dc:"æŽ’åº"`
}

// ActivityRewardUpdateRes 更新æ´»åŠ¨å¥–åŠ±è¡¨响应
type ActivityRewardUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityRewardDeleteReq 删除æ´»åŠ¨å¥–åŠ±è¡¨请求
type ActivityRewardDeleteReq struct {
	g.Meta `path:"/activity_reward/delete" method:"delete" tags:"æ´»åŠ¨å¥–åŠ±è¡¨" summary:"删除æ´»åŠ¨å¥–åŠ±è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨å¥–åŠ±è¡¨ID"`
}

// ActivityRewardDeleteRes 删除æ´»åŠ¨å¥–åŠ±è¡¨响应
type ActivityRewardDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityRewardDetailReq 获取æ´»åŠ¨å¥–åŠ±è¡¨详情请求
type ActivityRewardDetailReq struct {
	g.Meta `path:"/activity_reward/detail" method:"get" tags:"æ´»åŠ¨å¥–åŠ±è¡¨" summary:"获取æ´»åŠ¨å¥–åŠ±è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨å¥–åŠ±è¡¨ID"`
}

// ActivityRewardDetailRes 获取æ´»åŠ¨å¥–åŠ±è¡¨详情响应
type ActivityRewardDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ActivityRewardDetailOutput
}

// ActivityRewardListReq 获取æ´»åŠ¨å¥–åŠ±è¡¨列表请求
type ActivityRewardListReq struct {
	g.Meta   `path:"/activity_reward/list" method:"get" tags:"æ´»åŠ¨å¥–åŠ±è¡¨" summary:"获取æ´»åŠ¨å¥–åŠ±è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	RewardType int `json:"rewardType" dc:"å¥–åŠ±ç±»åž‹"`
}

// ActivityRewardListRes 获取æ´»åŠ¨å¥–åŠ±è¡¨列表响应
type ActivityRewardListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ActivityRewardListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

