package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	jwtSecret  = []byte("6MNSobBrcOjBkO0fS6MNSobBChainO0fS")
	expireTime = 8 * time.Hour
)

type JwtClaims struct {
	Payload interface{} `json:"payload,omitempty"`
	jwt.StandardClaims
}

// User 测试结构体作为payload数据
type User struct {
	Name string `json:"name,omitempty"`
	Addr string `json:"addr,omitempty"`
}

func main() {
	if encode, err := Encode(User{"pqr", "abc"}); err == nil {
		fmt.Println(encode)
		decode, _ := Decode(encode)
		fmt.Println(decode.(map[string]interface{})["name"])
	}
}

func Encode(payload interface{}) (string, error) {
	claims := JwtClaims{
		payload,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
		},
	}

	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)

	return sign, err
}

func Decode(token string) (interface{}, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}

	claims, err := jwt.ParseWithClaims(
		token,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)

	if err != nil {
		return token, err
	}

	return claims.Claims.(*JwtClaims).Payload, err
}
