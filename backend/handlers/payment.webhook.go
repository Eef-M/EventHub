package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Eef-M/EventHub/backend/models"
	"github.com/Eef-M/EventHub/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"
	"gorm.io/gorm"
)

func StripeWebhookHandler(db *gorm.DB, stripeService *services.StripeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const MaxBodyBytes = int64(65536)
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

		body, err := c.GetRawData()
		if err != nil {
			log.Printf("Error reading webhook body: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error reading request body",
			})
			return
		}

		sigHeader := c.GetHeader("Stripe-Signature")
		webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

		event, err := webhook.ConstructEvent(body, sigHeader, webhookSecret)
		if err != nil {
			log.Printf("Webhook signature verification failed: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Webhook signature verification failed",
			})
			return
		}

		log.Printf("Received Stripe webhook event: %s", event.Type)

		switch event.Type {
		case "payment_intent.succeeded":
			if err := handlePaymentSucceeded(event, stripeService); err != nil {
				log.Printf("Error handling payment success: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to process payment success",
				})
				return
			}

		case "payment_intent.payment_failed":
			if err := handlePaymentFailed(event, stripeService); err != nil {
				log.Printf("Error handling payment failure: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to process payment failure",
				})
				return
			}

		case "payment_intent.canceled":
			if err := handlePaymentCanceled(event, stripeService); err != nil {
				log.Printf("Error handling payment cancellation: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to process payment cancellation",
				})
				return
			}

		default:
			log.Printf("Unhandled event type: %s", event.Type)
		}

		c.Status(http.StatusOK)
	}
}

func handlePaymentSucceeded(event stripe.Event, stripeService *services.StripeService) error {
	var paymentIntent stripe.PaymentIntent
	if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
		return err
	}

	log.Printf("Payment succeeded: %s", paymentIntent.ID)

	if err := stripeService.UpdatePaymentStatus(paymentIntent.ID, models.PaymentStatusSucceeded); err != nil {
		return err
	}

	// Additional logic for successful payment (e.g., send confirmation email, update inventory)
	// TODO: Implement post-payment success logic

	return nil
}

func handlePaymentFailed(event stripe.Event, stripeService *services.StripeService) error {
	var paymentIntent stripe.PaymentIntent
	if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
		return err
	}

	log.Printf("Payment failed: %s", paymentIntent.ID)

	if err := stripeService.UpdatePaymentStatus(paymentIntent.ID, models.PaymentStatusFailed); err != nil {
		return err
	}

	// Additional logic for failed payment (e.g., notify user, release reserved tickets)
	// TODO: Implement post-payment failure logic

	return nil
}

func handlePaymentCanceled(event stripe.Event, stripeService *services.StripeService) error {
	var paymentIntent stripe.PaymentIntent
	if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
		return err
	}

	log.Printf("Payment canceled: %s", paymentIntent.ID)

	if err := stripeService.UpdatePaymentStatus(paymentIntent.ID, models.PaymentStatusCanceled); err != nil {
		return err
	}

	// Additional logic for canceled payment
	// TODO: Implement post-payment cancellation logic

	return nil
}
