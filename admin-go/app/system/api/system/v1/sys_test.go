package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/system/internal/model"
	"gbaseadmin/utility/snowflake"
)

// SysTest API

// SysTestCreateReq 创建æµ‹è¯•è¡¨请求
type SysTestCreateReq struct {
	g.Meta `path:"/sys_test/create" method:"post" tags:"æµ‹è¯•è¡¨" summary:"创建æµ‹è¯•è¡¨"`
	ParentID snowflake.JsonInt64 `json:"parentID"  dc:"ä¸Šçº§IDï¼Œ0è¡¨ç¤ºé¡¶çº§"`
	Title string `json:"title"  dc:"åç§°"`
	Code string `json:"code"  dc:"ç¼–ç "`
	Type int `json:"type"  dc:"ç±»åž‹"`
	Status int `json:"status"  dc:"çŠ¶æ€"`
	Sort int `json:"sort"  dc:"æŽ’åº"`
	Remark string `json:"remark"  dc:"å¤‡æ³¨"`
}

// SysTestCreateRes 创建æµ‹è¯•è¡¨响应
type SysTestCreateRes struct {
	g.Meta `mime:"application/json"`
}

// SysTestUpdateReq 更新æµ‹è¯•è¡¨请求
type SysTestUpdateReq struct {
	g.Meta `path:"/sys_test/update" method:"put" tags:"æµ‹è¯•è¡¨" summary:"更新æµ‹è¯•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æµ‹è¯•è¡¨ID"`
	ParentID snowflake.JsonInt64 `json:"parentID" dc:"ä¸Šçº§IDï¼Œ0è¡¨ç¤ºé¡¶çº§"`
	Title string `json:"title" dc:"åç§°"`
	Code string `json:"code" dc:"ç¼–ç "`
	Type int `json:"type" dc:"ç±»åž‹"`
	Status int `json:"status" dc:"çŠ¶æ€"`
	Sort int `json:"sort" dc:"æŽ’åº"`
	Remark string `json:"remark" dc:"å¤‡æ³¨"`
}

// SysTestUpdateRes 更新æµ‹è¯•è¡¨响应
type SysTestUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// SysTestDeleteReq 删除æµ‹è¯•è¡¨请求
type SysTestDeleteReq struct {
	g.Meta `path:"/sys_test/delete" method:"delete" tags:"æµ‹è¯•è¡¨" summary:"删除æµ‹è¯•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æµ‹è¯•è¡¨ID"`
}

// SysTestDeleteRes 删除æµ‹è¯•è¡¨响应
type SysTestDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// SysTestDetailReq 获取æµ‹è¯•è¡¨详情请求
type SysTestDetailReq struct {
	g.Meta `path:"/sys_test/detail" method:"get" tags:"æµ‹è¯•è¡¨" summary:"获取æµ‹è¯•è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æµ‹è¯•è¡¨ID"`
}

// SysTestDetailRes 获取æµ‹è¯•è¡¨详情响应
type SysTestDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.SysTestDetailOutput
}

// SysTestListReq 获取æµ‹è¯•è¡¨列表请求
type SysTestListReq struct {
	g.Meta   `path:"/sys_test/list" method:"get" tags:"æµ‹è¯•è¡¨" summary:"获取æµ‹è¯•è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Type int `json:"type" dc:"ç±»åž‹"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// SysTestListRes 获取æµ‹è¯•è¡¨列表响应
type SysTestListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.SysTestListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

// SysTestTreeReq 获取æµ‹è¯•è¡¨树形结构请求
type SysTestTreeReq struct {
	g.Meta `path:"/sys_test/tree" method:"get" tags:"æµ‹è¯•è¡¨" summary:"获取æµ‹è¯•è¡¨树形结构"`
	Type int `json:"type" dc:"ç±»åž‹"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// SysTestTreeRes 获取æµ‹è¯•è¡¨树形结构响应
type SysTestTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.SysTestTreeOutput `json:"list" dc:"树形数据"`
}

