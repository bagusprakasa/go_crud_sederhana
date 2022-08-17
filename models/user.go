package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Level     string    `json:"level"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func FormatUser(user User) User {
	userFormat := User{}
	userFormat.ID = user.ID
	userFormat.Name = user.Name
	userFormat.Gender = user.Gender
	userFormat.Level = user.Level

	return userFormat
}
