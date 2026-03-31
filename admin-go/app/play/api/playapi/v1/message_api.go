package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 我的消息列表（MemberAuth） ====================

type MessageListReq struct {
	g.Meta   `path:"/member/messages" method:"get" tags:"C端消息" summary:"我的消息列表"`
	Type     string `json:"type" dc:"消息类型筛选:order/system/activity，为空查全部"`
	Page     int    `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type MessageListRes struct {
	g.Meta `mime:"application/json"`
	Total  int           `json:"total" dc:"总数"`
	List   []MessageItem `json:"list" dc:"消息列表"`
}

type MessageItem struct {
	Id     string `json:"id" dc:"消息ID"`
	Type   string `json:"type" dc:"消息类型:order/system/activity"`
	Title  string `json:"title" dc:"消息标题"`
	Desc   string `json:"desc" dc:"消息描述"`
	Time   string `json:"time" dc:"消息时间"`
	Unread bool   `json:"unread" dc:"是否未读"`
}

// ==================== 标记单条消息已读（MemberAuth） ====================

type MessageReadReq struct {
	g.Meta `path:"/member/message/read" method:"post" tags:"C端消息" summary:"标记消息已读"`
	Id     string `json:"id" v:"required#消息ID不能为空" dc:"消息ID"`
}

type MessageReadRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 标记全部消息已读（MemberAuth） ====================

type MessageReadAllReq struct {
	g.Meta `path:"/member/message/read_all" method:"post" tags:"C端消息" summary:"全部标记已读"`
}

type MessageReadAllRes struct {
	g.Meta `mime:"application/json"`
}
