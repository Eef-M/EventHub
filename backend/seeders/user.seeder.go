package seeders

import (
	"fmt"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SeedUsers() []models.User {
	rawUsers := []struct {
		Username  string
		FirstName string
		LastName  string
		Email     string
		Password  string
		Role      models.Roles
	}{
		{
			Username:  "gojo",
			FirstName: "Satoru",
			LastName:  "Gojo",
			Email:     "gojo@example.com",
			Password:  "gojo123",
			Role:      "organizer",
		},
		{
			Username:  "geto",
			FirstName: "Suguru",
			LastName:  "Geto",
			Email:     "geto@example.com",
			Password:  "geto123",
			Role:      "organizer",
		},
		{
			Username:  "yuji",
			FirstName: "Yuji",
			LastName:  "Itadori",
			Email:     "yuji@example.com",
			Password:  "yuji123",
			Role:      "participant",
		},
		{
			Username:  "yuta",
			FirstName: "Yuta",
			LastName:  "Okkotsu",
			Email:     "yuta@example.com",
			Password:  "yuta123",
			Role:      "participant",
		},
	}

	var users []models.User

	for _, ru := range rawUsers {
		hashedPassword, err := hashPassword(ru.Password)
		if err != nil {
			panic("Failed to hash password: " + err.Error())
		}

		defaultAvatar := fmt.Sprintf(
			"%s/uploads/avatars/default_avatar.png",
			initializers.BaseURL,
		)

		user := models.User{
			ID:        uuid.New(),
			Username:  ru.Username,
			FirstName: ru.FirstName,
			LastName:  ru.LastName,
			Email:     ru.Email,
			Password:  hashedPassword,
			Role:      ru.Role,
			AvatarURL: defaultAvatar,
		}

		initializers.DB.Create(&user)
		users = append(users, user)
	}

	return users
}
