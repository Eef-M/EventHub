package main

import (
	"event-management-system/backend/initializers"
	"event-management-system/backend/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	app := gin.Default()

	routes.InitRoute(app)
	app.Run()
}
