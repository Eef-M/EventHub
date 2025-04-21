package initializers

import "event-management-system/backend/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
