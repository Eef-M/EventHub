package controllers

import (
	"net/http"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/Eef-M/EventHub/backend/repository"
	"github.com/Eef-M/EventHub/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SendFeedback(c *gin.Context) {
	eventIDParam := c.Param("id")
	eventID, err := uuid.Parse(eventIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

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

	rating, ok := utils.RequireIntegerField(c, "rating")
	if !ok {
		return
	}

	comment, ok := utils.RequireStringField(c, "comment")
	if !ok {
		return
	}

	feedback := models.EventFeedback{
		UserID:  user.ID,
		EventID: eventID,
		Rating:  rating,
		Comment: comment,
	}

	if err := initializers.DB.Create(&feedback).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send feedback",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": feedback,
	})
}

func GetFeedbacks(c *gin.Context) {
	eventIDParam := c.Param("id")
	eventID, err := uuid.Parse(eventIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	feedbacks, err := repository.GetAllFeedback(initializers.DB, eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch feedbacks: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": feedbacks,
	})
}
