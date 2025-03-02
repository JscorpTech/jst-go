package main

import (
	"fmt"
	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/JscorpTech/jst-go/models"
)

func main() {
	app := bootstrap.NewApp()
	defer bootstrap.CloseDatabaseConnection(app)
	err := app.DB.AutoMigrate(
		&models.UserModel{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Print("Migrations applied successfully!")
}
