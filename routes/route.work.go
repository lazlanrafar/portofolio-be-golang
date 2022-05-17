package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	resultsWork "lazlanrafar/controllers/work-controllers/results"
)

func InitWorkRoutes(db *gorm.DB, router *gin.RouterGroup) {
	resultsWorkRepository := resultsWork.NewRepositoryResults(db)
	resultsWorkService := resultsWork.NewServiceResults(resultsWorkRepository)
	resultsWorkHandler := resultsWork.NewHandlerResults(resultsWorkService)

	router.GET("/work", resultsWorkHandler.ResultsWorkHandler)
}