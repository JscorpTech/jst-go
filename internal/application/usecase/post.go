package usecase

import (
	"github.com/JscorpTech/jst-go/internal/bootstrap"
	"github.com/JscorpTech/jst-go/internal/domain/dto"
	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/internal/models"
	"github.com/JscorpTech/jst-go/internal/repository"
)

type postUsecase struct {
	interfaces.BaseUsecasePort[interfaces.PostRepositoryPort, models.PostModel]
	PostRepository interfaces.PostRepositoryPort
}

func NewPostUsecase(app *bootstrap.App) interfaces.PostUsecasePort {
	return &postUsecase{
		BaseUsecasePort: NewBaseUsecase(repository.NewPostRepository(app.DB)),
		PostRepository:  repository.NewPostRepository(app.DB),
	}
}

func (p *postUsecase) Create(data any) (*models.PostModel, error) {
	payload, _ := data.(dto.PostCreateReuqest)
	post := models.PostModel{
		Title: payload.Title,
		Desc:  payload.Desc,
		Image: payload.Image,
	}
	if err := p.PostRepository.Create(&post); err != nil {
		return nil, err
	}
	return &post, nil
}
