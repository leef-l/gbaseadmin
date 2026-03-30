package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// ActivityJoin API

// ActivityJoinCreateReq 创建æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨请求
type ActivityJoinCreateReq struct {
	g.Meta `path:"/activity_join/create" method:"post" tags:"æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨" summary:"创建æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨"`
	ActivityID snowflake.JsonInt64 `json:"activityID" v:"required#æ´»åŠ¨ID不能为空" dc:"æ´»åŠ¨ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#ä¼šå‘˜ID不能为空" dc:"ä¼šå‘˜ID"`
	JoinStatus int `json:"joinStatus"  dc:"å‚ä¸ŽçŠ¶æ€"`
	CurrentStep int `json:"currentStep"  dc:"å½“å‰å®Œæˆåˆ°ç¬¬å‡ æ­¥"`
	FinishAt *gtime.Time `json:"finishAt"  dc:"å®Œæˆæ—¶é—´"`
	RewardAt *gtime.Time `json:"rewardAt"  dc:"é¢†å¥–æ—¶é—´"`
	Remark string `json:"remark"  dc:"å¤‡æ³¨"`
}

// ActivityJoinCreateRes 创建æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨响应
type ActivityJoinCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityJoinUpdateReq 更新æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨请求
type ActivityJoinUpdateReq struct {
	g.Meta `path:"/activity_join/update" method:"put" tags:"æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨" summary:"更新æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨ID"`
	ActivityID snowflake.JsonInt64 `json:"activityID" dc:"æ´»åŠ¨ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"ä¼šå‘˜ID"`
	JoinStatus int `json:"joinStatus" dc:"å‚ä¸ŽçŠ¶æ€"`
	CurrentStep int `json:"currentStep" dc:"å½“å‰å®Œæˆåˆ°ç¬¬å‡ æ­¥"`
	FinishAt *gtime.Time `json:"finishAt" dc:"å®Œæˆæ—¶é—´"`
	RewardAt *gtime.Time `json:"rewardAt" dc:"é¢†å¥–æ—¶é—´"`
	Remark string `json:"remark" dc:"å¤‡æ³¨"`
}

// ActivityJoinUpdateRes 更新æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨响应
type ActivityJoinUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityJoinDeleteReq 删除æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨请求
type ActivityJoinDeleteReq struct {
	g.Meta `path:"/activity_join/delete" method:"delete" tags:"æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨" summary:"删除æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨ID"`
}

// ActivityJoinDeleteRes 删除æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨响应
type ActivityJoinDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityJoinDetailReq 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨详情请求
type ActivityJoinDetailReq struct {
	g.Meta `path:"/activity_join/detail" method:"get" tags:"æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨" summary:"获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨ID"`
}

// ActivityJoinDetailRes 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨详情响应
type ActivityJoinDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ActivityJoinDetailOutput
}

// ActivityJoinListReq 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨列表请求
type ActivityJoinListReq struct {
	g.Meta   `path:"/activity_join/list" method:"get" tags:"æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨" summary:"获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	JoinStatus int `json:"joinStatus" dc:"å‚ä¸ŽçŠ¶æ€"`
}

// ActivityJoinListRes 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨列表响应
type ActivityJoinListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ActivityJoinListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

