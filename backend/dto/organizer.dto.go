package dto

type RecentRegistrationDTO struct {
	Username   string `json:"username"`
	EventTitle string `json:"event_title"`
	TicketName string `json:"ticket_name"`
	Status     string `json:"status"`
}

type EventRegistrationsDTO struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	EventTitle   string `json:"event_title"`
	TicketName   string `json:"ticket_name"`
	Status       string `json:"status"`
	RegisteredAt string `json:"registered_at"`
}

type EventFeedbackDTO struct {
	Username   string `json:"username"`
	EventTitle string `json:"event_title"`
	Comment    string `json:"comment"`
	Rating     int    `json:"rating"`
	CreatedAt  string `json:"created_at"`
}

type TicketDTO struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quota      int     `json:"quota"`
	EventTitle string  `json:"event_title"`
}
