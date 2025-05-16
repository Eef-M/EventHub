package repository

import (
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetEventsByOrganizerID(db *gorm.DB, organizerID uuid.UUID) ([]models.Event, error) {
	var events []models.Event
	err := db.Where("organizer_id = ?", organizerID).Order("date ASC").Find(&events).Error
	return events, err
}
