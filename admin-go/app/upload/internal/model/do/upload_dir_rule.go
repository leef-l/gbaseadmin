// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadDirRule is the golang structure of table upload_dir_rule for DAO operations like Where/Data.
type UploadDirRule struct {
	g.Meta    `orm:"table:upload_dir_rule, do:true"`
	Id        any         // ID
	DirId     any         // ç›®å½•ID
	Category  any         // ç±»åˆ«:1=é»˜è®¤,2=ç±»åž‹,3=æŽ¥å£
	SavePath  any         // ä¿å­˜ç›®å½•
	Status    any         // çŠ¶æ€:0=ç¦ç”¨,1=å¯ç”¨
	CreatedAt *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt *gtime.Time // åˆ é™¤æ—¶é—´
	CreatedBy any         // åˆ›å»ºäºº
	DeptId    any         // éƒ¨é—¨ID
}
