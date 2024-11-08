package database

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type SQLServerDB struct {
	db *gorm.DB
}

func (d *SQLServerDB) Open(dsn string, config gorm.Config) (*gorm.DB, error) {
	var err error
	// Open connection to the actual database
	d.db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		DisableAutomaticPing: true,
	})

	if err != nil {
		return nil, err
	}

	return d.db, nil
}

func (d *SQLServerDB) PingContext() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	for i := 1; i <= 5; i++ {
		fmt.Printf("Attempt %d to ping the SQL Server database...\n", i)
		if err := sqlDB.Ping(); err != nil {
			fmt.Printf("Ping failed: %v\n", err)
			time.Sleep(2 * time.Second)
			continue
		}
		return nil
	}
	return fmt.Errorf("failed to ping database after retries")
}

func (d *SQLServerDB) DB() *gorm.DB {
	return d.db
}
