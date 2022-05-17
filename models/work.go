package models

type Work struct {
	Id       int    `gorm:"primary_key"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Tools    string `json:"tools"`
	DesignBy string `json:"designBy"`
	Link     string `json:"link"`
	Image    string `json:"image"`
}