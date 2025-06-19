package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OrganizerID uuid.UUID `gorm:"type:uuid;not null" json:"organizer_id"`
	Organizer   User      `gorm:"foreignKey:OrganizerID;references:ID" json:"-"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Location    string    `gorm:"type:varchar(255);not null" json:"location"`
	Category    string    `gorm:"type:varchar(255);not null" json:"category"`
	Date        time.Time `gorm:"type:date" json:"date"`
	Time        time.Time `gorm:"type:time" json:"time"`
	IsPublic    bool      `gorm:"default:true" json:"is_public"`
	IsOpen      bool      `gorm:"default:true" json:"is_open"`
	BannerURL   string    `gorm:"type:text" json:"banner_url"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAT   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Event) TableName() string {
	return "events"
}

func (u *Event) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
