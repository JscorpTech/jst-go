package bootstrap

import "github.com/labstack/echo/v4"

type App struct {
	Logger echo.Logger
	Env    *Env
	Server *echo.Echo
}

func NewApp() *App {
	e := echo.New()
	return &App{
		Logger: e.Logger,
		Env:    NewEnv(),
		Server: e,
	}
}
