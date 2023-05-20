package database

import (
	"log"
	"os"

	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	if os.Getenv("AUTO_MIGRATE") == "true" {
		dbErr := db.AutoMigrate(&entities.User{}, &entities.Card{}, &entities.Employee{})
		if dbErr != nil {
			return
		}
	}

	DB = Dbinstance{
		Db: db,
	}
}
