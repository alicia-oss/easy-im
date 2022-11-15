package jwt

import (
	"easy_im/pkg/log"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type ImClaims struct {
	UserId             uint64 `json:"user_id"`
	UserName           string `json:"user_name"`
	jwt.StandardClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

var secret = []byte("jinx") // 定义secret，后面会用到
var expireTime = 2 * time.Hour

func GenToken(userId uint64, username string) (string, error) {
	claim := ImClaims{
		UserId:   userId,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	signedString, err := token.SignedString(secret)
	if err != nil {
		log.Error(fmt.Sprintf("gen token err:%v", err), "token")
	}
	return signedString, err
}

func DecodeToken(token string) (*ImClaims, error) {
	t, err := jwt.ParseWithClaims(token, &ImClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	claims := t.Claims.(*ImClaims)
	return claims, err
}
