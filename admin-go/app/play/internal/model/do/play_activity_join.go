// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityJoin is the golang structure of table play_activity_join for DAO operations like Where/Data.
type PlayActivityJoin struct {
	g.Meta      `orm:"table:play_activity_join, do:true"`
	Id          any         // è®°å½•IDï¼ˆSnowflakeï¼‰
	ActivityId  any         // æ´»åŠ¨ID
	MemberId    any         // ä¼šå‘˜ID
	JoinStatus  any         // å‚ä¸ŽçŠ¶æ€:0=å·²æŠ¥å,1=è¿›è¡Œä¸­,2=å·²å®Œæˆ,3=å·²é¢†å¥–
	CurrentStep any         // å½“å‰å®Œæˆåˆ°ç¬¬å‡ æ­¥
	FinishAt    *gtime.Time // å®Œæˆæ—¶é—´
	RewardAt    *gtime.Time // é¢†å¥–æ—¶é—´
	Remark      any         // å¤‡æ³¨
	CreatedBy   any         // åˆ›å»ºäººID
	DeptId      any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt   *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt   *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
