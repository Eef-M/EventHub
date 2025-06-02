package repository

import (
	"github.com/Eef-M/EventHub/backend/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetMyTickets(db *gorm.DB, userID uuid.UUID) ([]dto.TicketDTO, error) {
	var tickets []dto.TicketDTO

	err := db.Table("event_registrations AS r").
		Select(`
            t.id AS ticket_id,
            e.id AS event_id,
            e.title,
            e.location,
            e.date,
            t.ticket_code,
            t.price
        `).
		Joins("JOIN events AS e ON r.event_id = e.id").
		Joins("JOIN tickets AS t ON t.event_id = e.id").
		Where("r.user_id = ?", userID).
		Scan(&tickets).Error

	return tickets, err
}
