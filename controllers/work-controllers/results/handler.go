package resultsWork

import "github.com/gin-gonic/gin"

type handler struct {
	service Service
}

func NewHandlerResults(service Service) *handler {
	return &handler{service}
}

func (h *handler) ResultsWorkHandler(c *gin.Context) {
	works, err := h.service.ResultsWorkRepository()
	if err != "success" {
		c.JSON(500, gin.H{
			"message": err,
		})
	} else {
		c.JSON(200, gin.H{
			"data": works,
		})
	}
}