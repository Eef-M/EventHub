package dto

type RecentRegistrationDTO struct {
	Username   string `json:"username"`
	EventTitle string `json:"event_title"`
	TicketName string `json:"ticket_name"`
	Status     string `json:"status"`
}
