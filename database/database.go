package database

import (
	"log"

	"github.com/Uallessonivo/go_card_manager/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&model.User{}, &model.Card{}, &model.Employee{})

	DB = Dbinstance{
		Db: db,
	}
}
