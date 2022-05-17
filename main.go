package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	config "lazlanrafar/configs"
)

func main() {
	router := setupRouter()
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	// Database Connection
	db := config.Connection()
	fmt.Println(db)

	// Init Router
	router := gin.Default()

	return router
}