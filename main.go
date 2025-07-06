package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/lunmy/go-demo-api/fixtures"
	"github.com/lunmy/go-demo-api/handler"
	"github.com/lunmy/go-demo-api/repository"
	"github.com/lunmy/go-demo-api/service/router"
	"github.com/lunmy/go-api-core/database"
	"github.com/lunmy/go-api-core/event"
	"github.com/lunmy/go-api-core/service/env"
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
	// Get current path
	path, _ := os.Getwd()

	env.Init(path + "/.env")
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

	
	db := database.Setup("config/database.yml")
	dispatcher := event.NewSimpleDispatcher()

	modelRepo := repository.NewGenericRepository(db)
	modelHandler := handler.NewModelHandler(modelRepo, dispatcher)

	mediaObjectRepo := repository.NewGenericRepository(db)
	mediaObjectHandler := handler.NewMediaObjectHandler(mediaObjectRepo, dispatcher)

	r := router.SetupRouter(modelHandler, mediaObjectHandler)
	r.Run(":8080")
}