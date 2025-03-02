package bootstrap

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		env.DB_HOST,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_DATABASE,
		env.DB_PORT,
		env.TIMEZONE,
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}

func CloseDatabaseConnection(app *App) {
	sqlDB, err := app.DB.DB()
	var count int8
	count = 127
	fmt.Print(count)
	if err != nil {
		app.Server.Logger.Fatal("❌ DB obyektini olishda xatolik:", err)
	}
	if err := sqlDB.Close(); err != nil {
		app.Server.Logger.Fatal("Connection close error")
	}
}
