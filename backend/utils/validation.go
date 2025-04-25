package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RequireStringField(c *gin.Context, field string) (string, bool) {
	value := c.PostForm(field)
	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": field + " is required",
		})
		return "", false
	}
	return value, true
}

func RequireFloatField(c *gin.Context, field string) (float64, bool) {
	value := c.PostForm(field)
	num, err := strconv.ParseFloat(value, 64)
	if err != nil || num < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid " + field,
		})
		return 0, false
	}
	return num, true
}

func RequireIntegerField(c *gin.Context, field string) (int, bool) {
	value := c.PostForm(field)
	num, err := strconv.Atoi(value)
	if err != nil || num < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid " + field,
		})
		return 0, false
	}
	return num, true
}
