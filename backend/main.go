package main

import (
	"fmt"
	"os"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/middleware"
	"github.com/Eef-M/EventHub/backend/routes"
	"github.com/Eef-M/EventHub/backend/seeders"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
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
	app.Use(middleware.CORSMiddleware())
	routes.InitRoute(app)
	app.Run()
}
