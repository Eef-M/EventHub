package controllers

import (
	"net/http"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/gin-gonic/gin"
)

func GetMyProfile(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateMyProfile(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	var body struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	updates := models.User{
		Username:  body.Username,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}

	if err := initializers.DB.Model(&currentUser).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"data":    currentUser,
	})
}

func DeleteMyAccount(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	if err := initializers.DB.Delete(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete account",
		})
		return
	}

	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Account deleted successfully",
	})
}
