package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// ActivityStep API

// ActivityStepCreateReq 创建æ´»åŠ¨æ­¥éª¤è¡¨请求
type ActivityStepCreateReq struct {
	g.Meta `path:"/activity_step/create" method:"post" tags:"æ´»åŠ¨æ­¥éª¤è¡¨" summary:"创建æ´»åŠ¨æ­¥éª¤è¡¨"`
	ActivityID snowflake.JsonInt64 `json:"activityID" v:"required#æ´»åŠ¨ID不能为空" dc:"æ´»åŠ¨ID"`
	StepNum int `json:"stepNum"  dc:"æ­¥éª¤åºå·"`
	Title string `json:"title" v:"required#æ­¥éª¤æ ‡é¢˜不能为空" dc:"æ­¥éª¤æ ‡é¢˜"`
	DescContent string `json:"descContent"  dc:"æ­¥éª¤è¯´æ˜Ž"`
	StepImage string `json:"stepImage"  dc:"æ­¥éª¤ç¤ºä¾‹å›¾ç‰‡"`
	Sort int `json:"sort"  dc:"æŽ’åº"`
}

// ActivityStepCreateRes 创建æ´»åŠ¨æ­¥éª¤è¡¨响应
type ActivityStepCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityStepUpdateReq 更新æ´»åŠ¨æ­¥éª¤è¡¨请求
type ActivityStepUpdateReq struct {
	g.Meta `path:"/activity_step/update" method:"put" tags:"æ´»åŠ¨æ­¥éª¤è¡¨" summary:"更新æ´»åŠ¨æ­¥éª¤è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨æ­¥éª¤è¡¨ID"`
	ActivityID snowflake.JsonInt64 `json:"activityID" dc:"æ´»åŠ¨ID"`
	StepNum int `json:"stepNum" dc:"æ­¥éª¤åºå·"`
	Title string `json:"title" dc:"æ­¥éª¤æ ‡é¢˜"`
	DescContent string `json:"descContent" dc:"æ­¥éª¤è¯´æ˜Ž"`
	StepImage string `json:"stepImage" dc:"æ­¥éª¤ç¤ºä¾‹å›¾ç‰‡"`
	Sort int `json:"sort" dc:"æŽ’åº"`
}

// ActivityStepUpdateRes 更新æ´»åŠ¨æ­¥éª¤è¡¨响应
type ActivityStepUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityStepDeleteReq 删除æ´»åŠ¨æ­¥éª¤è¡¨请求
type ActivityStepDeleteReq struct {
	g.Meta `path:"/activity_step/delete" method:"delete" tags:"æ´»åŠ¨æ­¥éª¤è¡¨" summary:"删除æ´»åŠ¨æ­¥éª¤è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨æ­¥éª¤è¡¨ID"`
}

// ActivityStepDeleteRes 删除æ´»åŠ¨æ­¥éª¤è¡¨响应
type ActivityStepDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityStepDetailReq 获取æ´»åŠ¨æ­¥éª¤è¡¨详情请求
type ActivityStepDetailReq struct {
	g.Meta `path:"/activity_step/detail" method:"get" tags:"æ´»åŠ¨æ­¥éª¤è¡¨" summary:"获取æ´»åŠ¨æ­¥éª¤è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"æ´»åŠ¨æ­¥éª¤è¡¨ID"`
}

// ActivityStepDetailRes 获取æ´»åŠ¨æ­¥éª¤è¡¨详情响应
type ActivityStepDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ActivityStepDetailOutput
}

// ActivityStepListReq 获取æ´»åŠ¨æ­¥éª¤è¡¨列表请求
type ActivityStepListReq struct {
	g.Meta   `path:"/activity_step/list" method:"get" tags:"æ´»åŠ¨æ­¥éª¤è¡¨" summary:"获取æ´»åŠ¨æ­¥éª¤è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
}

// ActivityStepListRes 获取æ´»åŠ¨æ­¥éª¤è¡¨列表响应
type ActivityStepListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ActivityStepListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

