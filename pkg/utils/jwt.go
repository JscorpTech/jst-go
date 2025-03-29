package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type Jwt struct {
	Sub  uint   `json:"sub"`
	Type string `json:"type"`
	Exp  int    `json:"exp"`
	Iat  int    `json:"iat"`
}

func GenerateJwt(data *Jwt) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"iat":  time.Now().Unix(),
		"type": data.Type,
		"sub":  data.Sub,
	})
	return token.SignedString([]byte(viper.GetString("KEY")))
}

func ParseJwt(tokenString string) (*Jwt, error) {
	secretKey := []byte(viper.GetString("KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("tokenni tekshirib bo‘lmadi: %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, _ := claims["sub"].(float64)
		exp, _ := claims["exp"].(float64)
		iat, _ := claims["iat"].(float64)
		tokenType, _ := claims["type"].(string)
		if tokenType != "access" {
			return nil, errors.New("this token not access")
		}
		return &Jwt{
			Sub:  uint(sub),
			Exp:  int(exp),
			Iat:  int(iat),
			Type: tokenType,
		}, nil
	} else {
		return nil, fmt.Errorf("token noto‘g‘ri yoki muddati tugagan")
	}
}
