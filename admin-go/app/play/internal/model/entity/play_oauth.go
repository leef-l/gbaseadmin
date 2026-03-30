// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayOauth is the golang structure for table play_oauth.
type PlayOauth struct {
	Id           uint64      `orm:"id"            description:"è®°å½•IDï¼ˆSnowflakeï¼‰"`              // è®°å½•IDï¼ˆSnowflakeï¼‰
	MemberId     uint64      `orm:"member_id"     description:"ä¼šå‘˜ID"`                             // ä¼šå‘˜ID
	Provider     int         `orm:"provider"      description:"ç¬¬ä¸‰æ–¹å¹³å°:1=å¾®ä¿¡,2=æ”¯ä»˜å®"` // ç¬¬ä¸‰æ–¹å¹³å°:1=å¾®ä¿¡,2=æ”¯ä»˜å®
	OpenId       string      `orm:"open_id"       description:"ç¬¬ä¸‰æ–¹OpenID"`                      // ç¬¬ä¸‰æ–¹OpenID
	UnionId      string      `orm:"union_id"      description:"ç¬¬ä¸‰æ–¹UnionID"`                     // ç¬¬ä¸‰æ–¹UnionID
	Nickname     string      `orm:"nickname"      description:"ç¬¬ä¸‰æ–¹æ˜µç§°"`                      // ç¬¬ä¸‰æ–¹æ˜µç§°
	Avatar       string      `orm:"avatar"        description:"ç¬¬ä¸‰æ–¹å¤´åƒ"`                      // ç¬¬ä¸‰æ–¹å¤´åƒ
	AccessToken  string      `orm:"access_token"  description:"è®¿é—®ä»¤ç‰Œ"`                         // è®¿é—®ä»¤ç‰Œ
	RefreshToken string      `orm:"refresh_token" description:"åˆ·æ–°ä»¤ç‰Œ"`                         // åˆ·æ–°ä»¤ç‰Œ
	ExpireAt     *gtime.Time `orm:"expire_at"     description:"ä»¤ç‰Œè¿‡æœŸæ—¶é—´"`                   // ä»¤ç‰Œè¿‡æœŸæ—¶é—´
	CreatedBy    uint64      `orm:"created_by"    description:"åˆ›å»ºäººID"`                          // åˆ›å»ºäººID
	DeptId       uint64      `orm:"dept_id"       description:"æ‰€å±žéƒ¨é—¨ID"`                       // æ‰€å±žéƒ¨é—¨ID
	CreatedAt    *gtime.Time `orm:"created_at"    description:"åˆ›å»ºæ—¶é—´"`                         // åˆ›å»ºæ—¶é—´
	UpdatedAt    *gtime.Time `orm:"updated_at"    description:"æ›´æ–°æ—¶é—´"`                         // æ›´æ–°æ—¶é—´
	DeletedAt    *gtime.Time `orm:"deleted_at"    description:"è½¯åˆ é™¤æ—¶é—´"`                      // è½¯åˆ é™¤æ—¶é—´
}
