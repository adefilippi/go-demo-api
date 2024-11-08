package repository

import (
	"gorm.io/gorm"
	"example/web-service-gin/database"
)

var db *gorm.DB

func Setup() {
	db = database.Setup()
}
