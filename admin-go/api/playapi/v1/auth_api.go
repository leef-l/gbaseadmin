package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AuthLoginReq 手机号验证码登录
type AuthLoginReq struct {
	g.Meta `path:"/auth/login" method:"post" tags:"C端认证" summary:"手机号验证码登录"`
	Phone  string `json:"phone" v:"required|phone#手机号不能为空|手机号格式不正确" dc:"手机号"`
	Code   string `json:"code" v:"required|length:4,6#验证码不能为空|验证码长度4-6位" dc:"短信验证码"`
}

type AuthLoginRes struct {
	g.Meta       `mime:"application/json"`
	Token        string `json:"token" dc:"访问令牌"`
	RefreshToken string `json:"refreshToken" dc:"刷新令牌"`
	ExpiresIn    int64  `json:"expiresIn" dc:"过期时间(秒)"`
	IsNew        bool   `json:"isNew" dc:"是否新注册用户"`
}

// AuthSendCodeReq 发送验证码
type AuthSendCodeReq struct {
	g.Meta `path:"/auth/send_code" method:"post" tags:"C端认证" summary:"发送验证码"`
	Phone  string `json:"phone" v:"required|phone#手机号不能为空|手机号格式不正确" dc:"手机号"`
	Scene  string `json:"scene" v:"required|in:login,bindPhone#场景不能为空|场景值不合法" dc:"场景:login=登录,bindPhone=绑定手机"`
}

type AuthSendCodeRes struct {
	g.Meta `mime:"application/json"`
}

// AuthRefreshTokenReq 刷新Token
type AuthRefreshTokenReq struct {
	g.Meta       `path:"/auth/refresh_token" method:"post" tags:"C端认证" summary:"刷新Token"`
	RefreshToken string `json:"refreshToken" v:"required#刷新令牌不能为空" dc:"刷新令牌"`
}

type AuthRefreshTokenRes struct {
	g.Meta       `mime:"application/json"`
	Token        string `json:"token" dc:"新访问令牌"`
	RefreshToken string `json:"refreshToken" dc:"新刷新令牌"`
	ExpiresIn    int64  `json:"expiresIn" dc:"过期时间(秒)"`
}

// AuthWxLoginReq 微信登录
type AuthWxLoginReq struct {
	g.Meta `path:"/auth/wx_login" method:"post" tags:"C端认证" summary:"微信登录"`
	Code   string `json:"code" v:"required#微信授权码不能为空" dc:"微信授权code"`
}

type AuthWxLoginRes struct {
	g.Meta        `mime:"application/json"`
	Token         string `json:"token" dc:"访问令牌"`
	RefreshToken  string `json:"refreshToken" dc:"刷新令牌"`
	ExpiresIn     int64  `json:"expiresIn" dc:"过期时间(秒)"`
	IsNew         bool   `json:"isNew" dc:"是否新注册用户"`
	NeedBindPhone bool   `json:"needBindPhone" dc:"是否需要绑定手机号"`
}

// AuthAlipayLoginReq 支付宝登录
type AuthAlipayLoginReq struct {
	g.Meta   `path:"/auth/alipay_login" method:"post" tags:"C端认证" summary:"支付宝登录"`
	AuthCode string `json:"authCode" v:"required#支付宝授权码不能为空" dc:"支付宝授权code"`
}

type AuthAlipayLoginRes struct {
	g.Meta        `mime:"application/json"`
	Token         string `json:"token" dc:"访问令牌"`
	RefreshToken  string `json:"refreshToken" dc:"刷新令牌"`
	ExpiresIn     int64  `json:"expiresIn" dc:"过期时间(秒)"`
	IsNew         bool   `json:"isNew" dc:"是否新注册用户"`
	NeedBindPhone bool   `json:"needBindPhone" dc:"是否需要绑定手机号"`
}
