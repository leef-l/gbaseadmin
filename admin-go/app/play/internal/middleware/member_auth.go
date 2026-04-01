package middleware

import (
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"

	"gbaseadmin/utility/jwt"
	"gbaseadmin/utility/response"
)

// MemberAuth C端会员JWT鉴权中间件
func MemberAuth(r *ghttp.Request) {
	tokenStr := r.GetHeader("Authorization")
	if tokenStr == "" {
		response.Unauthorized(r)
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	tokenStr = strings.TrimSpace(tokenStr)
	if tokenStr == "" {
		response.Unauthorized(r)
		return
	}

	claims, err := jwt.ParseMemberToken(tokenStr)
	if err != nil {
		response.Unauthorized(r, "Token无效或已过期")
		return
	}

	r.SetCtxVar("jwt_member_id", claims.MemberID)
	r.SetCtxVar("jwt_phone", claims.Phone)
	r.SetCtxVar("jwt_is_coach", claims.IsCoach)
	r.SetCtxVar("jwt_coach_id", claims.CoachID)
	r.SetCtxVar("jwt_current_role", claims.CurrentRole)
	r.SetCtxVar("jwt_member_claims", claims)

	r.Middleware.Next()
}

// MemberAuthOptional 可选会员鉴权，有 token 就解析写入 ctx，没有或无效则跳过（不返回 401）
func MemberAuthOptional(r *ghttp.Request) {
	tokenStr := strings.TrimSpace(strings.TrimPrefix(r.GetHeader("Authorization"), "Bearer "))
	if tokenStr != "" {
		if claims, err := jwt.ParseMemberToken(tokenStr); err == nil {
			r.SetCtxVar("jwt_member_id", claims.MemberID)
			r.SetCtxVar("jwt_phone", claims.Phone)
			r.SetCtxVar("jwt_is_coach", claims.IsCoach)
			r.SetCtxVar("jwt_coach_id", claims.CoachID)
			r.SetCtxVar("jwt_current_role", claims.CurrentRole)
			r.SetCtxVar("jwt_member_claims", claims)
		}
	}
	r.Middleware.Next()
}

// CoachOnly 陪玩师身份校验中间件（需嵌套在MemberAuth之后）
func CoachOnly(r *ghttp.Request) {
	isCoach := r.GetCtxVar("jwt_is_coach").Int()
	currentRole := r.GetCtxVar("jwt_current_role").String()

	if isCoach != 1 || currentRole != "coach" {
		response.Forbidden(r, "需要切换到陪玩师身份")
		return
	}

	r.Middleware.Next()
}
