package controllers

import (
	"net/http"

	"github.com/JscorpTech/jst-go/internal/application/usecase"
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

type AuthController struct {
	App             *bootstrap.App
	LoginUsecase    interfaces.LoginUsecasePort
	RegisterUsecase interfaces.RegisterUsecasePort
}

func NewAuthController(app *bootstrap.App) *AuthController {
	userRepository := repository.NewUserRepository(app.DB)
	return &AuthController{
		App:             app,
		LoginUsecase:    usecase.NewLoginUsecase(userRepository),
		RegisterUsecase: usecase.NewRegisterUsecase(userRepository),
	}
}

func (auth *AuthController) Login(c echo.Context) error {
	var payload dto.LoginRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateRequest(payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(messages.ValidationError, err))
	}

	user, err := auth.LoginUsecase.Login(payload.Phone, payload.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, utils.ErrorResponse(messages.PermissionDenied, nil))
	}
	token, err := auth.LoginUsecase.GetToken(user)
	if err != nil {
		return c.JSON(http.StatusForbidden, utils.ErrorResponse(messages.PermissionDenied, nil))
	}
	return c.JSON(http.StatusOK, token)
}

func (auth *AuthController) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}

func (auth *AuthController) Register(c echo.Context) error {
	var payload dto.RegisterRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(messages.BadRequest, nil))
	}

	if err := validator.ValidateRequest(payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(messages.ValidationError, err))
	}

	user, err := auth.RegisterUsecase.CreateUserIfNotExist(payload.Phone, payload.FirstName, payload.LastName, payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(messages.InternalError, dto.ValidationError{
			Field:   "phone",
			Type:    "unique",
			Message: messages.UserAlreadyExist,
		}))
	}

	token, err := auth.LoginUsecase.GetToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse(messages.InternalError, err))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("OK", dto.RegisterResponse{
		User: dto.User{
			ID:        user.ID,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Token: *token,
	}))
}

func (auth *AuthController) User(c echo.Context) error {
	user, _ := c.Get("user").(*models.User)
	return c.JSON(http.StatusOK, utils.SuccessResponse("OK", &dto.User{
		ID:        user.ID,
		Phone:     user.Phone,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}))
}

func (auth *AuthController) ResendCode(c echo.Context) error {
	return c.JSON(http.StatusOK, utils.SuccessResponse("OK", nil))
}

func (auth *AuthController) ChangePassword(c echo.Context) error {
	return c.JSON(http.StatusOK, utils.SuccessResponse("OK", nil))
}
