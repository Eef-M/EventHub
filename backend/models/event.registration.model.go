package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Registered Status = "registered"
	Cancelled  Status = "cancelled"
)

type EventRegistration struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	EventID      uuid.UUID `gorm:"type:uuid;not null" json:"event_id"`
	TicketID     uuid.UUID `gorm:"type:uuid;not null" json:"ticket_id"`
	Status       Status    `gorm:"type:registration_status;default:'registered'" json:"status"`
	RegisteredAt time.Time `gorm:"autoCreateTime" json:"registered_at"`

	User   User   `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Event  Event  `gorm:"foreignKey:EventID;references:ID" json:"-"`
	Ticket Ticket `gorm:"foreignKey:TicketID;references:ID" json:"-"`
}

func (EventRegistration) TableName() string {
	return "event_registrations"
}

func (u *EventRegistration) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
