package routes

import (
	"github.com/JscorpTech/jst-go/api/controllers"
)

func (r *Route) InitRootRoute() {
	controller := controllers.NewRootController()
	router := r.App.Server.Group("")
	router.GET("", controller.RootController)
	router.GET("/health", controller.HealthController)
}
