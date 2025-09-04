package repository

import (
	"github.com/Eef-M/EventHub/backend/dto"
	"github.com/Eef-M/EventHub/backend/models"
)

func GetMe(user models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      string(user.Role),
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt,
	}
}
