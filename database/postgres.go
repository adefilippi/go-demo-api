package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	"fmt"
)

func setupPostgres(dsn string) (db *gorm.DB, err error) {
	params := extractParamsFromDSN(dsn)
	createDBDsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", params["host"], params["user"], params["password"], params["port"])
	database, err := gorm.Open(postgres.Open(createDBDsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
	})

	_ = database.Exec("CREATE DATABASE " + params["dbname"] + ";")
	dbInstance, _ := database.DB()
	_ = dbInstance.Close()

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
