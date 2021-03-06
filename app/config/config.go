package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config variables
var (
	MongoDBPassword string
)

// Load config
func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file load error: " + err.Error())
		return
	}

	MongoDBPassword = os.Getenv("MONGODB_PASSWORD")

	fmt.Println("Configs loaded successfully.")

}
