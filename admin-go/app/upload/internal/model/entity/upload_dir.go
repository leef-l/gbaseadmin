// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadDir is the golang structure for table upload_dir.
type UploadDir struct {
	Id        uint64      `orm:"id"         description:"ID"`                       // ID
	ParentId  uint64      `orm:"parent_id"  description:"ä¸Šçº§ç›®å½•"`             // ä¸Šçº§ç›®å½•
	Name      string      `orm:"name"       description:"ç›®å½•åç§°"`             // ç›®å½•åç§°
	Path      string      `orm:"path"       description:"ç›®å½•è·¯å¾„"`             // ç›®å½•è·¯å¾„
	Sort      int         `orm:"sort"       description:"æŽ’åº"`                   // æŽ’åº
	Status    int         `orm:"status"     description:"çŠ¶æ€:0=ç¦ç”¨,1=å¯ç”¨"` // çŠ¶æ€:0=ç¦ç”¨,1=å¯ç”¨
	CreatedAt *gtime.Time `orm:"created_at" description:"åˆ›å»ºæ—¶é—´"`             // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time `orm:"updated_at" description:"æ›´æ–°æ—¶é—´"`             // æ›´æ–°æ—¶é—´
	DeletedAt *gtime.Time `orm:"deleted_at" description:"åˆ é™¤æ—¶é—´"`             // åˆ é™¤æ—¶é—´
	CreatedBy uint64      `orm:"created_by" description:"åˆ›å»ºäºº"`                // åˆ›å»ºäºº
	DeptId    uint64      `orm:"dept_id"    description:"éƒ¨é—¨ID"`                 // éƒ¨é—¨ID
}
