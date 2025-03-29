package routes

import (
	"github.com/JscorpTech/jst-go/internal/api/controllers"
	"github.com/JscorpTech/jst-go/internal/api/middlewares"
	"github.com/JscorpTech/jst-go/internal/repository"
)

func (r *Route) InitAuthRoute() {
	controller := controllers.NewAuthController(r.App)
	router := r.App.Server.Group("/auth")
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)
	router.GET("/me", controller.User, middlewares.AuthMiddleware(repository.NewUserRepository(r.App.DB)))
	router.POST("/change-password", controller.ChangePassword)
}
