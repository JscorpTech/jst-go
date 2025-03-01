package routes

import "github.com/JscorpTech/jst-go/api/controllers"

func (r *Route) InitAuthRoute() {
	controller := controllers.NewAuthController(r.App)
	router := r.App.Server.Group("/auth")
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)
	router.POST("/change-password", controller.ChangePassword)
}
