package repository

import (
	"github.com/Eef-M/EventHub/backend/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetMyRegistrations(db *gorm.DB, userID uuid.UUID) ([]dto.EventRegistrationDTO, error) {
	var results []dto.EventRegistrationDTO

	err := db.Table("event_registrations AS r").
		Select("r.id AS registration_id, "+
			"e.id AS event_id, "+
			"e.title, "+
			"e.description, "+
			"e.location, "+
			"e.date, "+
			"e.time, "+
			"e.banner_url, "+
			"r.status, "+
			"(SELECT COUNT(*) FROM event_registrations WHERE event_id = e.id) AS total_registrations").
		Joins("JOIN events AS e ON r.event_id = e.id").
		Where("r.user_id = ?", userID).
		Scan(&results).Error

	return results, err
}
