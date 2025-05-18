package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/Eef-M/EventHub/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetEvents(c *gin.Context) {
	var events []models.Event

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	if err := initializers.DB.Limit(limit).Offset(offset).Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

func GetEvent(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	var event models.Event
	if err := initializers.DB.Where("id = ?", id).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": event,
	})
}

func CreateEvent(c *gin.Context) {
	title, ok := utils.RequireStringField(c, "title")
	if !ok {
		return
	}

	description, ok := utils.RequireStringField(c, "description")
	if !ok {
		return
	}

	location, ok := utils.RequireStringField(c, "location")
	if !ok {
		return
	}

	category, ok := utils.RequireStringField(c, "category")
	if !ok {
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

	dateStr := c.PostForm("date")
	parseDate, err := utils.ParseDate(dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format. use YYYY-MM-DD",
		})
		return
	}

	timeStr := c.PostForm("time")
	parseTime, err := utils.ParseTime(timeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid time format. use HH:MM (24h)",
		})
		return
	}

	file, err := c.FormFile("banner_image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filename, err := utils.SaveImage(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	event := models.Event{
		OrganizerID: user.ID,
		Title:       title,
		Description: description,
		Location:    location,
		Category:    category,
		Date:        parseDate,
		Time:        parseTime,
		BannerURL:   filename,
		IsPublic:    true,
	}

	if err := initializers.DB.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create data successfully",
		"data":    event,
	})
}

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := userInterface.(models.User)

	var event models.Event
	if err := initializers.DB.Where("id = ? AND organizer_id = ?", id, user.ID).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "You are not allowed to update this event / " + err.Error(),
		})
		return
	}

	title, ok := utils.RequireStringField(c, "title")
	if !ok {
		return
	}

	description, ok := utils.RequireStringField(c, "description")
	if !ok {
		return
	}

	location, ok := utils.RequireStringField(c, "location")
	if !ok {
		return
	}

	category, ok := utils.RequireStringField(c, "category")
	if !ok {
		return
	}

	dateStr := c.PostForm("date")
	parseDate, err := utils.ParseDate(dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format. use YYYY-MM-DD",
		})
		return
	}

	timeStr := c.PostForm("time")
	parseTime, err := utils.ParseTime(timeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid time format. use HH:MM (24h)",
		})
		return
	}

	event.Title = title
	event.Description = description
	event.Location = location
	event.Category = category
	event.Date = parseDate
	event.Time = parseTime

	file, err := c.FormFile("banner_image")
	if err == nil {
		if event.BannerURL != "" {
			oldFilePath := filepath.Join("uploads", filepath.Base(event.BannerURL))
			if err := os.Remove(oldFilePath); err != nil && !os.IsNotExist(err) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to delete old image",
				})
				return
			}
		}

		filename, err := utils.SaveImage(c, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		event.BannerURL = filename
	}

	if err := initializers.DB.Save(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update successfully",
		"data":    event,
	})
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
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
	if err := initializers.DB.Where("id = ? AND organizer_id = ?", id, user.ID).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "You are not allowed to delete this event / " + err.Error(),
		})
		return
	}

	if event.BannerURL != "" {
		imagePath := filepath.Join("uploads", filepath.Base(event.BannerURL))
		if err := os.Remove(imagePath); err != nil && !os.IsNotExist(err) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete image file",
			})
			return
		}
	}

	if err := initializers.DB.Delete(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete successfully",
	})
}
