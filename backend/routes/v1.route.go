package routes

import (
	"event-management-system/backend/controllers"
	"event-management-system/backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	router := app
	api := router.Group("/api/v1")

	// Auth Group
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Events Group
	events := api.Group("/events")
	{
		events.GET("", middleware.RequireAuth, controllers.GetEvents)
		events.GET("/:id", middleware.RequireAuth, controllers.GetEvent)
		events.POST("", middleware.RequireAuth, controllers.CreateEvent)
		events.PUT("/:id", middleware.RequireAuth, controllers.UpdateEvent)
		events.DELETE("/:id", middleware.RequireAuth, controllers.DeleteEvent)

	}
}
