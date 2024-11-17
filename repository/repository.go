package repository

import (
	"gorm.io/gorm"
	"github.com/adefilippi/go-demo-api/database"
)

var db *gorm.DB

func Setup(configPath string) {
	db = database.Setup(configPath)
}
