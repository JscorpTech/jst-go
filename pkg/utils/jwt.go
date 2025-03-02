package utils

import (
	"encoding/json"
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/models"
)

type Jwt struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func GenerateToken(user *models.UserModel) (*domain.Token, error) {
	accessToken, err := GenerateJwt(&Jwt{})
	if err != nil {
		return nil, err
	}
	refreshToken, err := GenerateJwt(&Jwt{})
	if err != nil {
		return nil, err
	}
	return &domain.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GenerateJwt(data *Jwt) (string, error) {
	res, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func ParseJwt(tokenString []byte) (*Jwt, error) {
	var res Jwt
	if err := json.Unmarshal(tokenString, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
