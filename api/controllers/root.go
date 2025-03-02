package controllers

import (
	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RootController struct {
	App *bootstrap.App
}

func NewRootController(app *bootstrap.App) *RootController {
	return &RootController{
		App: app,
	}
}
func (r *RootController) RootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
func (r *RootController) HealthHandler(c echo.Context) error {
	db, err := r.App.DB.DB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if err := db.Ping(); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "OK")
}
