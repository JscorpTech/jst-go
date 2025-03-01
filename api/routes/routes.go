package routes

import (
	"github.com/JscorpTech/jst-go/bootstrap"
)

type Route struct {
	App *bootstrap.App
}

func InitRoutes(app *bootstrap.App) *Route {
	route := &Route{
		App: app,
	}
	route.InitRootRoute()
	route.InitAuthRoute()
	return route
}
