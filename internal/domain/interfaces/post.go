package interfaces

import (
	"github.com/JscorpTech/jst-go/internal/models"
)

type PostRepositoryPort interface {
	BaseRepositoryPort[models.PostModel]
}

type PostUsecasePort interface {
	BaseUsecasePort[PostRepositoryPort, models.PostModel]
}
