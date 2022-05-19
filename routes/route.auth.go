package route

import (
	util "lazlanrafar/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, router *gin.RouterGroup) {
	authMiddleware := util.Sign()
	router.POST("/login", authMiddleware.LoginHandler)
}