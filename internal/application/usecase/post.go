package usecase

import (
	"github.com/JscorpTech/jst-go/internal/bootstrap"
	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/internal/models"
	"github.com/JscorpTech/jst-go/internal/repository"
)

type postUsecase struct {
	interfaces.BaseUsecasePort[interfaces.PostRepositoryPort, models.PostModel]
}

func NewPostUsecase(app *bootstrap.App) interfaces.PostUsecasePort {
	return &postUsecase{
		BaseUsecasePort: NewBaseUsecase(repository.NewPostRepository(app.DB)),
	}
}
