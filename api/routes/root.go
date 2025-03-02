package routes

import (
	"github.com/JscorpTech/jst-go/api/controllers"
)

func (r *Route) InitRootRoute() {
	controller := controllers.NewRootController(r.App)
	router := r.App.Server.Group("")
	router.GET("", controller.RootHandler)
	router.GET("/health", controller.HealthHandler)
}
