package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/upload/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// File API

// FileCreateReq 创建æ–‡ä»¶è®°å½•请求
type FileCreateReq struct {
	g.Meta `path:"/file/create" method:"post" tags:"æ–‡ä»¶è®°å½•" summary:"创建æ–‡ä»¶è®°å½•"`
	DirID snowflake.JsonInt64 `json:"dirID"  dc:"æ‰€å±žç›®å½•"`
	Name string `json:"name" v:"required#æ–‡ä»¶åç§°不能为空" dc:"æ–‡ä»¶åç§°"`
	URL string `json:"url" v:"required#æ–‡ä»¶åœ°å€不能为空" dc:"æ–‡ä»¶åœ°å€"`
	Ext string `json:"ext"  dc:"æ–‡ä»¶æ‰©å±•å"`
	Size int64 `json:"size"  dc:"æ–‡ä»¶å¤§å°"`
	Mime string `json:"mime"  dc:"MIMEç±»åž‹"`
	Storage int `json:"storage"  dc:"å­˜å‚¨ç±»åž‹"`
	IsImage int `json:"isImage"  dc:"æ˜¯å¦å›¾ç‰‡"`
}

// FileCreateRes 创建æ–‡ä»¶è®°å½•响应
type FileCreateRes struct {
	g.Meta `mime:"application/json"`
}

// FileUpdateReq 更新æ–‡ä»¶è®°å½•请求
type FileUpdateReq struct {
	g.Meta `path:"/file/update" method:"put" tags:"æ–‡ä»¶è®°å½•" summary:"更新æ–‡ä»¶è®°å½•"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ–‡ä»¶è®°å½•ID"`
	DirID snowflake.JsonInt64 `json:"dirID" dc:"æ‰€å±žç›®å½•"`
	Name string `json:"name" dc:"æ–‡ä»¶åç§°"`
	URL string `json:"url" dc:"æ–‡ä»¶åœ°å€"`
	Ext string `json:"ext" dc:"æ–‡ä»¶æ‰©å±•å"`
	Size int64 `json:"size" dc:"æ–‡ä»¶å¤§å°"`
	Mime string `json:"mime" dc:"MIMEç±»åž‹"`
	Storage int `json:"storage" dc:"å­˜å‚¨ç±»åž‹"`
	IsImage int `json:"isImage" dc:"æ˜¯å¦å›¾ç‰‡"`
}

// FileUpdateRes 更新æ–‡ä»¶è®°å½•响应
type FileUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// FileDeleteReq 删除æ–‡ä»¶è®°å½•请求
type FileDeleteReq struct {
	g.Meta `path:"/file/delete" method:"delete" tags:"æ–‡ä»¶è®°å½•" summary:"删除æ–‡ä»¶è®°å½•"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ–‡ä»¶è®°å½•ID"`
}

// FileDeleteRes 删除æ–‡ä»¶è®°å½•响应
type FileDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// FileDetailReq 获取æ–‡ä»¶è®°å½•详情请求
type FileDetailReq struct {
	g.Meta `path:"/file/detail" method:"get" tags:"æ–‡ä»¶è®°å½•" summary:"获取æ–‡ä»¶è®°å½•详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ–‡ä»¶è®°å½•ID"`
}

// FileDetailRes 获取æ–‡ä»¶è®°å½•详情响应
type FileDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.FileDetailOutput
}

// FileListReq 获取æ–‡ä»¶è®°å½•列表请求
type FileListReq struct {
	g.Meta   `path:"/file/list" method:"get" tags:"æ–‡ä»¶è®°å½•" summary:"获取æ–‡ä»¶è®°å½•列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Storage int `json:"storage" dc:"å­˜å‚¨ç±»åž‹"`
	IsImage int `json:"isImage" dc:"æ˜¯å¦å›¾ç‰‡"`
}

// FileListRes 获取æ–‡ä»¶è®°å½•列表响应
type FileListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.FileListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

