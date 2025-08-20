package config

import (
	"fmt"
	"os"

	"github.com/Eef-M/EventHub/backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var BaseURL string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	BaseURL = fmt.Sprintf("http://localhost:%s", os.Getenv("PORT")) // must be replaced during production !!!
}

func ConnectDatabase() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not added to env!")
	}
	return port
}

func SyncDB() {
	DB.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Ticket{},
		&models.EventRegistration{},
		&models.EventFeedback{},
		&models.Payment{},
	)
}
