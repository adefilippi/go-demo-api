package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"example/web-service-gin/repository"
	"example/web-service-gin/service"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	repository.Setup()
	router := service.SetupRouter()
	router.Run("localhost:8080")
}
