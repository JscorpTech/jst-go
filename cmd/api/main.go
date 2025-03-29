package main

import (
	"fmt"

	"github.com/JscorpTech/jst-go/internal/api/routes"
	"github.com/JscorpTech/jst-go/internal/bootstrap"
)

func main() {
	app := bootstrap.NewApp()
	defer bootstrap.CloseApp(app)
	routes.InitRoutes(app)
	app.Server.Logger.Fatal(app.Server.Start(fmt.Sprintf(":%s", app.Env.Port)))
}
