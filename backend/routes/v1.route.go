package routes

import (
	"github.com/Eef-M/EventHub/backend/controllers"
	"github.com/Eef-M/EventHub/backend/middleware"
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
		auth.POST("/logout", middleware.RequireAuth, controllers.Logout)
	}

	// Events Group
	events := api.Group("/events")
	{
		events.GET("", controllers.GetEvents)
		events.GET("/:id", middleware.RequireAuth, controllers.GetEvent)
		events.POST("", middleware.RequireAuth, controllers.CreateEvent)
		events.PUT("/:id", middleware.RequireAuth, controllers.UpdateEvent)
		events.DELETE("/:id", middleware.RequireAuth, controllers.DeleteEvent)

		events.POST("/:id/register", middleware.RequireAuth, controllers.RegisterEvent)
		events.GET("/my-registrations", middleware.RequireAuth, controllers.MyRegistrations)
		events.GET("/:id/attendees", middleware.RequireAuth, controllers.EventAttendees)
		events.PUT("/:id/cancel", middleware.RequireAuth, controllers.CancelRegistration)

		events.POST("/:id/feedback", middleware.RequireAuth, controllers.SendFeedback)
		events.GET("/:id/feedback", middleware.RequireAuth, controllers.GetFeedbacks)
	}

	// Ticket Group
	ticket := api.Group("/tickets")
	{
		ticket.GET("", middleware.RequireAuth, controllers.GetTickets)
		ticket.GET("/:id", middleware.RequireAuth, controllers.GetTicket)
		ticket.POST("", middleware.RequireAuth, controllers.CreateTicket)
		ticket.PUT("/:id", middleware.RequireAuth, controllers.UpdateTicket)
		ticket.DELETE("/:id", middleware.RequireAuth, controllers.DeleteTicket)
	}
}
