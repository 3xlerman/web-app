package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	MongoDBPassword string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file load error: " + err.Error())
	} else {

		MongoDBPassword = os.Getenv("MONGODB_PASSWORD")

		fmt.Println("Configs loaded successfully.")
	}
}
