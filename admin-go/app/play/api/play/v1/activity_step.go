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

// ActivityStepCreateReq 创建活动步骤表请求
type ActivityStepCreateReq struct {
	g.Meta `path:"/activity_step/create" method:"post" tags:"活动步骤表" summary:"创建活动步骤表"`
	ActivityID snowflake.JsonInt64 `json:"activityID" v:"required#活动ID不能为空" dc:"活动ID"`
	StepNum int `json:"stepNum"  dc:"步骤序号"`
	Title string `json:"title" v:"required#步骤标题不能为空" dc:"步骤标题"`
	DescContent string `json:"descContent"  dc:"步骤说明（富文本，支持图文）"`
	StepImage string `json:"stepImage"  dc:"步骤示例图片"`
	Sort int `json:"sort"  dc:"排序（升序）"`
}

// ActivityStepCreateRes 创建活动步骤表响应
type ActivityStepCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityStepUpdateReq 更新活动步骤表请求
type ActivityStepUpdateReq struct {
	g.Meta `path:"/activity_step/update" method:"put" tags:"活动步骤表" summary:"更新活动步骤表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动步骤表ID"`
	ActivityID snowflake.JsonInt64 `json:"activityID" dc:"活动ID"`
	StepNum int `json:"stepNum" dc:"步骤序号"`
	Title string `json:"title" dc:"步骤标题"`
	DescContent string `json:"descContent" dc:"步骤说明（富文本，支持图文）"`
	StepImage string `json:"stepImage" dc:"步骤示例图片"`
	Sort int `json:"sort" dc:"排序（升序）"`
}

// ActivityStepUpdateRes 更新活动步骤表响应
type ActivityStepUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityStepDeleteReq 删除活动步骤表请求
type ActivityStepDeleteReq struct {
	g.Meta `path:"/activity_step/delete" method:"delete" tags:"活动步骤表" summary:"删除活动步骤表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动步骤表ID"`
}

// ActivityStepDeleteRes 删除活动步骤表响应
type ActivityStepDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityStepDetailReq 获取活动步骤表详情请求
type ActivityStepDetailReq struct {
	g.Meta `path:"/activity_step/detail" method:"get" tags:"活动步骤表" summary:"获取活动步骤表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动步骤表ID"`
}

// ActivityStepDetailRes 获取活动步骤表详情响应
type ActivityStepDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ActivityStepDetailOutput
}

// ActivityStepListReq 获取活动步骤表列表请求
type ActivityStepListReq struct {
	g.Meta   `path:"/activity_step/list" method:"get" tags:"活动步骤表" summary:"获取活动步骤表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	ActivityID snowflake.JsonInt64 `json:"activityID" dc:"活动ID"`
}

// ActivityStepListRes 获取活动步骤表列表响应
type ActivityStepListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ActivityStepListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

