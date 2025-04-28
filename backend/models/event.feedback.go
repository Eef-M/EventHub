package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventFeedback struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	EventID   uuid.UUID `gorm:"type:uuid;not null" json:"event_id"`
	Rating    int       `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAT time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	User  User  `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Event Event `gorm:"foreignKey:EventID;references:ID" json:"-"`
}

func (EventFeedback) TableName() string {
	return "event_feedbacks"
}

func (u *EventFeedback) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
