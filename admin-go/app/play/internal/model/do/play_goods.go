// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayGoods is the golang structure of table play_goods for DAO operations like Where/Data.
type PlayGoods struct {
	g.Meta      `orm:"table:play_goods, do:true"`
	Id          any         // å•†å“IDï¼ˆSnowflakeï¼‰
	CategoryId  any         // åˆ†ç±»ID
	CoachId     any         // é™ªçŽ©å¸ˆID
	Title       any         // å•†å“åç§°
	CoverImage  any         // å•†å“å°é¢å›¾
	DescContent any         // å•†å“è¯¦æƒ…æè¿°
	Price       any         // å•ä»·ï¼ˆåˆ†ï¼‰
	Unit        any         // è®¡é‡å•ä½
	SalesNum    any         // é”€é‡
	Sort        any         // æŽ’åºï¼ˆå‡åºï¼‰
	Status      any         // çŠ¶æ€:0=ä¸‹æž¶,1=ä¸Šæž¶
	CreatedBy   any         // åˆ›å»ºäººID
	DeptId      any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt   *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt   *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
