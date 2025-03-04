package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"library/models"
)

var DB *gorm.DB

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// func ConnectDB() {
// 	var err error
// 	dsn := os.Getenv("DB_URL")
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal("Failed to connect to database!")
// 	}

// 	fmt.Println("Database connected successfully.")
// }

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	// Auto migrate tables
	DB.AutoMigrate(&models.User{}, &models.Session{})

	fmt.Println("Database connected successfully.")
}