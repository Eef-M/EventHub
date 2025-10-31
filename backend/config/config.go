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
	// Load .env file, but don't panic if not found (for Docker environment)
	err := godotenv.Load()
	if err != nil {
		// In Docker, environment variables are set via docker-compose
		// So we just log the warning instead of panic
		fmt.Println("Warning: .env file not found, using environment variables from system")
	}

	// Set BaseURL - get from environment or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	// For Docker, use the service name or environment variable
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL != "" {
		BaseURL = backendURL
	} else {
		BaseURL = fmt.Sprintf("http://localhost:%s", port)
	}
}

func ConnectDatabase() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		getEnvOrDefault("DB_HOST", "localhost"),
		getEnvOrDefault("DB_USER", "postgres"),
		getEnvOrDefault("DB_PASSWORD", "postgres"),
		getEnvOrDefault("DB_NAME", "eventhub_db"),
		getEnvOrDefault("DB_PORT", "5432"),
		getEnvOrDefault("DB_SSLMODE", "disable"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	fmt.Println("Database connected successfully!")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		// Return default port instead of panic for better Docker compatibility
		fmt.Println("Warning: PORT not set in environment, using default 8080")
		return "8080"
	}
	return port
}

func SyncDB() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Ticket{},
		&models.EventRegistration{},
		&models.EventFeedback{},
		&models.Payment{},
	)

	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	fmt.Println("Database tables synced successfully!")

	// Auto-run seeder after migration
	autoSeed()
}

// autoSeed runs the seeder automatically after migration
func autoSeed() {
	// Check if AUTO_SEED is enabled (default: true)
	autoSeedEnabled := getEnvOrDefault("AUTO_SEED", "true")

	if autoSeedEnabled != "true" {
		fmt.Println("Auto-seeding is disabled")
		return
	}

	// Check if database is empty (no users exist)
	var count int64
	DB.Model(&models.User{}).Count(&count)

	if count > 0 {
		fmt.Println("Database already contains data, skipping auto-seed")
		return
	}

	fmt.Println("Running database seeder automatically...")
	// Import seeders package here to avoid circular dependency
	// You'll need to call the seeder functions from main package instead
	fmt.Println("⚠️  Please run seeders manually from main package to avoid circular import")
}

// Helper function to get environment variable with default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
