package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"fmt"
	"strings"

	"example/web-service-gin/entity"
	"example/web-service-gin/service/env"
)

var db *gorm.DB

func Setup() {
	var err error
	dsn := env.GetEnvVariable("DATABASE_URL")

	if dsn == "" {
		panic("No DSN provided")
	}

	params := extractParamsFromDSN(dsn)
	createDBDsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", params["host"], params["user"], params["password"], params["port"])
	database, err := gorm.Open(postgres.Open(createDBDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	_ = database.Exec("CREATE DATABASE " + params["dbname"] + ";")
	dbInstance, _ := database.DB()
	_ = dbInstance.Close()

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.AutoMigrate(&entity.Model{}, &entity.MediaObject{})
}

func extractParamsFromDSN(dsn string) map[string]string {
	params := make(map[string]string)

	// Split the DSN into key=value pairs
	parts := strings.Split(dsn, " ")

	// Iterate over the parts and split each by '='
	for _, part := range parts {
		keyValue := strings.SplitN(part, "=", 2)
		if len(keyValue) == 2 {
			params[keyValue[0]] = keyValue[1]
		}
	}

	return params
}
