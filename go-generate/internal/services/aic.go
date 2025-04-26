package services

import (
	"github.com/JscorpTech/jst-go/go-generate/internal/domain"
	"github.com/JscorpTech/jst-go/go-generate/internal/providers"
)

type aicService struct{}

func NewAicService() domain.AicServicePort {
	return &aicService{}
}

func (a *aicService) Generate(prompt string) string {
	provider := providers.NewPizzaGptProvider()
	return provider.Generate(prompt)
}
