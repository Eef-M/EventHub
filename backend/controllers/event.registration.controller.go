package controllers

import (
	"net/http"

	"github.com/Eef-M/EventHub/backend/config"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/Eef-M/EventHub/backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterEvent(c *gin.Context) {
	eventIDParam := c.Param("id")
	eventID, err := uuid.Parse(eventIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	var input struct {
		UserID   string `form:"user_id" binding:"required"`
		TicketID string `form:"ticket_id" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID and Ticket ID are required",
		})
		return
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	ticketID, err := uuid.Parse(input.TicketID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ticket ID",
		})
		return
	}

	registration := models.EventRegistration{
		UserID:   userID,
		EventID:  eventID,
		TicketID: ticketID,
		Status:   "registered",
	}

	if err := config.DB.Create(&registration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register event",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": registration,
	})
}

func MyRegistrations(c *gin.Context) {
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

	registrations, err := repository.GetMyRegistrations(config.DB, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch my registrations: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": registrations,
	})
}

func EventAttendees(c *gin.Context) {
	eventIDParam := c.Param("id")
	eventID, err := uuid.Parse(eventIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	var attendees []models.EventRegistration
	if err := config.DB.Where("event_id = ?", eventID).Find(&attendees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": attendees,
	})
}

func CancelRegistration(c *gin.Context) {
	ID := c.Param("id")
	registrationID, err := uuid.Parse(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event registration ID",
		})
		return
	}

	var registration models.EventRegistration
	if err := config.DB.First(&registration, "id = ?", registrationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	registration.Status = "cancelled"

	if err := config.DB.Save(&registration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": registration,
	})
}
