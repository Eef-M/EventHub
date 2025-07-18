package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Eef-M/EventHub/backend/handlers"
	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/Eef-M/EventHub/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetMyProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found in context",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateMyProfile(c *gin.Context) {
	username, ok := utils.RequireStringField(c, "username")
	if !ok {
		return
	}

	firstName, ok := utils.RequireStringField(c, "first_name")
	if !ok {
		return
	}

	lastName, ok := utils.RequireStringField(c, "last_name")
	if !ok {
		return
	}

	email, ok := utils.RequireStringField(c, "email")
	if !ok {
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found in context",
		})
		return
	}

	currentUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast user",
		})
		return
	}

	updates := models.User{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	if file, err := c.FormFile("avatar_url"); err == nil {
		newAvatarURL, err := handlers.SaveAvatarImage(c, file, currentUser.AvatarURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to upload avatar",
			})
			return
		}
		updates.AvatarURL = newAvatarURL
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
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found in context",
		})
		return
	}

	currentUser := user.(models.User)

	if strings.Contains(currentUser.AvatarURL, "/uploads/avatars/") {
		parts := strings.Split(currentUser.AvatarURL, "/uploads/avatars/")
		if len(parts) == 2 {
			filename := parts[1]
			if filename != "default_avatar.png" {
				avatarPath := filepath.Join("uploads", "avatars", filename)
				_ = os.Remove(avatarPath)
			}
		}
	}

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
