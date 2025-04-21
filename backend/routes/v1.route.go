package routes

import (
	"event-management-system/backend/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	router := app

	api := router.Group("/api/v1")
	{
		// Auth
		api.POST("auth/register", controllers.Register)
	}
}
