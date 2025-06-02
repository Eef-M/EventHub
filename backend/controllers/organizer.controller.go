package controllers

import (
	"net/http"

	"github.com/Eef-M/EventHub/backend/dto"
	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/Eef-M/EventHub/backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func OrganizerDashboard(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	organizer := user.(models.User)

	totalEvents, totalRegistrations, registered, cancelledRegistrations, totalFeedbacks, err := repository.GetDashboardStatsByOrganizerID(initializers.DB, organizer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch stats. | " + err.Error(),
		})
		return
	}

	recentRegistrations, err := repository.GetRecentRegistrationsByOrganizerID(initializers.DB, organizer.ID, 5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch recent registrations. | " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_events":            totalEvents,
		"total_registrations":     totalRegistrations,
		"registered":              registered,
		"cancelled_registrations": cancelledRegistrations,
		"feedback_received":       totalFeedbacks,
		"recent_registrations":    recentRegistrations,
	})
}

func GetMyEventsHandler(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found in context",
		})
		return
	}

	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast user",
		})
		return
	}

	userID := user.ID
	uid, err := uuid.Parse(userID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID. | " + err.Error(),
		})
		return
	}

	events, err := repository.GetEventsByOrganizerID(initializers.DB, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get events. | " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

func GetAllRegistrationsHandler(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast user",
		})
		return
	}

	uid, err := uuid.Parse(user.ID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID. | " + err.Error(),
		})
		return
	}

	registrations, err := repository.GetAllRegistrationsByOrganizerID(initializers.DB, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch registrations. | " + err.Error(),
		})
		return
	}

	var result []dto.EventRegistrationsOrganizerDTO
	for _, r := range registrations {
		result = append(result, dto.EventRegistrationsOrganizerDTO{
			ID:           r.ID.String(),
			Username:     r.User.Username,
			Email:        r.User.Email,
			EventTitle:   r.Event.Title,
			TicketName:   r.Ticket.Name,
			Status:       r.Status,
			RegisteredAt: r.RegisteredAt.Format("2006-01-02 15:04"),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func GetAllFeedbackHandler(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast user",
		})
		return
	}

	uid, err := uuid.Parse(user.ID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID. | " + err.Error(),
		})
		return
	}

	feedbacks, err := repository.GetAllFeedbackByOrganizerID(initializers.DB, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch feedbacks" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": feedbacks,
	})
}

func GetAllTicketsHandler(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast user",
		})
		return
	}

	uid, err := uuid.Parse(user.ID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID. | " + err.Error(),
		})
		return
	}

	tickets, err := repository.GetAllTicketsByOrganizerID(initializers.DB, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch tickets. | " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}
