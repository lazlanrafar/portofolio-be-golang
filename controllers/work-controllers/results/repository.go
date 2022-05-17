package resultsWork

import (
	model "lazlanrafar/models"

	"gorm.io/gorm"
)

type Repository interface {
	ResultsWorkRepository() (*[]model.Work, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ResultsWorkRepository() (*[]model.Work, string) {
	var works []model.Work
	db := r.db.Model(&works)
	errorCode := make(chan string, 1)

	resultsWorks := db.Debug().Select("*").Find(&works)
	if resultsWorks.Error != nil {
		errorCode <- resultsWorks.Error.Error()
	}else {
		errorCode <- "success"
	}

	return &works, <-errorCode
}