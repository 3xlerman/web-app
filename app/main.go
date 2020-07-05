package main

import (
	"fmt"
	"github.com/3xlerman/web-app/app/config"
	"github.com/3xlerman/web-app/app/database"
	"github.com/3xlerman/web-app/app/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	// Load config
	config.Load()

	// Create nginx-server
	app := gin.Default()

	// Connect to MongoDB
	database.Connect()

	// Test-request
	app.GET("/home", handlers.Home)
	app.GET("/person", handlers.Person)

	// Run application
	err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
}
