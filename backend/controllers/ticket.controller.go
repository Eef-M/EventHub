package controllers

import (
	"event-management-system/backend/initializers"
	"event-management-system/backend/models"
	"net/http"
	"strconv"

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

	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}

	// Description is optional
	description := c.PostForm("description")

	priceStr := c.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid price",
		})
		return
	}

	quotaStr := c.PostForm("quota")
	quota, err := strconv.Atoi(quotaStr)
	if err != nil || quota < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid quota",
		})
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

	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}

	description := c.PostForm("description")

	priceStr := c.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid price",
		})
		return
	}

	quotaStr := c.PostForm("quota")
	quota, err := strconv.Atoi(quotaStr)
	if err != nil || quota < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid quota",
		})
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
