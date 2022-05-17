package main

import (
	"lazlanrafar/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	db.SetupDB()

	// Router
	r := gin.Default()
	rv1 := r.Group("/api/v1")

	rv1.GET("/works", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.Run()
}