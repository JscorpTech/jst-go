package bootstrap

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	Logger echo.Logger
	Env    *Env
	Server *echo.Echo
	DB     *gorm.DB
}

func NewApp() *App {
	e := echo.New()
	env := NewEnv()
	return &App{
		Logger: e.Logger,
		Env:    env,
		Server: e,
		DB:     NewDatabase(env),
	}
}
