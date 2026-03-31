// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadDir is the golang structure of table upload_dir for DAO operations like Where/Data.
type UploadDir struct {
	g.Meta    `orm:"table:upload_dir, do:true"`
	Id        any         // ID
	ParentId  any         // ä¸Šçº§ç›®å½•
	Name      any         // ç›®å½•åç§°
	Path      any         // ç›®å½•è·¯å¾„
	Sort      any         // æŽ’åº
	Status    any         // çŠ¶æ€:0=ç¦ç”¨,1=å¯ç”¨
	CreatedAt *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt *gtime.Time // åˆ é™¤æ—¶é—´
	CreatedBy any         // åˆ›å»ºäºº
	DeptId    any         // éƒ¨é—¨ID
}
