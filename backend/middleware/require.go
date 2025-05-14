package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil || tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: No token provided",
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Invalid claims",
		})
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Token expired",
		})
		return
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Invalid subject",
		})
		return
	}

	userID, err := uuid.Parse(sub)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Invalid UUID format",
		})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "id = ?", userID)
	if result.Error != nil || user.ID == uuid.Nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: User not found",
		})
		return
	}

	c.Set("user", user)
	c.Next()
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		currentUser := user.(models.User)
		if string(currentUser.Role) != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden – You do not have access"})
			c.Abort()
			return
		}

		c.Next()
	}
}
