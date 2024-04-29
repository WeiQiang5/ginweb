package util

import (
	"ginWeb/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id int, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		strconv.Itoa(id), username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginBlog",
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
