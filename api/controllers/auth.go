package controllers

import (
	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/JscorpTech/jst-go/domain"
	auth2 "github.com/JscorpTech/jst-go/domain/auth"
	"github.com/JscorpTech/jst-go/pkg/messages"
	"github.com/JscorpTech/jst-go/pkg/validator"
	"github.com/JscorpTech/jst-go/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	App         *bootstrap.App
	AuthUsecase *usecase.AuthUsecase
}

func NewAuthController(app *bootstrap.App) *AuthController {
	return &AuthController{
		App: app,
		AuthUsecase: &usecase.AuthUsecase{
			DB: app.DB,
		},
	}
}

func (auth *AuthController) Login(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}
func (auth *AuthController) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}
func (auth *AuthController) Register(c echo.Context) error {
	var validatedData auth2.RegisterRequest
	if err := c.Bind(&validatedData); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.BadRequest, nil))
	}
	if err := validator.ValidateRequest(validatedData); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.ValidationError, err))
	}
	user, err := auth.AuthUsecase.CreateUserIfNotExist(validatedData.Phone, validatedData.FirstName, validatedData.LastName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse(messages.InternalError, err))
	}
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", auth2.RegisterResponse{
		User: auth2.User{
			Id:        user.ID,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Token: auth2.Token{
			AccessToken:  "111",
			RefreshToken: "222",
		},
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
