package services

import (
	"context"
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

func (s *StripeService) CreatePaymentIntent(ctx context.Context, userID, ticketID uuid.UUID, quantity int) (*stripe.PaymentIntent, error) {
	var ticket models.Ticket
	if err := s.db.First(&ticket, "id = ?", ticketID).Error; err != nil {
		return nil, err
	}

	totalPrice := ticket.Price * float64(quantity)
	amountInCents := int64(totalPrice * 100)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountInCents),
		Currency: stripe.String("usd"),
		Metadata: map[string]string{
			"ticket_id": ticketID.String(),
			"user_id":   userID.String(),
			"quantity":  string(rune(quantity)),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	payment := models.Payment{
		ID:              uuid.New(),
		TicketID:        ticketID,
		UserID:          userID,
		PaymentIntentID: pi.ID,
		Amount:          totalPrice,
		Currency:        "usd",
		Status:          models.PaymentStatusPending,
		Quantity:        quantity,
	}

	if err := s.db.Create(&payment).Error; err != nil {
		return nil, err
	}

	return pi, nil
}

func (s *StripeService) UpdatePaymentStatus(paymentIntentID string, status models.PaymentStatus) error {
	return s.db.Model(&models.Payment{}).
		Where("payment_intent_id = ?", paymentIntentID).
		Update("status", status).Error
}

func (s *StripeService) GetPaymentByIntentID(paymentIntentID string) (*models.Payment, error) {
	var payment models.Payment
	err := s.db.Where("payment_intent_id = ?", paymentIntentID).First(&payment).Error
	return &payment, err
}

func (s *StripeService) GetUserPayments(userID uuid.UUID, page, limit int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	offset := (page - 1) * limit

	if err := s.db.Model(&models.Payment{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := s.db.Where("user_id = ?", userID).
		Preload("Ticket").
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&payments).Error

	return payments, total, err
}

func (s *StripeService) GetPaymentByID(paymentID, userID uuid.UUID) (*models.Payment, error) {
	var payment models.Payment
	err := s.db.Where("id = ? AND user_id = ?", paymentID, userID).
		Preload("Ticket").
		Preload("User").
		First(&payment).Error
	return &payment, err
}

func (s *StripeService) GetAllPayments(page, limit int, status string) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	offset := (page - 1) * limit
	query := s.db.Model(&models.Payment{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("Ticket").
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&payments).Error

	return payments, total, err
}

func (s *StripeService) GetPaymentByIDOrganizer(paymentID uuid.UUID) (*models.Payment, error) {
	var payment models.Payment
	err := s.db.Where("id = ?", paymentID).
		Preload("Ticket").
		Preload("User").
		First(&payment).Error
	return &payment, err
}
