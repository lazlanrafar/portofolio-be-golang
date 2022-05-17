package model

import "time"

type EntityWork struct {
	ID        int    `gorm:"primary_key"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Tools     string `json:"tools"`
	DesignBy  string `json:"designBy"`
	Link      string `json:"link"`
	Image     string `json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
