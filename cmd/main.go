package main

import (
	"fmt"
	"github.com/JscorpTech/jst-go/api/routes"
	"github.com/JscorpTech/jst-go/bootstrap"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	app := bootstrap.NewApp()
	defer bootstrap.CloseDatabaseConnection(app)
	routes.InitRoutes(app)
	go func() {
		fmt.Println("Starting profiler server at port 6060")
		log.Println(http.ListenAndServe("localhost:6060", nil)) // Profiling server
	}()
	app.Server.Logger.Fatal(app.Server.Start(fmt.Sprintf(":%s", app.Env.Port)))
}
