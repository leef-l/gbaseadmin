// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayBanner is the golang structure of table play_banner for DAO operations like Where/Data.
type PlayBanner struct {
	g.Meta    `orm:"table:play_banner, do:true"`
	Id        any         // Banner ID
	Title     any         // Banner标题
	Image     any         // 图片URL
	LinkType  any         // 跳转类型: 1内页 2外链 3活动页 4商品页 5陪玩师页 6唤醒App
	LinkValue any         // 跳转值(页面路径/URL/业务ID/App Scheme)
	Sort      any         // 排序(越大越前)
	Status    any         // 状态: 0禁用 1启用
	StartTime *gtime.Time // 生效开始时间
	EndTime   *gtime.Time // 生效结束时间
	Remark    any         // 备注
	CreatedBy any         // 创建人
	DeptId    any         // 部门ID
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
