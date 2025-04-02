package controllers

import (
	"net/http"
	"strconv"

	"github.com/JscorpTech/jst-go/internal/bootstrap"
	"github.com/JscorpTech/jst-go/internal/domain/dto"
	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/internal/models"
	"github.com/JscorpTech/jst-go/internal/repository"
	"github.com/JscorpTech/jst-go/internal/shared/messages"
	"github.com/JscorpTech/jst-go/pkg/utils"
	"github.com/JscorpTech/jst-go/pkg/validator"
	"github.com/labstack/echo/v4"
)

type PostController struct {
	PostUsecase interfaces.PostUsecasePort
}

func NewPostController(app *bootstrap.App) *PostController {
	return &PostController{
		PostUsecase: repository.NewPostRepository(app.DB),
	}
}

func (p *PostController) List(c echo.Context) error {
	// List api view
	posts, err := p.PostUsecase.FindAll()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, utils.ErrorResponse(err.Error(), nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("OK", posts))
}

func (p *PostController) Detail(c echo.Context) error {
	// Detail api view one object
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 23)
	post, err := p.PostUsecase.FindByID(uint(postID))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Not found", nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("OK", post))
}

func (p *PostController) Create(c echo.Context) error {
	// Create api view
	var payload dto.PostCreateReuqest
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), nil))
	}
	if err := validator.ValidateRequest(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(messages.ValidationError, err))
	}
	post := models.PostModel{
		Title: payload.Title,
		Desc:  payload.Desc,
		Image: payload.Image,
	}
	p.PostUsecase.Create(&post)
	return c.JSON(http.StatusCreated, post)
}

func (p *PostController) Delete(c echo.Context) error {
	// Delete api view
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := p.PostUsecase.Delete(uint(postID)); err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Not found", nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("Deleted", nil))
}

func (p *PostController) Update(c echo.Context) error {
	// Update api view all fields required
	var payload dto.PostCreateReuqest
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), nil))
	}
	p.PostUsecase.Filter(&models.PostModel{}, "id = ?", uint(postID)).Updates(payload)
	return c.JSON(http.StatusOK, utils.SuccessResponse("Updated", nil))
}
