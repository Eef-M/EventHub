package main

import (
	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/routes"
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
