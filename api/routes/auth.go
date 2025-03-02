package routes

import (
	"github.com/JscorpTech/jst-go/api/controllers"
	"github.com/JscorpTech/jst-go/api/middlewares"
	"github.com/JscorpTech/jst-go/repository"
)

func (r *Route) InitAuthRoute() {
	controller := controllers.NewAuthController(r.App)
	router := r.App.Server.Group("/auth")
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)
	router.GET("/me", controller.User, middlewares.AuthMiddleware(repository.NewUserRepository(r.App.DB)))
	router.POST("/change-password", controller.ChangePassword)
}
