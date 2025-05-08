package seeders

import (
	"time"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/google/uuid"
)

func SeedEventRegistrations(users []models.User, events []models.Event, tickets []models.Ticket) {
	registrations := []models.EventRegistration{
		{
			ID:           uuid.New(),
			UserID:       users[2].ID,
			EventID:      events[0].ID,
			TicketID:     tickets[0].ID,
			Status:       "registered",
			RegisteredAt: time.Now(),
		},
		{
			ID:           uuid.New(),
			UserID:       users[3].ID,
			EventID:      events[1].ID,
			TicketID:     tickets[1].ID,
			Status:       "registered",
			RegisteredAt: time.Now(),
		},
		{
			ID:           uuid.New(),
			UserID:       users[3].ID,
			EventID:      events[2].ID,
			TicketID:     tickets[2].ID,
			Status:       "registered",
			RegisteredAt: time.Now(),
		},
		{
			ID:           uuid.New(),
			UserID:       users[2].ID,
			EventID:      events[3].ID,
			TicketID:     tickets[3].ID,
			Status:       "registered",
			RegisteredAt: time.Now(),
		},
		{
			ID:           uuid.New(),
			UserID:       users[2].ID,
			EventID:      events[4].ID,
			TicketID:     tickets[4].ID,
			Status:       "registered",
			RegisteredAt: time.Now(),
		},
		{
			ID:           uuid.New(),
			UserID:       users[3].ID,
			EventID:      events[5].ID,
			TicketID:     tickets[5].ID,
			Status:       "registered",
			RegisteredAt: time.Now(),
		},
		{
			ID:           uuid.New(),
			UserID:       users[1].ID,
			EventID:      events[1].ID,
			TicketID:     tickets[1].ID,
			Status:       "registered",
			RegisteredAt: time.Now(),
		},
		{
			ID:           uuid.New(),
			UserID:       users[0].ID,
			EventID:      events[4].ID,
			TicketID:     tickets[4].ID,
			Status:       "registered",
			RegisteredAt: time.Now(),
		},
	}

	for _, registration := range registrations {
		initializers.DB.Create(&registration)
	}
}
