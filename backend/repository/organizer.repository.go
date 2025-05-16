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

func GetDashboardStatsByOrganizerID(db *gorm.DB, organizerID uuid.UUID) (
	totalEvents, totalRegistrations, registered, cancelledRegistrations, feedbackReceived int64,
	err error,
) {
	err = db.Model(&models.Event{}).
		Where("organizer_id = ?", organizerID).
		Count(&totalEvents).Error
	if err != nil {
		return
	}

	err = db.Table("event_registrations").
		Joins("JOIN events ON event_registrations.event_id = events.id").
		Where("events.organizer_id = ?", organizerID).
		Count(&totalRegistrations).Error
	if err != nil {
		return
	}

	err = db.Table("event_registrations").
		Joins("JOIN events ON event_registrations.event_id = events.id").
		Where("events.organizer_id = ? AND event_registrations.status = ?", organizerID, "registered").
		Count(&registered).Error
	if err != nil {
		return
	}

	err = db.Table("event_registrations").
		Joins("JOIN events ON event_registrations.event_id = events.id").
		Where("events.organizer_id = ? AND event_registrations.status = ?", organizerID, "cancelled").
		Count(&cancelledRegistrations).Error
	if err != nil {
		return
	}

	err = db.Table("event_feedbacks").
		Joins("JOIN events ON event_feedbacks.event_id = events.id").
		Where("events.organizer_id = ?", organizerID).
		Count(&feedbackReceived).Error

	return
}

func GetRecentRegistrationsByOrganizerID(db *gorm.DB, organizerID uuid.UUID, limit int) ([]models.EventRegistration, error) {
	var regs []models.EventRegistration
	err := db.
		Joins("JOIN events ON event_registrations.event_id = events.id").
		Where("events.organizer_id = ?", organizerID).
		Preload("Event").
		Preload("User").
		Order("event_registrations.registered_at DESC").
		Limit(limit).
		Find(&regs).Error
	return regs, err
}
