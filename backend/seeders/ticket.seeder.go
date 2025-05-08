package seeders

import (
	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/google/uuid"
)

func SeedTickets(events []models.Event) []models.Ticket {
	tickets := []models.Ticket{
		{
			ID:          uuid.New(),
			EventID:     events[0].ID,
			Name:        "Early Bird",
			Description: "Ticket for early bird.",
			Price:       25,
			Quota:       100,
		},
		{
			ID:          uuid.New(),
			EventID:     events[1].ID,
			Name:        "Regular",
			Description: "Ticket for reguler.",
			Price:       40,
			Quota:       200,
		},
		{
			ID:          uuid.New(),
			EventID:     events[2].ID,
			Name:        "VIP",
			Description: "Ticket for VIP.",
			Price:       145,
			Quota:       50,
		},
		{
			ID:          uuid.New(),
			EventID:     events[3].ID,
			Name:        "Regular",
			Description: "Ticket for reguler.",
			Price:       45,
			Quota:       230,
		},
		{
			ID:          uuid.New(),
			EventID:     events[4].ID,
			Name:        "Early Bird",
			Description: "Ticket for early bird.",
			Price:       15,
			Quota:       125,
		},
		{
			ID:          uuid.New(),
			EventID:     events[5].ID,
			Name:        "VIP",
			Description: "Tikcet for VIP.",
			Price:       95,
			Quota:       100,
		},
	}

	for _, ticket := range tickets {
		initializers.DB.Create(&ticket)
	}

	return tickets
}
