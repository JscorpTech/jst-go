package repository

import (
	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	interfaces.BaseRepositoryPort[models.PostModel]
}

func NewPostRepository(db *gorm.DB) interfaces.PostRepositoryPort {
	return &PostRepository{
		BaseRepositoryPort: NewBaseRepository[models.PostModel](db),
	}
}
