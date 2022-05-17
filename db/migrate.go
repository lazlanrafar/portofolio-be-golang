package db

import (
	"lazlanrafar/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.Work{})
}