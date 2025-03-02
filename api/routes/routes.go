package routes

import (
	"github.com/JscorpTech/jst-go/bootstrap"
)

type Route struct {
	App *bootstrap.App
}

func NewRoute(app *bootstrap.App) *Route {
	return &Route{
		App: app,
	}
}

func InitRoutes(app *bootstrap.App) *Route {
	route := NewRoute(app)
	route.InitRootRoute()
	route.InitAuthRoute()
	return route
}
