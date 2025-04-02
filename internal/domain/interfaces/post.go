package interfaces

import (
	"github.com/JscorpTech/jst-go/internal/domain/dto"
	"github.com/JscorpTech/jst-go/internal/models"
	"gorm.io/gorm"
)

type PostRepositoryPort interface {
	BaseRepositoryPort[models.PostModel]
}

type PostUsecasePort interface {
	Create(payload dto.PostCreateReuqest) (*dto.PostDetailResponse, error)
	FindByID(id uint) (*dto.PostDetailResponse, error)
	FindAll() ([]dto.PostListResponse, error)
	Update(entity *models.PostModel) error
	Delete(id uint) error
	Filter(model *models.PostModel, query string, params ...any) *gorm.DB
}
