package routes

import (
	"event-management-system/backend/controllers"
	"event-management-system/backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	router := app

	api := router.Group("/api/v1")
	{
		// Auth
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)

		// Event
		api.GET("/events", middleware.RequireAuth, controllers.GetEvents)
		api.GET("/events/:id", middleware.RequireAuth, controllers.GetEvent)
		api.POST("/events", middleware.RequireAuth, controllers.CreateEvent)
		api.PUT("/events/:id", middleware.RequireAuth, controllers.UpdateEvent)
		api.DELETE("/events/:id", middleware.RequireAuth, controllers.DeleteEvent)

		// Valdate auth (for testing)
		api.GET("testing", middleware.RequireAuth, controllers.ValidateTesting)
	}
}
