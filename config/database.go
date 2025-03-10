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

func ConnectDBTest() {
	var err error
	DB_URL_TEST := "postgres://postgres:password@localhost:5432/test?sslmode=disable"
	dsn := DB_URL_TEST
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to test database!")
	}

	// defer DeleteAllData()
	// Auto migrate tables
	DB.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.Library{},
		&models.BookInventory{},
		&models.RequestEvent{},
		&models.IssueRegistery{},
	)

	fmt.Println("Test database connected successfully.")
}

// Delete all data from tables
func DeleteAllData() {
	DB.Exec("DELETE FROM users")
	DB.Exec("DELETE FROM sessions")
	DB.Exec("DELETE FROM libraries")
	DB.Exec("DELETE FROM book_inventories")
	DB.Exec("DELETE FROM request_events")
	DB.Exec("DELETE FROM issue_registries")
}
