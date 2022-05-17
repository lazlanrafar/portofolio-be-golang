package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitProjectRoutes(db *gorm.DB, router *gin.RouterGroup) {

	router.GET("/project", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}