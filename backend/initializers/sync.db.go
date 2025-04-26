package initializers

import "event-management-system/backend/models"

func SyncDB() {
	DB.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Ticket{},
		&models.EventRegistration{},
	)
}
