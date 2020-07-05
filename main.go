package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")


	// Create nginx-server
	app := gin.Default()

	// Test-request
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Run application
	err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
}
