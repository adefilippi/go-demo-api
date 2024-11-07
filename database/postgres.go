package database

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"fmt"
	"time"
)

func setupPostgres(dsn string) (db *gorm.DB, err error) {
	params := extractParamsFromDSN(dsn)
	createDBDsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", params["host"], params["user"], params["password"], params["port"])
	database, err := gorm.Open(postgres.Open(createDBDsn), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		DisableAutomaticPing: true,
	})

	_ = database.Exec("CREATE DATABASE " + params["dbname"] + ";")
	dbInstance, _ := database.DB()
	_ = dbInstance.Close()

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Errorf("Failed to get sql.DB: %v", err)
	}

	if err := manualPingWithRetry(sqlDB, 5, 2*time.Second); err != nil {
		return nil, fmt.Errorf("Database ping failed after retries: %v", err)
	}

	return db, err
}

func manualPingWithRetry(sqlDB *sql.DB, maxRetries int, interval time.Duration) error {
	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("Attempt %d to ping the database...\n", i)
		if err := sqlDB.Ping(); err != nil {
			fmt.Printf("Ping failed: %v\n", err)
			if i == maxRetries {
				return fmt.Errorf("all retries failed: %w", err)
			}
			time.Sleep(interval)
			continue
		}
		return nil // Ping successful
	}
	return nil // Should never reach this line
}
