package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_URL")

	if dsn == "" {
		log.Fatal("DB_URL is missing")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get generic database object: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	DB = db

	log.Println("database connected")

}
