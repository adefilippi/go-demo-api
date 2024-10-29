package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"example/web-service-gin/fixtures"
	"example/web-service-gin/repository"
	"example/web-service-gin/service/env"
	"example/web-service-gin/service/router"
)

// swag init --parseDependency --parseInternal && go run .

//	@title			Commercial Info API
//	@version		1.0
//	@description	This is a sample API .
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "fixtures":
			log.Println("Loading fixtures...")
			fixtures.SetupFixtures()
		case "dump":
			log.Println("Dumping fixtures...")
			fixtures.DumpFixtures()
		}

		return
	}
	time.Sleep(8 * time.Second)
	env.Init(".env")
	repository.Setup()
	r := router.SetupRouter()
	r.Run(":8080")
}
