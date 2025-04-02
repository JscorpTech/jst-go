package interfaces

import (
	"github.com/JscorpTech/jst-go/internal/models"
)

type PostRepositoryPort interface {
	BaseRepository[models.PostModel]
}

type PostUsecasePort interface {
	BaseUsecase[PostRepositoryPort, models.PostModel]
}
