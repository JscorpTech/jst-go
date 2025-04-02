package repository

import (
	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	interfaces.BaseRepository[models.PostModel]
}

func NewPostRepository(db *gorm.DB) interfaces.PostRepositoryPort {
	return &PostRepository{
		BaseRepository: NewBaseRepository[models.PostModel](db),
	}
}
