package routes

import (
	"github.com/Eef-M/EventHub/backend/config"
	"github.com/Eef-M/EventHub/backend/controllers"
	"github.com/Eef-M/EventHub/backend/handlers"
	"github.com/Eef-M/EventHub/backend/middleware"
	"github.com/Eef-M/EventHub/backend/services"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	router := app
	api := router.Group("/api/v1")

	// Initialize services
	stripeService := services.NewStripeService(config.DB)

	// Auth Group
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)
	}

	// User Group
	user := api.Group("/user")
	{
		user.GET("/me", middleware.AuthMiddleware(), controllers.GetMyProfile)
		user.PUT("/update", middleware.AuthMiddleware(), controllers.UpdateMyProfile)
		user.DELETE("/delete", middleware.AuthMiddleware(), controllers.DeleteMyAccount)
	}

	// Events Group
	events := api.Group("/events")
	{
		events.GET("", controllers.GetEvents)
		events.GET("/:id", controllers.GetEvent)
		events.POST("", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.CreateEvent)
		events.PUT("/:id", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.UpdateEvent)
		events.DELETE("/:id", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.DeleteEvent)

		events.POST("/:id/register", middleware.AuthMiddleware(), controllers.RegisterEvent)
		events.GET("/my-registrations", middleware.AuthMiddleware(), controllers.MyRegistrations)
		events.GET("/:id/attendees", middleware.AuthMiddleware(), controllers.EventAttendees)
		events.PUT("/:id/cancel", middleware.AuthMiddleware(), controllers.CancelRegistration)

		events.POST("/:id/feedbacks", middleware.AuthMiddleware(), controllers.SendFeedback)
		events.GET("/:id/feedbacks", controllers.GetFeedbacks)

		events.PATCH("/:id/availability", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.IsPublicAndIsOpenHandler)
	}

	// Ticket Group
	ticket := api.Group("/tickets")
	{
		ticket.GET("", controllers.GetTickets)
		ticket.GET("/my-tickets", middleware.AuthMiddleware(), controllers.GetMyTickets)
		ticket.GET("/:id", middleware.AuthMiddleware(), controllers.GetTicket)
		ticket.POST("", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.CreateTicket)
		ticket.PUT("/:id", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.UpdateTicket)
		ticket.DELETE("/:id", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.DeleteTicket)
	}

	// Organizer Group
	organizer := api.Group("/organizer")
	{
		organizer.GET("/dashboard", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.OrganizerDashboard)
		organizer.GET("/events", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.GetMyEventsHandler)
		organizer.GET("/tickets", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.GetAllTicketsHandler)
		organizer.GET("/registrations", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.GetAllRegistrationsHandler)
		organizer.GET("/feedbacks", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), controllers.GetAllFeedbackHandler)
	}

	// Payment Group
	payment := api.Group("/payments")
	{
		payment.POST("/create", middleware.AuthMiddleware(), handlers.CreatePaymentHandler(stripeService))
		payment.GET("/history", middleware.AuthMiddleware(), handlers.GetPaymentHistoryHandler(stripeService))
		payment.GET("/:id", middleware.AuthMiddleware(), handlers.GetPaymentHandler(stripeService))
		payment.POST("/webhook", handlers.StripeWebhookHandler(config.DB, stripeService))
		payment.GET("/organizer/all", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), handlers.GetAllPaymentsHandler(stripeService))
		payment.GET("/organizer/:id", middleware.AuthMiddleware(), middleware.RoleAccess("organizer"), handlers.GetPaymentDetailsHandler(stripeService))
	}
}
