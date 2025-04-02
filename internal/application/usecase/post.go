package usecase

import (
	"github.com/JscorpTech/jst-go/internal/bootstrap"
	"github.com/JscorpTech/jst-go/internal/domain/dto"
	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/internal/models"
	"github.com/JscorpTech/jst-go/internal/repository"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type postUsecase struct {
	PostRepository interfaces.PostRepositoryPort
}

func NewPostUsecase(app *bootstrap.App) interfaces.PostUsecasePort {
	return &postUsecase{
		PostRepository: repository.NewPostRepository(app.DB),
	}
}

func (u *postUsecase) FindByID(id uint) (*dto.PostDetailResponse, error) {
	var response dto.PostDetailResponse
	post, err := u.PostRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	copier.Copy(&response, &post)
	return &response, nil
}

func (u *postUsecase) FindAll() ([]dto.PostListResponse, error) {
	var response []dto.PostListResponse
	posts, err := u.PostRepository.FindAll()
	if err != nil {
		return nil, err
	}
	copier.Copy(&response, &posts)
	return response, nil
}

func (u *postUsecase) Update(entity *models.PostModel) error {
	return u.PostRepository.Update(entity)
}

func (u *postUsecase) Delete(id uint) error {
	return u.PostRepository.Delete(id)
}

func (u *postUsecase) Filter(model *models.PostModel, query string, params ...any) *gorm.DB {
	return u.PostRepository.Filter(model, query, params...)
}

func (p *postUsecase) Create(payload dto.PostCreateReuqest) (*dto.PostDetailResponse, error) {
	var response dto.PostDetailResponse
	post := models.PostModel{
		Title: payload.Title,
		Desc:  payload.Desc,
		Image: payload.Image,
	}
	if err := p.PostRepository.Create(&post); err != nil {
		return nil, err
	}
	copier.Copy(&response, &post)
	return &response, nil
}
