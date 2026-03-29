package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/utility/snowflake"
)

// Role API

// RoleCreateReq 创建角色表请求
type RoleCreateReq struct {
	g.Meta `path:"/role/create" method:"post" tags:"角色表" summary:"创建角色表"`
	ParentID snowflake.JsonInt64 `json:"parentID"  dc:"上级角色ID，0 表示顶级角色"`
	Title string `json:"title" v:"required#角色名称不能为空" dc:"角色名称"`
	DataScope int `json:"dataScope"  dc:"数据范围"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// RoleCreateRes 创建角色表响应
type RoleCreateRes struct {
	g.Meta `mime:"application/json"`
}

// RoleUpdateReq 更新角色表请求
type RoleUpdateReq struct {
	g.Meta `path:"/role/update" method:"put" tags:"角色表" summary:"更新角色表"`
	Id     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"角色表ID"`
	ParentID snowflake.JsonInt64 `json:"parentID" dc:"上级角色ID，0 表示顶级角色"`
	Title string `json:"title" dc:"角色名称"`
	DataScope int `json:"dataScope" dc:"数据范围"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// RoleUpdateRes 更新角色表响应
type RoleUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// RoleDeleteReq 删除角色表请求
type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"delete" tags:"角色表" summary:"删除角色表"`
	Id     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"角色表ID"`
}

// RoleDeleteRes 删除角色表响应
type RoleDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// RoleDetailReq 获取角色表详情请求
type RoleDetailReq struct {
	g.Meta `path:"/role/detail" method:"get" tags:"角色表" summary:"获取角色表详情"`
	Id     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"角色表ID"`
}

// RoleDetailRes 获取角色表详情响应
type RoleDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.RoleDetailOutput
}

// RoleListReq 获取角色表列表请求
type RoleListReq struct {
	g.Meta   `path:"/role/list" method:"get" tags:"角色表" summary:"获取角色表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	DataScope int `json:"dataScope" dc:"数据范围"`
	Status int `json:"status" dc:"状态"`
}

// RoleListRes 获取角色表列表响应
type RoleListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.RoleListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

// RoleTreeReq 获取角色表树形结构请求
type RoleTreeReq struct {
	g.Meta `path:"/role/tree" method:"get" tags:"角色表" summary:"获取角色表树形结构"`
	DataScope int `json:"dataScope" dc:"数据范围"`
	Status int `json:"status" dc:"状态"`
}

// RoleTreeRes 获取角色表树形结构响应
type RoleTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.RoleTreeOutput `json:"list" dc:"树形数据"`
}

