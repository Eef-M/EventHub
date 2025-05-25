package repository

import (
	"github.com/Eef-M/EventHub/backend/dto"
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

func GetRecentRegistrationsByOrganizerID(db *gorm.DB, organizerID uuid.UUID, limit int) ([]dto.RecentRegistrationDTO, error) {
	var result []dto.RecentRegistrationDTO

	err := db.
		Table("event_registrations").
		Select("users.username AS username, events.title AS event_title, tickets.name AS ticket_name, event_registrations.status").
		Joins("JOIN events ON event_registrations.event_id = events.id").
		Joins("JOIN users ON event_registrations.user_id = users.id").
		Joins("JOIN tickets ON event_registrations.ticket_id = tickets.id").
		Where("events.organizer_id = ?", organizerID).
		Order("event_registrations.registered_at DESC").
		Limit(limit).
		Scan(&result).Error

	return result, err
}

func GetAllRegistrationsByOrganizerID(db *gorm.DB, organizerID uuid.UUID) ([]models.EventRegistration, error) {
	var registrations []models.EventRegistration

	err := db.Preload("User").
		Preload("Event").
		Preload("Ticket").
		Joins("JOIN events ON event_registrations.event_id = events.id").
		Where("events.organizer_id = ?", organizerID).
		Find(&registrations).Error

	return registrations, err
}

func GetAllFeedbackByOrganizerID(db *gorm.DB, organizerID uuid.UUID) ([]dto.EventFeedbackDTO, error) {
	var feedbacks []dto.EventFeedbackDTO

	err := db.Table("event_feedbacks").
		Select("event_feedbacks.comment, event_feedbacks.rating, event_feedbacks.created_at, users.username, events.title AS event_title").
		Joins("JOIN events ON event_feedbacks.event_id = events.id").
		Joins("JOIN users ON event_feedbacks.user_id = users.id").
		Where("events.organizer_id = ?", organizerID).
		Order("event_feedbacks.created_at DESC").
		Scan(&feedbacks).Error

	return feedbacks, err
}

func GetAllTicketsByOrganizerID(db *gorm.DB, organizerID uuid.UUID) ([]dto.TicketDTO, error) {
	var tickets []dto.TicketDTO

	err := db.Table("tickets").
		Select("tickets.id, tickets.name, tickets.price, tickets.quota, events.title AS event_title").
		Joins("JOIN events ON tickets.event_id = events.id").
		Where("events.organizer_id = ?", organizerID).
		Order("tickets.created_at DESC").
		Scan(&tickets).Error

	return tickets, err
}
