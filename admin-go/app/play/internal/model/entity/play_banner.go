// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayBanner is the golang structure for table play_banner.
type PlayBanner struct {
	Id        int64       `orm:"id"         description:"Banner ID"`                            // Banner ID
	Title     string      `orm:"title"      description:"Banner标题"`                             // Banner标题
	Image     string      `orm:"image"      description:"图片URL"`                                // 图片URL
	LinkType  int         `orm:"link_type"  description:"跳转类型: 1内页 2外链 3活动页 4商品页 5陪玩师页 6唤醒App"` // 跳转类型: 1内页 2外链 3活动页 4商品页 5陪玩师页 6唤醒App
	LinkValue string      `orm:"link_value" description:"跳转值(页面路径/URL/业务ID/App Scheme)"`        // 跳转值(页面路径/URL/业务ID/App Scheme)
	Sort      int         `orm:"sort"       description:"排序(越大越前)"`                             // 排序(越大越前)
	Status    int         `orm:"status"     description:"状态: 0禁用 1启用"`                          // 状态: 0禁用 1启用
	StartTime *gtime.Time `orm:"start_time" description:"生效开始时间"`                               // 生效开始时间
	EndTime   *gtime.Time `orm:"end_time"   description:"生效结束时间"`                               // 生效结束时间
	Remark    string      `orm:"remark"     description:"备注"`                                   // 备注
	CreatedBy int64       `orm:"created_by" description:"创建人"`                                  // 创建人
	DeptId    int64       `orm:"dept_id"    description:"部门ID"`                                 // 部门ID
	CreatedAt *gtime.Time `orm:"created_at" description:"创建时间"`                                 // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at" description:"更新时间"`                                 // 更新时间
	DeletedAt *gtime.Time `orm:"deleted_at" description:"删除时间"`                                 // 删除时间
}
