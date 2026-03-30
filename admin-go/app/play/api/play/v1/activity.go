package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// Activity API

// ActivityCreateReq 创建æ´»åŠ¨è¡¨请求
type ActivityCreateReq struct {
	g.Meta `path:"/activity/create" method:"post" tags:"æ´»åŠ¨è¡¨" summary:"创建æ´»åŠ¨è¡¨"`
	Title string `json:"title" v:"required#æ´»åŠ¨åç§°不能为空" dc:"æ´»åŠ¨åç§°"`
	CoverImage string `json:"coverImage"  dc:"æ´»åŠ¨å°é¢å›¾"`
	DescContent string `json:"descContent"  dc:"æ´»åŠ¨è¯¦æƒ…æè¿°"`
	Type int `json:"type"  dc:"æ´»åŠ¨ç±»åž‹"`
	ConditionType int `json:"conditionType"  dc:"å‚ä¸Žæ¡ä»¶"`
	ConditionValue int64 `json:"conditionValue"  dc:"æ¡ä»¶å€¼"`
	IsAutoReward int `json:"isAutoReward"  dc:"æ˜¯å¦è‡ªåŠ¨å‘å¥–"`
	StartAt *gtime.Time `json:"startAt" v:"required#æ´»åŠ¨å¼€å§‹æ—¶é—´不能为空" dc:"æ´»åŠ¨å¼€å§‹æ—¶é—´"`
	EndAt *gtime.Time `json:"endAt" v:"required#æ´»åŠ¨ç»“æŸæ—¶é—´不能为空" dc:"æ´»åŠ¨ç»“æŸæ—¶é—´"`
	MaxNum int `json:"maxNum"  dc:"å‚ä¸Žäººæ•°ä¸Šé™ï¼ˆ0ä¸é™ï¼‰"`
	JoinNum int `json:"joinNum"  dc:"å·²å‚ä¸Žäººæ•°"`
	Sort int `json:"sort"  dc:"æŽ’åº"`
	Status int `json:"status"  dc:"çŠ¶æ€"`
}

// ActivityCreateRes 创建æ´»åŠ¨è¡¨响应
type ActivityCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityUpdateReq 更新æ´»åŠ¨è¡¨请求
type ActivityUpdateReq struct {
	g.Meta `path:"/activity/update" method:"put" tags:"æ´»åŠ¨è¡¨" summary:"更新æ´»åŠ¨è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨è¡¨ID"`
	Title string `json:"title" dc:"æ´»åŠ¨åç§°"`
	CoverImage string `json:"coverImage" dc:"æ´»åŠ¨å°é¢å›¾"`
	DescContent string `json:"descContent" dc:"æ´»åŠ¨è¯¦æƒ…æè¿°"`
	Type int `json:"type" dc:"æ´»åŠ¨ç±»åž‹"`
	ConditionType int `json:"conditionType" dc:"å‚ä¸Žæ¡ä»¶"`
	ConditionValue int64 `json:"conditionValue" dc:"æ¡ä»¶å€¼"`
	IsAutoReward int `json:"isAutoReward" dc:"æ˜¯å¦è‡ªåŠ¨å‘å¥–"`
	StartAt *gtime.Time `json:"startAt" dc:"æ´»åŠ¨å¼€å§‹æ—¶é—´"`
	EndAt *gtime.Time `json:"endAt" dc:"æ´»åŠ¨ç»“æŸæ—¶é—´"`
	MaxNum int `json:"maxNum" dc:"å‚ä¸Žäººæ•°ä¸Šé™ï¼ˆ0ä¸é™ï¼‰"`
	JoinNum int `json:"joinNum" dc:"å·²å‚ä¸Žäººæ•°"`
	Sort int `json:"sort" dc:"æŽ’åº"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// ActivityUpdateRes 更新æ´»åŠ¨è¡¨响应
type ActivityUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityDeleteReq 删除æ´»åŠ¨è¡¨请求
type ActivityDeleteReq struct {
	g.Meta `path:"/activity/delete" method:"delete" tags:"æ´»åŠ¨è¡¨" summary:"删除æ´»åŠ¨è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨è¡¨ID"`
}

// ActivityDeleteRes 删除æ´»åŠ¨è¡¨响应
type ActivityDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityDetailReq 获取æ´»åŠ¨è¡¨详情请求
type ActivityDetailReq struct {
	g.Meta `path:"/activity/detail" method:"get" tags:"æ´»åŠ¨è¡¨" summary:"获取æ´»åŠ¨è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨è¡¨ID"`
}

// ActivityDetailRes 获取æ´»åŠ¨è¡¨详情响应
type ActivityDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ActivityDetailOutput
}

// ActivityListReq 获取æ´»åŠ¨è¡¨列表请求
type ActivityListReq struct {
	g.Meta   `path:"/activity/list" method:"get" tags:"æ´»åŠ¨è¡¨" summary:"获取æ´»åŠ¨è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Type int `json:"type" dc:"æ´»åŠ¨ç±»åž‹"`
	ConditionType int `json:"conditionType" dc:"å‚ä¸Žæ¡ä»¶"`
	IsAutoReward int `json:"isAutoReward" dc:"æ˜¯å¦è‡ªåŠ¨å‘å¥–"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// ActivityListRes 获取æ´»åŠ¨è¡¨列表响应
type ActivityListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ActivityListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

