package jwt

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	gojwt "github.com/golang-jwt/jwt/v5"
)

// Claims 自定义 JWT 载荷
type Claims struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
	DeptID   int64  `json:"deptId"`
	gojwt.RegisteredClaims
}

var (
	secret     []byte
	expireTime time.Duration
)

func init() {
	ctx := gctx.New()
	key, _ := g.Cfg().Get(ctx, "jwt.secret", "gbaseadmin-secret-key")
	secret = []byte(key.String())
	hours, _ := g.Cfg().Get(ctx, "jwt.expire", 24)
	expireTime = time.Duration(hours.Int()) * time.Hour
}

// GenerateToken 生成 JWT Token
func GenerateToken(userID int64, username string, deptID int64) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:   userID,
		Username: username,
		DeptID:   deptID,
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(now.Add(expireTime)),
			IssuedAt:  gojwt.NewNumericDate(now),
			Issuer:    "gbaseadmin",
		},
	}
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// ParseToken 解析 JWT Token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := gojwt.ParseWithClaims(tokenStr, &Claims{}, func(t *gojwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, gojwt.ErrTokenInvalidClaims
}
