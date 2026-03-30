// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayOauth is the golang structure of table play_oauth for DAO operations like Where/Data.
type PlayOauth struct {
	g.Meta       `orm:"table:play_oauth, do:true"`
	Id           any         // è®°å½•IDï¼ˆSnowflakeï¼‰
	MemberId     any         // ä¼šå‘˜ID
	Provider     any         // ç¬¬ä¸‰æ–¹å¹³å°:1=å¾®ä¿¡,2=æ”¯ä»˜å®
	OpenId       any         // ç¬¬ä¸‰æ–¹OpenID
	UnionId      any         // ç¬¬ä¸‰æ–¹UnionID
	Nickname     any         // ç¬¬ä¸‰æ–¹æ˜µç§°
	Avatar       any         // ç¬¬ä¸‰æ–¹å¤´åƒ
	AccessToken  any         // è®¿é—®ä»¤ç‰Œ
	RefreshToken any         // åˆ·æ–°ä»¤ç‰Œ
	ExpireAt     *gtime.Time // ä»¤ç‰Œè¿‡æœŸæ—¶é—´
	CreatedBy    any         // åˆ›å»ºäººID
	DeptId       any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt    *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt    *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt    *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
