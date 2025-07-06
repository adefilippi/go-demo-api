package repository

import (
	"github.com/syneido/go-api-core/database"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup(configPath string) {
	db = database.Setup(configPath)
}
