// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadConfig is the golang structure of table upload_config for DAO operations like Where/Data.
type UploadConfig struct {
	g.Meta       `orm:"table:upload_config, do:true"`
	Id           any         // ID
	Name         any         // é…ç½®åç§°
	Storage      any         // å­˜å‚¨ç±»åž‹:1=æœ¬åœ°,2=é˜¿é‡Œäº‘OSS,3=è…¾è®¯äº‘COS
	IsDefault    any         // æ˜¯å¦é»˜è®¤:0=å¦,1=æ˜¯
	LocalPath    any         // æœ¬åœ°å­˜å‚¨è·¯å¾„
	OssEndpoint  any         // OSS Endpoint
	OssBucket    any         // OSS Bucket
	OssAccessKey any         // OSS AccessKey
	OssSecretKey any         // OSS SecretKey
	CosRegion    any         // COS Region
	CosBucket    any         // COS Bucket
	CosSecretId  any         // COS SecretId
	CosSecretKey any         // COS SecretKey
	MaxSize      any         // æœ€å¤§æ–‡ä»¶å¤§å°(MB)
	Status       any         // çŠ¶æ€:0=ç¦ç”¨,1=å¯ç”¨
	CreatedAt    *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt    *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt    *gtime.Time // åˆ é™¤æ—¶é—´
	CreatedBy    any         // åˆ›å»ºäºº
	DeptId       any         // éƒ¨é—¨ID
}
