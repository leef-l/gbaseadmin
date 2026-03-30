// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayGoods is the golang structure for table play_goods.
type PlayGoods struct {
	Id          uint64      `orm:"id"           description:"å•†å“IDï¼ˆSnowflakeï¼‰"`  // å•†å“IDï¼ˆSnowflakeï¼‰
	CategoryId  uint64      `orm:"category_id"  description:"åˆ†ç±»ID"`                 // åˆ†ç±»ID
	CoachId     uint64      `orm:"coach_id"     description:"é™ªçŽ©å¸ˆID"`              // é™ªçŽ©å¸ˆID
	Title       string      `orm:"title"        description:"å•†å“åç§°"`             // å•†å“åç§°
	CoverImage  string      `orm:"cover_image"  description:"å•†å“å°é¢å›¾"`          // å•†å“å°é¢å›¾
	DescContent string      `orm:"desc_content" description:"å•†å“è¯¦æƒ…æè¿°"`       // å•†å“è¯¦æƒ…æè¿°
	Price       int64       `orm:"price"        description:"å•ä»·ï¼ˆåˆ†ï¼‰"`          // å•ä»·ï¼ˆåˆ†ï¼‰
	Unit        string      `orm:"unit"         description:"è®¡é‡å•ä½"`             // è®¡é‡å•ä½
	SalesNum    int         `orm:"sales_num"    description:"é”€é‡"`                   // é”€é‡
	Sort        int         `orm:"sort"         description:"æŽ’åºï¼ˆå‡åºï¼‰"`       // æŽ’åºï¼ˆå‡åºï¼‰
	Status      int         `orm:"status"       description:"çŠ¶æ€:0=ä¸‹æž¶,1=ä¸Šæž¶"` // çŠ¶æ€:0=ä¸‹æž¶,1=ä¸Šæž¶
	CreatedBy   uint64      `orm:"created_by"   description:"åˆ›å»ºäººID"`              // åˆ›å»ºäººID
	DeptId      uint64      `orm:"dept_id"      description:"æ‰€å±žéƒ¨é—¨ID"`           // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   *gtime.Time `orm:"created_at"   description:"åˆ›å»ºæ—¶é—´"`             // åˆ›å»ºæ—¶é—´
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"æ›´æ–°æ—¶é—´"`             // æ›´æ–°æ—¶é—´
	DeletedAt   *gtime.Time `orm:"deleted_at"   description:"è½¯åˆ é™¤æ—¶é—´"`          // è½¯åˆ é™¤æ—¶é—´
}
