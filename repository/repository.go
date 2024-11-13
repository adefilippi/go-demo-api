package repository

import (
	"gorm.io/gorm"
	"github.com/syneido/go-demo-api/database"
)

var db *gorm.DB

func Setup() {
	db = database.Setup()
}
