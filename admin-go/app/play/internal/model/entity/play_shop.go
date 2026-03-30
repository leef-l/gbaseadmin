// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayShop is the golang structure for table play_shop.
type PlayShop struct {
	Id             uint64      `orm:"id"              description:"店铺ID（Snowflake）"`         // 店铺ID（Snowflake）
	Title          string      `orm:"title"           description:"店铺名称"`                    // 店铺名称
	LogoImage      string      `orm:"logo_image"      description:"店铺LOGO"`                  // 店铺LOGO
	CoverImage     string      `orm:"cover_image"     description:"封面图"`                     // 封面图
	ContactName    string      `orm:"contact_name"    description:"联系人姓名"`                   // 联系人姓名
	ContactPhone   string      `orm:"contact_phone"   description:"联系电话"`                    // 联系电话
	Intro          string      `orm:"intro"           description:"店铺简介"`                    // 店铺简介
	CommissionRate int         `orm:"commission_rate" description:"店铺抽成比例（百分比，如 10 表示 10%）"` // 店铺抽成比例（百分比，如 10 表示 10%）
	CoachNum       int         `orm:"coach_num"       description:"陪玩师数量"`                   // 陪玩师数量
	Sort           int         `orm:"sort"            description:"排序（升序）"`                  // 排序（升序）
	Status         int         `orm:"status"          description:"状态:0=关闭,1=开启"`            // 状态:0=关闭,1=开启
	CreatedBy      uint64      `orm:"created_by"      description:"创建人ID"`                   // 创建人ID
	DeptId         uint64      `orm:"dept_id"         description:"所属部门ID"`                  // 所属部门ID
	CreatedAt      *gtime.Time `orm:"created_at"      description:"创建时间"`                    // 创建时间
	UpdatedAt      *gtime.Time `orm:"updated_at"      description:"更新时间"`                    // 更新时间
	DeletedAt      *gtime.Time `orm:"deleted_at"      description:"软删除时间"`                   // 软删除时间
}
