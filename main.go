package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
}
