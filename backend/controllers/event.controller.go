package controllers

import (
	"event-management-system/backend/handlers"
	"event-management-system/backend/initializers"
	"event-management-system/backend/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	file, err := c.FormFile("banner_image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filename, err := handlers.SaveImage(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
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

	dateStr := c.PostForm("date")
	parseDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format. use YYYY-MM-DD",
		})
		return
	}

	timeStr := c.PostForm("time")
	parseTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid time format. use HH:MM (24h)",
		})
		return
	}

	event := models.Event{
		OrganizerID: user.ID,
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Location:    c.PostForm("location"),
		Category:    c.PostForm("category"),
		Date:        parseDate,
		Time:        parseTime,
		BannerURL:   "uploads/" + filename,
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

	var event models.Event
	if err := initializers.DB.Where("id = ?", id).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	dateStr := c.PostForm("date")
	parseDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format. use YYYY-MM-DD",
		})
		return
	}

	timeStr := c.PostForm("time")
	parseTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid time format. use HH:MM (24h)",
		})
		return
	}

	event.Title = c.PostForm("title")
	event.Description = c.PostForm("description")
	event.Location = c.PostForm("location")
	event.Category = c.PostForm("category")
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

		filename, err := handlers.SaveImage(c, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		event.BannerURL = "uploads/" + filename
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

	var event models.Event
	if err := initializers.DB.Where("id = ?", id).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
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
