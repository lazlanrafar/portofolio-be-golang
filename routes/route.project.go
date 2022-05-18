package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	resultsProject "lazlanrafar/controllers/project-controllers/results"
)

func InitProjectRoutes(db *gorm.DB, router *gin.RouterGroup) {
	// Project results
	resultsProjectRepository := resultsProject.NewRepositoryResults(db)
	resultsProjectService := resultsProject.NewServiceResults(resultsProjectRepository)
	resultsProjectHandler := resultsProject.NewHandlerResults(resultsProjectService)

	router.GET("/project", resultsProjectHandler.ResultsProjectHandler)

}