package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var BaseURL string

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	BaseURL = os.Getenv("BASE_URL")
}
