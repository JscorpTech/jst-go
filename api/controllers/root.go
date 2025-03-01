package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type RootController struct{}

func NewRootController() *RootController {
	return &RootController{}
}

func (r *RootController) RootController(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
func (r *RootController) HealthController(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
