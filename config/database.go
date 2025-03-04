package config

import (
	"fmt"
	"log"
	"os"

	"library/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	// Auto migrate tables
	DB.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.Library{},
		&models.BookInventory{},
		&models.RequestEvent{},
		&models.IssueRegistery{},
	)

	fmt.Println("Database connected successfully.")
}
