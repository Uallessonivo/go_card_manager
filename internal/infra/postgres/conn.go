package postgres

import (
	entities2 "github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"log"
	"os"

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
		log.Fatal("Failed to connect to postgres")
	}

	log.Println("Connected to postgres")
	db.Logger = logger.Default.LogMode(logger.Info)

	if os.Getenv("AUTO_MIGRATE") == "true" {
		dbErr := db.AutoMigrate(&entities2.User{}, &entities2.Card{}, &entities2.Employee{})
		if dbErr != nil {
			return
		}
	}

	DB = Dbinstance{
		Db: db,
	}
}
