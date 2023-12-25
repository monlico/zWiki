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
	expireTime int64
	jwtSecret  = setting.JwtSecret
)

func GenerateToken(username, platform string, id uint) (string, error) {

	switch platform {
	case "pc":
		expireTime = int64(setting.PcExpireTime.Seconds())
	case "app":
		expireTime = time.Now().Add(setting.AppExpireTime).Unix()
	default:
		expireTime = int64(setting.PcExpireTime.Seconds())
	}

	claims := Claims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "wiki",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
