package dto

type EventFeedbacksDTO struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Comment   string `json:"comment"`
	Rating    int    `json:"rating"`
	CreatedAt string `json:"created_at"`
}
