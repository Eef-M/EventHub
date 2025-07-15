package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSucceeded PaymentStatus = "succeeded"
	PaymentStatusFailed    PaymentStatus = "failed"
)

type Payment struct {
	ID              uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID          uuid.UUID     `gorm:"type:uuid;not null" json:"user_id"`
	EventID         uuid.UUID     `gorm:"type:uuid;not null" json:"event_id"`
	Amount          int64         `gorm:"not null" json:"amount"`
	Currency        string        `gorm:"type:varchar(10);not null" json:"currency"`
	Status          PaymentStatus `gorm:"type:varchar(20);not null" json:"status"`
	PaymentIntentID string        `gorm:"type:varchar(255)" json:"payment_intent_id"`
	CreatedAt       time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAT       time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}

func (u *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
