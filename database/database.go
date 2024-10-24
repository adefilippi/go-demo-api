package database

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"

	"example/web-service-gin/entity"
	"example/web-service-gin/service/env"
)

var db *gorm.DB
var err error

func Setup() {

	dsn := env.GetEnvVariable("DATABASE_URL")

	if dsn == "" {
		panic("No DSN provided")
	}

	db_type := env.GetEnvVariable("DATABASE_TYPE")

	if db_type == "" {
		panic("No Database Type provided")
	}

	if db_type == "postgres" {
		db, err = setupPostgres(dsn)
		if err != nil {
			panic(fmt.Sprintf("failed to connect database: %v", err))
		}
	} else {
		panic("Unsupported Database Type")
	}

	sqlDB, err := db.DB()

	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
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

func GetDB() *gorm.DB {
	if db == nil {
		Setup()
	}
	return db
}
