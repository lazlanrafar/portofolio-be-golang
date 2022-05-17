package model

import "time"

type Project struct {
	ID        int    `gorm:"primary_key"`
	Title     string `json:"title"`
	DesignBy  string `json:"designBy"`
	Link      string `json:"link"`
	Image     string `json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}