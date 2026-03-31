// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadFile is the golang structure for table upload_file.
type UploadFile struct {
	Id        uint64      `orm:"id"         description:"ID"`                                                  // ID
	DirId     uint64      `orm:"dir_id"     description:"æ‰€å±žç›®å½•"`                                        // æ‰€å±žç›®å½•
	Name      string      `orm:"name"       description:"æ–‡ä»¶åç§°"`                                        // æ–‡ä»¶åç§°
	Url       string      `orm:"url"        description:"æ–‡ä»¶åœ°å€"`                                        // æ–‡ä»¶åœ°å€
	Ext       string      `orm:"ext"        description:"æ–‡ä»¶æ‰©å±•å"`                                     // æ–‡ä»¶æ‰©å±•å
	Size      uint64      `orm:"size"       description:"æ–‡ä»¶å¤§å°"`                                        // æ–‡ä»¶å¤§å°
	Mime      string      `orm:"mime"       description:"MIMEç±»åž‹"`                                          // MIMEç±»åž‹
	Storage   int         `orm:"storage"    description:"å­˜å‚¨ç±»åž‹:1=æœ¬åœ°,2=é˜¿é‡Œäº‘OSS,3=è…¾è®¯äº‘COS"` // å­˜å‚¨ç±»åž‹:1=æœ¬åœ°,2=é˜¿é‡Œäº‘OSS,3=è…¾è®¯äº‘COS
	IsImage   int         `orm:"is_image"   description:"æ˜¯å¦å›¾ç‰‡:0=å¦,1=æ˜¯"`                            // æ˜¯å¦å›¾ç‰‡:0=å¦,1=æ˜¯
	CreatedAt *gtime.Time `orm:"created_at" description:"åˆ›å»ºæ—¶é—´"`                                        // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time `orm:"updated_at" description:"æ›´æ–°æ—¶é—´"`                                        // æ›´æ–°æ—¶é—´
	DeletedAt *gtime.Time `orm:"deleted_at" description:"åˆ é™¤æ—¶é—´"`                                        // åˆ é™¤æ—¶é—´
	CreatedBy uint64      `orm:"created_by" description:"åˆ›å»ºäºº"`                                           // åˆ›å»ºäºº
	DeptId    uint64      `orm:"dept_id"    description:"éƒ¨é—¨ID"`                                            // éƒ¨é—¨ID
}
