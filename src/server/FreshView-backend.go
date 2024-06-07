package main

import (
	"krzysztofRoz/FreshView/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()
	repository.InitializeConfig()
	repository.ConnectDataBase()
	repository.SyncDB()

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Run the server on port 8080
	router.Run(":8080")
}
