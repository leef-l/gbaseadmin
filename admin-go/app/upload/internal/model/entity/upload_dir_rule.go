// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadDirRule is the golang structure for table upload_dir_rule.
type UploadDirRule struct {
	Id        uint64      `orm:"id"         description:"ID"`                                // ID
	DirId     uint64      `orm:"dir_id"     description:"ç›®å½•ID"`                          // ç›®å½•ID
	Category  int         `orm:"category"   description:"ç±»åˆ«:1=é»˜è®¤,2=ç±»åž‹,3=æŽ¥å£"` // ç±»åˆ«:1=é»˜è®¤,2=ç±»åž‹,3=æŽ¥å£
	SavePath  string      `orm:"save_path"  description:"ä¿å­˜ç›®å½•"`                      // ä¿å­˜ç›®å½•
	Status    int         `orm:"status"     description:"çŠ¶æ€:0=ç¦ç”¨,1=å¯ç”¨"`          // çŠ¶æ€:0=ç¦ç”¨,1=å¯ç”¨
	CreatedAt *gtime.Time `orm:"created_at" description:"åˆ›å»ºæ—¶é—´"`                      // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time `orm:"updated_at" description:"æ›´æ–°æ—¶é—´"`                      // æ›´æ–°æ—¶é—´
	DeletedAt *gtime.Time `orm:"deleted_at" description:"åˆ é™¤æ—¶é—´"`                      // åˆ é™¤æ—¶é—´
	CreatedBy uint64      `orm:"created_by" description:"åˆ›å»ºäºº"`                         // åˆ›å»ºäºº
	DeptId    uint64      `orm:"dept_id"    description:"éƒ¨é—¨ID"`                          // éƒ¨é—¨ID
}
