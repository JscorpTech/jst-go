package routes

import (
	"github.com/JscorpTech/jst-go/internal/api/controllers"
)

func (r *Route) InitPostRoute() {
	controller := controllers.NewPostController(r.App)
	router := r.App.Server.Group("/post")
	router.GET("/", controller.List)
	router.POST("/", controller.Create)
	router.GET("/:id/", controller.Detail)
	router.DELETE("/:id/", controller.Delete)
	router.PUT("/:id/", controller.Update)
	router.PATCH("/:id/", controller.Update)
}
