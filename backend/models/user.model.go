package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Roles string

const (
	IsUser  Roles = "user"
	IsAdmin Roles = "admin"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	FirstName string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Role      Roles     `gorm:"type:user_role;default:'user'" json:"role"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAT time.Time `gorm:"default:now()" json:"updated_now()"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	validRoles := map[Roles]bool{
		IsUser:  true,
		IsAdmin: true,
	}

	if _, exists := validRoles[u.Role]; !exists {
		return fmt.Errorf("invalid role: %s", u.Role)
	}

	return nil
}
