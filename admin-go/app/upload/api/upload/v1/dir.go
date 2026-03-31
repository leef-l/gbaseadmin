package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/upload/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Dir API

// DirCreateReq 创建æ–‡ä»¶ç›®å½•请求
type DirCreateReq struct {
	g.Meta `path:"/dir/create" method:"post" tags:"æ–‡ä»¶ç›®å½•" summary:"创建æ–‡ä»¶ç›®å½•"`
	ParentID snowflake.JsonInt64 `json:"parentID"  dc:"ä¸Šçº§ç›®å½•"`
	Name string `json:"name" v:"required#ç›®å½•åç§°不能为空" dc:"ç›®å½•åç§°"`
	Path string `json:"path" v:"required#ç›®å½•è·¯å¾„不能为空" dc:"ç›®å½•è·¯å¾„"`
	Sort int `json:"sort"  dc:"æŽ’åº"`
	Status int `json:"status"  dc:"çŠ¶æ€"`
}

// DirCreateRes 创建æ–‡ä»¶ç›®å½•响应
type DirCreateRes struct {
	g.Meta `mime:"application/json"`
}

// DirUpdateReq 更新æ–‡ä»¶ç›®å½•请求
type DirUpdateReq struct {
	g.Meta `path:"/dir/update" method:"put" tags:"æ–‡ä»¶ç›®å½•" summary:"更新æ–‡ä»¶ç›®å½•"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ–‡ä»¶ç›®å½•ID"`
	ParentID snowflake.JsonInt64 `json:"parentID" dc:"ä¸Šçº§ç›®å½•"`
	Name string `json:"name" dc:"ç›®å½•åç§°"`
	Path string `json:"path" dc:"ç›®å½•è·¯å¾„"`
	Sort int `json:"sort" dc:"æŽ’åº"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// DirUpdateRes 更新æ–‡ä»¶ç›®å½•响应
type DirUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// DirDeleteReq 删除æ–‡ä»¶ç›®å½•请求
type DirDeleteReq struct {
	g.Meta `path:"/dir/delete" method:"delete" tags:"æ–‡ä»¶ç›®å½•" summary:"删除æ–‡ä»¶ç›®å½•"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ–‡ä»¶ç›®å½•ID"`
}

// DirDeleteRes 删除æ–‡ä»¶ç›®å½•响应
type DirDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// DirDetailReq 获取æ–‡ä»¶ç›®å½•详情请求
type DirDetailReq struct {
	g.Meta `path:"/dir/detail" method:"get" tags:"æ–‡ä»¶ç›®å½•" summary:"获取æ–‡ä»¶ç›®å½•详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ–‡ä»¶ç›®å½•ID"`
}

// DirDetailRes 获取æ–‡ä»¶ç›®å½•详情响应
type DirDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.DirDetailOutput
}

// DirListReq 获取æ–‡ä»¶ç›®å½•列表请求
type DirListReq struct {
	g.Meta   `path:"/dir/list" method:"get" tags:"æ–‡ä»¶ç›®å½•" summary:"获取æ–‡ä»¶ç›®å½•列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// DirListRes 获取æ–‡ä»¶ç›®å½•列表响应
type DirListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.DirListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

// DirTreeReq 获取æ–‡ä»¶ç›®å½•树形结构请求
type DirTreeReq struct {
	g.Meta `path:"/dir/tree" method:"get" tags:"æ–‡ä»¶ç›®å½•" summary:"获取æ–‡ä»¶ç›®å½•树形结构"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// DirTreeRes 获取æ–‡ä»¶ç›®å½•树形结构响应
type DirTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.DirTreeOutput `json:"list" dc:"树形数据"`
}

