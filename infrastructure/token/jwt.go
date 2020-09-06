package token

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var tokenKey = "90036a7e01159610b81215d88c63becd"
var TokenDuration = int64(864000)

type TokenData struct {
	UserId     string `json:"user_id"`
	ExpireTime int64  `json:"expire_time"`
	Token      string `json:"token"`
}

func CreateToken(m *TokenData) string {
	token := jwt.New(jwt.SigningMethodHS256)
	m.ExpireTime = time.Now().Unix() + TokenDuration
	claims := make(jwt.MapClaims)
	claims["user_id"] = m.UserId
	claims["expire_time"] = strconv.FormatInt(m.ExpireTime, 10)
	// fmt.Println(_map)
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(tokenKey))
	return tokenString
}

func ParseToken(tokenString string) (interface{}, bool) {

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return "", false
	}
}
