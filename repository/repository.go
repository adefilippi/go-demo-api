package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"example/web-service-gin/entity"
	"example/web-service-gin/service/env"
)

var db *gorm.DB

func Setup() {
	var err error
	dsn := env.GetEnvVariable("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.AutoMigrate(&entity.Model{}, &entity.MediaObject{})
}
