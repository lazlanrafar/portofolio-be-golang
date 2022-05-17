package controllers

import (
	"lazlanrafar/db"
	"lazlanrafar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAll(c *gin.Context) {
	db := db.SetupDB()
	var works []models.Work

	if err := db.Find(&works).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, works)
}