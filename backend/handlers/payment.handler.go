package handlers

import (
	"net/http"

	"github.com/Eef-M/EventHub/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePaymentRequest struct {
	TicketID uuid.UUID `json:"ticket_id" binding:"required"`
	Quantity int       `json:"quantity"`
}

type CreatePaymentResponse struct {
	ClientSecret    string  `json:"client_secret"`
	PaymentIntentID string  `json:"payment_intent_id"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
}

func CreatePaymentHandler(stripeService *services.StripeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreatePaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request: " + err.Error(),
			})
			return
		}

		if req.Quantity <= 0 {
			req.Quantity = 1
		}

		if req.Quantity > 10 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Quantity cannot exceed 10",
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

		var userID uuid.UUID
		switch v := userIDVal.(type) {
		case uuid.UUID:
			userID = v
		case string:
			parsed, err := uuid.Parse(v)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Invalid user ID format",
				})
				return
			}
			userID = parsed
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid user ID type",
			})
			return
		}

		pi, err := stripeService.CreatePaymentIntent(c.Request.Context(), userID, req.TicketID, req.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create payment: " + err.Error(),
			})
			return
		}

		response := CreatePaymentResponse{
			ClientSecret:    pi.ClientSecret,
			PaymentIntentID: pi.ID,
			Amount:          float64(pi.Amount) / 100,
			Currency:        string(pi.Currency),
		}

		c.JSON(http.StatusOK, response)
	}
}
