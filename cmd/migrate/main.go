package main

import (
	"fmt"
	"log"

	"github.com/JscorpTech/jst-go/internal/bootstrap"
	"github.com/JscorpTech/jst-go/internal/models"
	"gorm.io/gorm"
)

var modelsToMigrate = []interface{}{
	&models.User{},
	&models.Token{},
	&models.PostModel{},
}

func getTableName(db *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(model)
	return stmt.Schema.Table
}

func migrateTablesIndividually(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		tableName := getTableName(db, model)

		if db.Migrator().HasTable(model) {
			var confirm string
			fmt.Printf("'%s' jadvali mavjud! Uni o‘chirib qayta yaratishni xohlaysizmi? (yes/no): ", tableName)
			_, _ = fmt.Scanln(&confirm)

			if confirm == "yes" {
				err := db.Migrator().DropTable(model)
				if err != nil {
					log.Fatalf("'%s' jadvalini o‘chirishda xatolik: %v", tableName, err)
				}
				fmt.Printf("'%s' jadvali o‘chirildi!\n", tableName)
			} else {
				fmt.Printf("'%s' jadvali o‘chirilmasdan qoldirildi.\n", tableName)
				continue
			}
		}

		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("'%s' jadvalini migratsiya qilishda xatolik: %v", tableName, err)
		}
		fmt.Printf("'%s' jadvali muvaffaqiyatli yaratildi!\n", tableName)
	}
}

func main() {
	app := bootstrap.NewApp()
	defer bootstrap.CloseDatabaseConnection(app)

	migrateTablesIndividually(app.DB, modelsToMigrate)
}
