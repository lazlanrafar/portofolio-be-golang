package resultsProject

import (
	model "lazlanrafar/models"

	"gorm.io/gorm"
)

type Repository interface {
	ResultsProjectRepository() (*[]model.Project, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ResultsProjectRepository() (*[]model.Project, string) {
	var projects []model.Project
	db := r.db.Model(&projects)
	errorCode := make(chan string, 1)

	resultsProjects := db.Debug().Select("*").Find(&projects)
	if resultsProjects.Error != nil {
		errorCode <- resultsProjects.Error.Error()
	}else {
		errorCode <- "success"
	}

	return &projects, <-errorCode
}