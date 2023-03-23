package database

import (
	"log"
	"os"

	"github.com/nanorex07/cafe_app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

var all_models = []interface{}{
	&models.User{},
	&models.Menu{},
	&models.Item{},
}

func ConnectDB() {
	dsn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database.")
	log.Println("Running migrations")

	// execute if need to reset all tables
	// db.Migrator().DropTable(all_models...)

	db.AutoMigrate(all_models...)
	DB = DBInstance{
		Db: db,
	}
}
