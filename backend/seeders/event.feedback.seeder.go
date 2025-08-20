package seeders

import (
	"github.com/Eef-M/EventHub/backend/config"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/google/uuid"
)

func SeedEventFeedbacks(users []models.User, events []models.Event) {
	feedbacks := []models.EventFeedback{
		{
			ID:      uuid.New(),
			UserID:  users[2].ID,
			EventID: events[0].ID,
			Rating:  5,
			Comment: "Best event!",
		},
		{
			ID:      uuid.New(),
			UserID:  users[3].ID,
			EventID: events[1].ID,
			Rating:  4,
			Comment: "Good Event",
		},
		{
			ID:      uuid.New(),
			UserID:  users[3].ID,
			EventID: events[2].ID,
			Rating:  4,
			Comment: "Nice",
		},
		{
			ID:      uuid.New(),
			UserID:  users[2].ID,
			EventID: events[6].ID,
			Rating:  5,
			Comment: "Good!",
		},
		{
			ID:      uuid.New(),
			UserID:  users[3].ID,
			EventID: events[6].ID,
			Rating:  4,
			Comment: "Not Bad!",
		},
	}

	for _, feedback := range feedbacks {
		config.DB.Create(&feedback)
	}
}
