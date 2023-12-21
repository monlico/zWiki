package util

import (
	"time"
	"zWiki/pkg/setting"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	expireTime time.Duration
	jwtSecret  = setting.JwtSecret
)

func GenerateToken(username, platform string, id uint) (string, error) {

	switch platform {
	case "pc":
		expireTime = setting.PcExpireTime
	case "app":
		expireTime = setting.AppExpireTime
	default:
		expireTime = setting.PcExpireTime
	}

	nowTime := time.Now()

	claims := Claims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: nowTime.Add(expireTime).Unix(),
			Issuer:    "wiki",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
