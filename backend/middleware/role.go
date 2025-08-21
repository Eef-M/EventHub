package middleware

import (
	"net/http"

	"github.com/Eef-M/EventHub/backend/models"
	"github.com/gin-gonic/gin"
)

func RoleAccess(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		currentUser := user.(models.User)
		if string(currentUser.Role) != role {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Forbidden – You do not have access",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
