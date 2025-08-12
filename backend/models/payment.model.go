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
	PaymentStatusCanceled  PaymentStatus = "canceled"
)

type Payment struct {
	ID              uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TicketID        uuid.UUID     `gorm:"type:uuid;not null" json:"ticket_id"`
	UserID          uuid.UUID     `gorm:"type:uuid;not null" json:"user_id"`
	PaymentIntentID string        `gorm:"unique;not null;index" json:"payment_intent_id"`
	Amount          float64       `gorm:"not null" json:"amount"`
	Currency        string        `gorm:"type:varchar(10);default:'usd'" json:"currency"`
	Status          PaymentStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Quantity        int           `gorm:"default:1" json:"quantity" `
	CreatedAt       time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAT       time.Time     `gorm:"autoUpdateTime" json:"updated_at"`

	User   User   `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Ticket Ticket `gorm:"foreignKey:TicketID;references:ID" json:"-"`
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
