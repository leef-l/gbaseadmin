package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Goods API

// GoodsCreateReq 创建å•†å“è¡¨请求
type GoodsCreateReq struct {
	g.Meta `path:"/goods/create" method:"post" tags:"å•†å“è¡¨" summary:"创建å•†å“è¡¨"`
	CategoryID snowflake.JsonInt64 `json:"categoryID" v:"required#åˆ†ç±»ID不能为空" dc:"åˆ†ç±»ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" v:"required#é™ªçŽ©å¸ˆID不能为空" dc:"é™ªçŽ©å¸ˆID"`
	Title string `json:"title" v:"required#å•†å“åç§°不能为空" dc:"å•†å“åç§°"`
	CoverImage string `json:"coverImage"  dc:"å•†å“å°é¢å›¾"`
	DescContent string `json:"descContent"  dc:"å•†å“è¯¦æƒ…æè¿°"`
	Price int64 `json:"price"  dc:"å•ä»·ï¼ˆåˆ†ï¼‰"`
	Unit string `json:"unit"  dc:"è®¡é‡å•ä½"`
	SalesNum int `json:"salesNum"  dc:"é”€é‡"`
	Sort int `json:"sort"  dc:"æŽ’åºï¼ˆå‡åºï¼‰"`
	Status int `json:"status"  dc:"çŠ¶æ€"`
}

// GoodsCreateRes 创建å•†å“è¡¨响应
type GoodsCreateRes struct {
	g.Meta `mime:"application/json"`
}

// GoodsUpdateReq 更新å•†å“è¡¨请求
type GoodsUpdateReq struct {
	g.Meta `path:"/goods/update" method:"put" tags:"å•†å“è¡¨" summary:"更新å•†å“è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"å•†å“è¡¨ID"`
	CategoryID snowflake.JsonInt64 `json:"categoryID" dc:"åˆ†ç±»ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" dc:"é™ªçŽ©å¸ˆID"`
	Title string `json:"title" dc:"å•†å“åç§°"`
	CoverImage string `json:"coverImage" dc:"å•†å“å°é¢å›¾"`
	DescContent string `json:"descContent" dc:"å•†å“è¯¦æƒ…æè¿°"`
	Price int64 `json:"price" dc:"å•ä»·ï¼ˆåˆ†ï¼‰"`
	Unit string `json:"unit" dc:"è®¡é‡å•ä½"`
	SalesNum int `json:"salesNum" dc:"é”€é‡"`
	Sort int `json:"sort" dc:"æŽ’åºï¼ˆå‡åºï¼‰"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// GoodsUpdateRes 更新å•†å“è¡¨响应
type GoodsUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// GoodsDeleteReq 删除å•†å“è¡¨请求
type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" method:"delete" tags:"å•†å“è¡¨" summary:"删除å•†å“è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"å•†å“è¡¨ID"`
}

// GoodsDeleteRes 删除å•†å“è¡¨响应
type GoodsDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// GoodsDetailReq 获取å•†å“è¡¨详情请求
type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail" method:"get" tags:"å•†å“è¡¨" summary:"获取å•†å“è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"å•†å“è¡¨ID"`
}

// GoodsDetailRes 获取å•†å“è¡¨详情响应
type GoodsDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.GoodsDetailOutput
}

// GoodsListReq 获取å•†å“è¡¨列表请求
type GoodsListReq struct {
	g.Meta   `path:"/goods/list" method:"get" tags:"å•†å“è¡¨" summary:"获取å•†å“è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// GoodsListRes 获取å•†å“è¡¨列表响应
type GoodsListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.GoodsListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

