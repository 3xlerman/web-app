package main

import (
	"fmt"
	"github.com/3xlerman/web-app/app/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	// Create nginx-server
	app := gin.Default()

	// Test-request
	app.GET("/home", handlers.Home)

	// Run application
	err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
}
