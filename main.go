package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/adefilippi/go-demo-api/fixtures"
	"github.com/adefilippi/go-demo-api/repository"
	"github.com/adefilippi/go-demo-api/service/env"
	"github.com/adefilippi/go-demo-api/service/router"
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
//	@description				Authentication with Bearer Token (JWT)

// 	@securityDefinitions.apikey BearerAuth
// 	@in 						header
// @name 						Authorization
// @description 				Authentication with Bearer Token (JWT)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	env.Init(".env")
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

	repository.Setup()
	r := router.SetupRouter()
	r.Run(":8080")
}
