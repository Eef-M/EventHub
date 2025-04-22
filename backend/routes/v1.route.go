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
		api.POST("auth/register", controllers.Register)
		api.POST("auth/login", controllers.Login)

		// Valdate auth (for testing)
		api.GET("testing", middleware.RequireAuth, controllers.ValidateTesting)
	}
}
