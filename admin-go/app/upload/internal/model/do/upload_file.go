// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadFile is the golang structure of table upload_file for DAO operations like Where/Data.
type UploadFile struct {
	g.Meta    `orm:"table:upload_file, do:true"`
	Id        any         // ID
	DirId     any         // æ‰€å±žç›®å½•
	Name      any         // æ–‡ä»¶åç§°
	Url       any         // æ–‡ä»¶åœ°å€
	Ext       any         // æ–‡ä»¶æ‰©å±•å
	Size      any         // æ–‡ä»¶å¤§å°
	Mime      any         // MIMEç±»åž‹
	Storage   any         // å­˜å‚¨ç±»åž‹:1=æœ¬åœ°,2=é˜¿é‡Œäº‘OSS,3=è…¾è®¯äº‘COS
	IsImage   any         // æ˜¯å¦å›¾ç‰‡:0=å¦,1=æ˜¯
	CreatedAt *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt *gtime.Time // åˆ é™¤æ—¶é—´
	CreatedBy any         // åˆ›å»ºäºº
	DeptId    any         // éƒ¨é—¨ID
}
