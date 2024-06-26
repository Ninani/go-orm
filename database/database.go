package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Ninani/go-orm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Info),
		CreateBatchSize: 1000,
	})
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
		os.Exit(2)
	}
	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running db models migrations")
	db.AutoMigrate(&models.Fact{}, &models.User{}, &models.Quiz{})

	DB = Dbinstance{Db: db}
}
