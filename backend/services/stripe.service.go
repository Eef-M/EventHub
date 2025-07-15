package services

import (
	"context"
	"fmt"
	"os"

	"github.com/Eef-M/EventHub/backend/models"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
	"gorm.io/gorm"
)

type StripeService struct {
	db *gorm.DB
}

func NewStripeService(db *gorm.DB) *StripeService {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	return &StripeService{db: db}
}

func (s *StripeService) CreatePaymentIntent(ctx context.Context, userID, eventID uuid.UUID, amount int64, currency string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent: %w", err)
	}

	payment := &models.Payment{
		ID:              uuid.New(),
		UserID:          userID,
		EventID:         eventID,
		Amount:          amount,
		Currency:        currency,
		Status:          models.PaymentStatusPending,
		PaymentIntentID: pi.ID,
	}

	if err := s.db.WithContext(ctx).Create(payment).Error; err != nil {
		return nil, fmt.Errorf("failed to save payment: %w", err)
	}

	return pi, nil
}
