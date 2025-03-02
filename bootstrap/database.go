package bootstrap

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		env.DbHost,
		env.DbUser,
		env.DbPassword,
		env.DbDatabase,
		env.DbPort,
		env.Timezone,
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}

func CloseDatabaseConnection(app *App) {
	sqlDB, err := app.DB.DB()
	if err != nil {
		app.Server.Logger.Fatal("❌ DB obyektini olishda xatolik:", err)
	}
	if err := sqlDB.Close(); err != nil {
		app.Server.Logger.Fatal("Connection close error")
	}
}
