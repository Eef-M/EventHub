package initializers

func SyncDB() {
	DB.AutoMigrate()
}
