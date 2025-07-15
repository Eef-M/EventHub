package initializers

import "github.com/Eef-M/EventHub/backend/models"

func SyncDB() {
	DB.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Ticket{},
		&models.EventRegistration{},
		&models.EventFeedback{},
		&models.Payment{},
	)
}
