package main

import (
	"fmt"
	"os"

	"github.com/Eef-M/EventHub/backend/config"
	"github.com/Eef-M/EventHub/backend/middleware"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/Eef-M/EventHub/backend/routes"
	"github.com/Eef-M/EventHub/backend/seeders"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDatabase()
	config.SyncDB()

	// Auto-run seeder after migration if enabled and database is empty
	autoSeed()
	config.ConnectRedis()
}

func main() {
	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "seed":
		runSeeder()
	default:
		runServer()
	}
}

// autoSeed runs the seeder automatically after migration
func autoSeed() {
	// Check if AUTO_SEED environment variable is set (default: true)
	autoSeedEnabled := os.Getenv("AUTO_SEED")
	if autoSeedEnabled == "" {
		autoSeedEnabled = "true" // default to true
	}

	if autoSeedEnabled != "true" {
		fmt.Println("Auto-seeding is disabled (set AUTO_SEED=true to enable)")
		return
	}

	// Check if database is empty (no users exist)
	var count int64
	config.DB.Model(&models.User{}).Count(&count)

	if count > 0 {
		fmt.Println("Database already contains data, skipping auto-seed")
		return
	}

	fmt.Println("\n=== Running database seeder automatically ===")
	runSeeder()
	fmt.Println("=== Auto-seeding completed! ===")
}

func runSeeder() {
	fmt.Println("Running database seeder...")
	users := seeders.SeedUsers()
	events := seeders.SeedEvents(users)
	tickets := seeders.SeedTickets(events)
	seeders.SeedEventRegistrations(users, events, tickets)
	seeders.SeedEventFeedbacks(users, events)
	fmt.Println("Seeding completed!")
}

func runServer() {
	app := gin.Default()
	app.Static("/uploads", "./uploads")
	app.Use(middleware.CORSMiddleware())
	routes.InitRoute(app)
	app.Run()
}
