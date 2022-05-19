package model

import "time"

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

