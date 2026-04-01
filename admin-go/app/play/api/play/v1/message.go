package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Message API

// MessageCreateReq 创建会员消息请求
type MessageCreateReq struct {
	g.Meta `path:"/message/create" method:"post" tags:"会员消息" summary:"创建会员消息"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#接收者会员ID不能为空" dc:"接收者会员ID"`
	Title string `json:"title"  dc:"消息标题"`
	Content string `json:"content"  dc:"消息内容"`
	MsgType int `json:"msgType"  dc:"消息类型 1=系统通知 2=订单消息 3=活动消息"`
	BizID string `json:"bizID"  dc:"关联业务ID（订单ID/活动ID等）"`
	IsRead int `json:"isRead"  dc:"是否已读 0=未读 1=已读"`
	Status int `json:"status"  dc:"状态 1=正常 0=禁用"`
}

// MessageCreateRes 创建会员消息响应
type MessageCreateRes struct {
	g.Meta `mime:"application/json"`
}

// MessageUpdateReq 更新会员消息请求
type MessageUpdateReq struct {
	g.Meta `path:"/message/update" method:"put" tags:"会员消息" summary:"更新会员消息"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员消息ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"接收者会员ID"`
	Title string `json:"title" dc:"消息标题"`
	Content string `json:"content" dc:"消息内容"`
	MsgType int `json:"msgType" dc:"消息类型 1=系统通知 2=订单消息 3=活动消息"`
	BizID string `json:"bizID" dc:"关联业务ID（订单ID/活动ID等）"`
	IsRead int `json:"isRead" dc:"是否已读 0=未读 1=已读"`
	Status int `json:"status" dc:"状态 1=正常 0=禁用"`
}

// MessageUpdateRes 更新会员消息响应
type MessageUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// MessageDeleteReq 删除会员消息请求
type MessageDeleteReq struct {
	g.Meta `path:"/message/delete" method:"delete" tags:"会员消息" summary:"删除会员消息"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员消息ID"`
}

// MessageDeleteRes 删除会员消息响应
type MessageDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// MessageBatchDeleteReq 批量删除会员消息请求
type MessageBatchDeleteReq struct {
	g.Meta `path:"/message/batch-delete" method:"delete" tags:"会员消息" summary:"批量删除会员消息"`
	IDs    []snowflake.JsonInt64 `json:"ids" v:"required#ID列表不能为空" dc:"会员消息ID列表"`
}

// MessageBatchDeleteRes 批量删除会员消息响应
type MessageBatchDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// MessageDetailReq 获取会员消息详情请求
type MessageDetailReq struct {
	g.Meta `path:"/message/detail" method:"get" tags:"会员消息" summary:"获取会员消息详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员消息ID"`
}

// MessageDetailRes 获取会员消息详情响应
type MessageDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.MessageDetailOutput
}

// MessageListReq 获取会员消息列表请求
type MessageListReq struct {
	g.Meta    `path:"/message/list" method:"get" tags:"会员消息" summary:"获取会员消息列表"`
	PageNum   int    `json:"pageNum" d:"1" dc:"页码"`
	PageSize  int    `json:"pageSize" d:"10" dc:"每页数量"`
	OrderBy   string `json:"orderBy" dc:"排序字段"`
	OrderDir  string `json:"orderDir" d:"asc" dc:"排序方向:asc/desc"`
	StartTime string `json:"startTime" dc:"开始时间"`
	EndTime   string `json:"endTime" dc:"结束时间"`
	Title string `json:"title" dc:"消息标题"`
}

// MessageListRes 获取会员消息列表响应
type MessageListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.MessageListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}
// MessageExportReq 导出会员消息请求
type MessageExportReq struct {
	g.Meta    `path:"/message/export" method:"get" tags:"会员消息" summary:"导出会员消息"`
	StartTime string `json:"startTime" dc:"开始时间"`
	EndTime   string `json:"endTime" dc:"结束时间"`
	Title string `json:"title" dc:"消息标题"`
}

// MessageExportRes 导出会员消息响应
type MessageExportRes struct {
	g.Meta `mime:"text/csv"`
}


