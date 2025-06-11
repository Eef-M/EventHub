package repository

import (
	"github.com/Eef-M/EventHub/backend/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllFeedback(db *gorm.DB, eventID uuid.UUID) ([]dto.EventFeedbacksDTO, error) {
	var feedbacks []dto.EventFeedbacksDTO

	err := db.Table("event_feedbacks").
		Select("event_feedbacks.comment, event_feedbacks.rating, event_feedbacks.created_at, users.username, users.avatar_url").
		Joins("JOIN users ON event_feedbacks.user_id = users.id").
		Order("event_feedbacks.created_at DESC").
		Where("event_feedbacks.event_id = ?", eventID).
		Scan(&feedbacks).Error

	return feedbacks, err
}
