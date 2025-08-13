package handlers

import (
	"net/http"
	"strconv"

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

func GetPaymentHistoryHandler(stripeService *services.StripeService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid user ID",
				})
				return
			}
			userID = parsed
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user ID",
			})
			return
		}

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

		payments, total, err := stripeService.GetUserPayments(userID, page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"payments": payments,
			"total":    total,
			"page":     page,
			"limit":    limit,
		})
	}
}

func GetPaymentHandler(stripeService *services.StripeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		paymentID := c.Param("id")
		if paymentID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Payment ID is required",
			})
			return
		}

		paymentUUID, err := uuid.Parse(paymentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid payment ID",
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
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid user ID",
				})
				return
			}
			userID = parsed
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user ID",
			})
			return
		}

		payment, err := stripeService.GetPaymentByID(paymentUUID, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Payment not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"payment": payment,
		})
	}
}

func GetAllPaymentsHandler(stripeService *services.StripeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
		status := c.Query("status")

		payments, total, err := stripeService.GetAllPayments(page, limit, status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"payments": payments,
			"total":    total,
			"page":     page,
			"limit":    limit,
		})
	}
}

func GetPaymentDetailsHandler(stripeService *services.StripeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		paymentID := c.Param("id")
		if paymentID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Payment ID is required",
			})
			return
		}

		paymentUUID, err := uuid.Parse(paymentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid payment ID",
			})
			return
		}

		payment, err := stripeService.GetPaymentByIDOrganizer(paymentUUID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Payment not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"payment": payment,
		})
	}
}
