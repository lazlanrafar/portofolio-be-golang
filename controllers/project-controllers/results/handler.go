package resultsProject

import "github.com/gin-gonic/gin"

type handler struct {
	service Service
}

func NewHandlerResults(service Service) *handler {
	return &handler{service}
}

func (h *handler) ResultsProjectHandler(c *gin.Context) {
	projects, err := h.service.ResultsProjectRepository()
	if err != "success" {
		c.JSON(500, gin.H{
			"message": err,
		})
	} else {
		c.JSON(200, gin.H{
			"data": projects,
		})
	}
}