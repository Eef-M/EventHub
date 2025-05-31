package dto

type EventRegistrationDTO struct {
	RegistrationID     string `json:"registration_id"`
	EventID            string `json:"event_id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Time               string `json:"time"`
	BannerURL          string `json:"banner_url"`
	Status             string `json:"status"`
	TotalRegistrations int    `json:"total_registrations"`
}
