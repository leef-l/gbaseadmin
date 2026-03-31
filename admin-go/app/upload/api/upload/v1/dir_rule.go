package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/upload/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// DirRule API

// DirRuleCreateReq 创建æ–‡ä»¶ç›®å½•è§„åˆ™请求
type DirRuleCreateReq struct {
	g.Meta `path:"/dir_rule/create" method:"post" tags:"æ–‡ä»¶ç›®å½•è§„åˆ™" summary:"创建æ–‡ä»¶ç›®å½•è§„åˆ™"`
	DirID snowflake.JsonInt64 `json:"dirID" v:"required#ç›®å½•ID不能为空" dc:"ç›®å½•ID"`
	Category int `json:"category"  dc:"ç±»åˆ«"`
	SavePath string `json:"savePath"  dc:"ä¿å­˜ç›®å½•"`
	Status int `json:"status"  dc:"çŠ¶æ€"`
}

// DirRuleCreateRes 创建æ–‡ä»¶ç›®å½•è§„åˆ™响应
type DirRuleCreateRes struct {
	g.Meta `mime:"application/json"`
}

// DirRuleUpdateReq 更新æ–‡ä»¶ç›®å½•è§„åˆ™请求
type DirRuleUpdateReq struct {
	g.Meta `path:"/dir_rule/update" method:"put" tags:"æ–‡ä»¶ç›®å½•è§„åˆ™" summary:"更新æ–‡ä»¶ç›®å½•è§„åˆ™"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ–‡ä»¶ç›®å½•è§„åˆ™ID"`
	DirID snowflake.JsonInt64 `json:"dirID" dc:"ç›®å½•ID"`
	Category int `json:"category" dc:"ç±»åˆ«"`
	SavePath string `json:"savePath" dc:"ä¿å­˜ç›®å½•"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// DirRuleUpdateRes 更新æ–‡ä»¶ç›®å½•è§„åˆ™响应
type DirRuleUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// DirRuleDeleteReq 删除æ–‡ä»¶ç›®å½•è§„åˆ™请求
type DirRuleDeleteReq struct {
	g.Meta `path:"/dir_rule/delete" method:"delete" tags:"æ–‡ä»¶ç›®å½•è§„åˆ™" summary:"删除æ–‡ä»¶ç›®å½•è§„åˆ™"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ–‡ä»¶ç›®å½•è§„åˆ™ID"`
}

// DirRuleDeleteRes 删除æ–‡ä»¶ç›®å½•è§„åˆ™响应
type DirRuleDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// DirRuleDetailReq 获取æ–‡ä»¶ç›®å½•è§„åˆ™详情请求
type DirRuleDetailReq struct {
	g.Meta `path:"/dir_rule/detail" method:"get" tags:"æ–‡ä»¶ç›®å½•è§„åˆ™" summary:"获取æ–‡ä»¶ç›®å½•è§„åˆ™详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ–‡ä»¶ç›®å½•è§„åˆ™ID"`
}

// DirRuleDetailRes 获取æ–‡ä»¶ç›®å½•è§„åˆ™详情响应
type DirRuleDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.DirRuleDetailOutput
}

// DirRuleListReq 获取æ–‡ä»¶ç›®å½•è§„åˆ™列表请求
type DirRuleListReq struct {
	g.Meta   `path:"/dir_rule/list" method:"get" tags:"æ–‡ä»¶ç›®å½•è§„åˆ™" summary:"获取æ–‡ä»¶ç›®å½•è§„åˆ™列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Category int `json:"category" dc:"ç±»åˆ«"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// DirRuleListRes 获取æ–‡ä»¶ç›®å½•è§„åˆ™列表响应
type DirRuleListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.DirRuleListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

