package repository

import (
	"github.com/Eef-M/EventHub/backend/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetMyTickets(db *gorm.DB, organizerID uuid.UUID) ([]dto.TicketDTO, error) {
	var tickets []dto.TicketDTO

	err := db.Table("tickets").
		Select(`
            tickets.id AS ticket_id,
            tickets.event_id,
            events.title,
            events.location,
            events.date,
            tickets.ticket_code,
            tickets.price
        `).
		Joins("JOIN events ON tickets.event_id = events.id").
		Where("events.organizer_id = ?", organizerID).
		Scan(&tickets).Error

	return tickets, err
}
