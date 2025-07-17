package utils

import (
	"net/http"

	"github.com/Eef-M/EventHub/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePaymentRequest struct {
	EventID  uuid.UUID `json:"event_id"`
	Amount   int64     `json:amount`
	Currency string    `json:"currency"`
}

func CreatePaymentHandler(stripeService *services.StripeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreatePaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalild request: " + err.Error(),
			})
			return
		}

		userIDVal, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		userID, ok := userIDVal.(uuid.UUID)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid user ID",
			})
			return
		}

		pi, err := stripeService.CreatePaymentIntent(c.Request.Context(), userID, req.EventID, req.Amount, req.Currency)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"client_secret": pi.ClientSecret,
		})
	}
}
