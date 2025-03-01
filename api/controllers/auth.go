package controllers

import (
	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	App *bootstrap.App
}

func NewAuthController(app *bootstrap.App) *AuthController {
	return &AuthController{
		App: app,
	}
}

func (auth *AuthController) Login(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}
func (auth *AuthController) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}
func (auth *AuthController) Register(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}
func (auth *AuthController) User(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}
func (auth *AuthController) ResendCode(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}
func (auth *AuthController) ChangePassword(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}
