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

	// User Group
	user := api.Group("/user")
	{
		user.GET("/me", middleware.RequireAuth, controllers.GetMyProfile)
		user.PUT("/update", middleware.RequireAuth, controllers.UpdateMyProfile)
		user.DELETE("/delete", middleware.RequireAuth, controllers.DeleteMyAccount)
	}

	// Events Group
	events := api.Group("/events")
	{
		events.GET("", controllers.GetEvents)
		events.GET("/:id", controllers.GetEvent)
		events.POST("", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.CreateEvent)
		events.PUT("/:id", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.UpdateEvent)
		events.DELETE("/:id", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.DeleteEvent)

		events.POST("/:id/register", middleware.RequireAuth, controllers.RegisterEvent)
		events.GET("/my-registrations", middleware.RequireAuth, controllers.MyRegistrations)
		events.GET("/:id/attendees", middleware.RequireAuth, controllers.EventAttendees)
		events.PUT("/:id/cancel", middleware.RequireAuth, controllers.CancelRegistration)

		events.POST("/:id/feedbacks", middleware.RequireAuth, controllers.SendFeedback)
		events.GET("/:id/feedbacks", controllers.GetFeedbacks)

		events.PATCH("/:id/availability", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.IsPublicAndIsOpenHandler)
	}

	// Ticket Group
	ticket := api.Group("/tickets")
	{
		ticket.GET("", controllers.GetTickets)
		ticket.GET("/my-tickets", middleware.RequireAuth, controllers.GetMyTickets)
		ticket.GET("/:id", middleware.RequireAuth, controllers.GetTicket)
		ticket.POST("", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.CreateTicket)
		ticket.PUT("/:id", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.UpdateTicket)
		ticket.DELETE("/:id", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.DeleteTicket)
	}

	// Organizer Group
	organizer := api.Group("/organizer")
	{
		organizer.GET("/dashboard", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.OrganizerDashboard)
		organizer.GET("/events", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.GetMyEventsHandler)
		organizer.GET("/tickets", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.GetAllTicketsHandler)
		organizer.GET("/registrations", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.GetAllRegistrationsHandler)
		organizer.GET("/feedbacks", middleware.RequireAuth, middleware.RequireRole("organizer"), controllers.GetAllFeedbackHandler)
	}
}
