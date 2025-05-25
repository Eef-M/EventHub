package controllers

import (
	"net/http"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/Eef-M/EventHub/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTickets(c *gin.Context) {
	var tickets []models.Ticket
	if err := initializers.DB.Find(&tickets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func GetTicket(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ticket ID",
		})
		return
	}

	var ticket models.Ticket
	if err := initializers.DB.Where("id = ?", id).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ticket,
	})
}

func CreateTicket(c *gin.Context) {
	eventIDStr := c.PostForm("event_id")
	eventID, err := uuid.Parse(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event_id UUID format",
		})
		return
	}

	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	user := userInterface.(models.User)

	var event models.Event
	if err := initializers.DB.First(&event, "id = ?", eventID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}
	if event.OrganizerID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not own this event",
		})
		return
	}

	name, ok := utils.RequireStringField(c, "name")
	if !ok {
		return
	}

	// Description is optional
	description := c.PostForm("description")

	price, ok := utils.RequireFloatField(c, "price")
	if !ok {
		return
	}

	quota, ok := utils.RequireIntegerField(c, "quota")
	if !ok {
		return
	}

	ticket := models.Ticket{
		EventID:     eventID,
		Name:        name,
		Description: description,
		Price:       price,
		Quota:       quota,
	}

	if err := initializers.DB.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket created successfully",
		"data":    ticket,
	})
}

func UpdateTicket(c *gin.Context) {
	ticketIDStr := c.Param("id")
	ticketID, err := uuid.Parse(ticketIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ticket ID",
		})
		return
	}

	var ticket models.Ticket
	if err := initializers.DB.First(&ticket, "id = ?", ticketID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	user := userInterface.(models.User)

	var event models.Event
	if err := initializers.DB.First(&event, "id = ?", ticket.EventID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}
	if event.OrganizerID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not own this ticket's event",
		})
		return
	}

	name, ok := utils.RequireStringField(c, "name")
	if !ok {
		return
	}

	description := c.PostForm("description")

	price, ok := utils.RequireFloatField(c, "price")
	if !ok {
		return
	}

	quota, ok := utils.RequireIntegerField(c, "quota")
	if !ok {
		return
	}

	ticket.Name = name
	ticket.Description = description
	ticket.Price = price
	ticket.Quota = quota

	if err := initializers.DB.Save(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket updated successfully",
		"data":    ticket,
	})
}

func DeleteTicket(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ticket ID",
		})
		return
	}

	var ticket models.Ticket
	if err := initializers.DB.First(&ticket, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	user := userInterface.(models.User)

	var event models.Event
	if err := initializers.DB.First(&event, "id = ?", ticket.EventID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}
	if event.OrganizerID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not own this ticket's event",
		})
		return
	}

	if err := initializers.DB.Delete(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete successfully",
	})
}
