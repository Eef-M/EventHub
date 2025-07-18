package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Eef-M/EventHub/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"
	"gorm.io/gorm"
)

func StripeWebhookhandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		const MaxBodyBytes = int64(65536)
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Read error: " + err.Error(),
			})
			return
		}

		sigHeader := c.GetHeader("Stripe-Signature")
		webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

		event, err := webhook.ConstructEvent(body, sigHeader, webhookSecret)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid signature | " + err.Error(),
			})
			return
		}

		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse data",
			})
			return
		}

		switch event.Type {
		case "payment_intent.succeeded":
			updatePaymentStatus(db, paymentIntent.ID, models.PaymentStatusSucceeded)

		case "payment_intent.payment_failed":
			updatePaymentStatus(db, paymentIntent.ID, models.PaymentStatusFailed)
		}

		c.Status(http.StatusOK)
	}
}

func updatePaymentStatus(db *gorm.DB, paymentIntentID string, status models.PaymentStatus) {
	db.Model(&models.Payment{}).
		Where("payment_intent_id = ?", paymentIntentID).
		Update("status", status)
}
