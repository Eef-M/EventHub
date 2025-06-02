package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	EventID     uuid.UUID `gorm:"type:uuid;not null" json:"event_id"`
	Event       Event     `gorm:"foreignKey:EventID;references:ID" json:"-"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	Price       float64   `gorm:"type:decimal(10,2);not null;default:0.00" json:"price"`
	Quota       int       `gorm:"not null" json:"quota"`
	TicketCode  string    `gorm:"type:varchar(50);not null;unique" json:"ticket_code"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Ticket) TableName() string {
	return "tickets"
}

func (u *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
