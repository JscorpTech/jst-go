package controllers

import (
	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/pkg/messages"
	"github.com/JscorpTech/jst-go/pkg/utils"
	"github.com/JscorpTech/jst-go/pkg/validator"
	"github.com/JscorpTech/jst-go/repository"
	"github.com/JscorpTech/jst-go/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	App             *bootstrap.App
	LoginUsecase    domain.LoginUsecase
	RegisterUsecase domain.RegisterUsecase
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
	var payload domain.LoginRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateRequest(payload); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.ValidationError, err))
	}

	user, err := auth.LoginUsecase.Login(payload.Username, payload.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, domain.ErrorResponse(messages.PermissionDenied, nil))
	}

	token, err := auth.LoginUsecase.GetToken(user)
	if err != nil {
		return c.JSON(http.StatusForbidden, domain.ErrorResponse(messages.PermissionDenied, nil))
	}

	return c.JSON(http.StatusOK, token)
}
func (auth *AuthController) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}
func (auth *AuthController) Register(c echo.Context) error {
	var payload domain.RegisterRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.BadRequest, nil))
	}

	if err := validator.ValidateRequest(payload); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.ValidationError, err))
	}

	user, err := auth.RegisterUsecase.CreateUserIfNotExist(payload.Phone, payload.FirstName, payload.LastName, payload.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse(messages.InternalError, domain.ValidationError{
			Field:   "phone",
			Type:    "unique",
			Message: messages.UserAlreadyExist,
		}))
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse(messages.InternalError, err))
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", domain.RegisterResponse{
		User: domain.User{
			Id:        user.ID,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Token: *token,
	}))
}
func (auth *AuthController) User(c echo.Context) error {
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", nil))
}
func (auth *AuthController) ResendCode(c echo.Context) error {
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", nil))
}
func (auth *AuthController) ChangePassword(c echo.Context) error {
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", nil))
}
