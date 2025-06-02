package dto

type TicketDTO struct {
	TicketID   string  `json:"ticket_id"`
	EventID    string  `json:"event_id"`
	Title      string  `json:"title"`
	Location   string  `json:"location"`
	Date       string  `json:"date"`
	TicketCode string  `json:"ticket_code"`
	Price      float64 `json:"price"`
}
