package seeders

import (
	"github.com/Eef-M/EventHub/backend/initializers"
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
	}

	for _, feedback := range feedbacks {
		initializers.DB.Create(&feedback)
	}
}
