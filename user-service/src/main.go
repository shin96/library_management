package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware: logger and recovery middleware
	router := gin.Default()

	// Define a GET route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the Gin server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
