// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayOrder is the golang structure of table play_order for DAO operations like Where/Data.
type PlayOrder struct {
	g.Meta         `orm:"table:play_order, do:true"`
	Id             any         // è®¢å•IDï¼ˆSnowflakeï¼‰
	OrderNo        any         // è®¢å•ç¼–å·
	MemberId       any         // ä¸‹å•ä¼šå‘˜ID
	CoachId        any         // é™ªçŽ©å¸ˆID
	ShopId         any         // åº—é“ºID
	GoodsId        any         // å•†å“ID
	GoodsTitle     any         // å•†å“åç§°ï¼ˆå†—ä½™ï¼‰
	GoodsPrice     any         // å•†å“å•ä»·ï¼ˆåˆ†ï¼‰
	Quantity       any         // æ•°é‡
	TotalAmount    any         // è®¢å•æ€»é¢ï¼ˆåˆ†ï¼‰
	DiscountAmount any         // ä¼šå‘˜æŠ˜æ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰
	CouponAmount   any         // ä¼˜æƒ åˆ¸æŠµæ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰
	PayAmount      any         // å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰
	CouponMemberId any         // ä½¿ç”¨çš„ä¼˜æƒ åˆ¸é¢†å–è®°å½•ID
	PayType        any         // æ”¯ä»˜æ–¹å¼:0=æœªæ”¯ä»˜,1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜,3=ä½™é¢æ”¯ä»˜
	OrderStatus    any         // è®¢å•çŠ¶æ€:0=å¾…æ”¯ä»˜,1=å·²æ”¯ä»˜,2=è¿›è¡Œä¸­,3=å·²å®Œæˆ,4=å·²å–æ¶ˆ,5=é€€æ¬¾ä¸­,6=å·²é€€æ¬¾
	PayAt          *gtime.Time // æ”¯ä»˜æ—¶é—´
	StartAt        *gtime.Time // æœåŠ¡å¼€å§‹æ—¶é—´
	FinishAt       *gtime.Time // æœåŠ¡å®Œæˆæ—¶é—´
	CancelAt       *gtime.Time // å–æ¶ˆæ—¶é—´
	CancelReason   any         // å–æ¶ˆåŽŸå›
	Remark         any         // è®¢å•å¤‡æ³¨
	CreatedBy      any         // åˆ›å»ºäººID
	DeptId         any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt      *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt      *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
