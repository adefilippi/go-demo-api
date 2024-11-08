package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

type PostgresDB struct {
	db *gorm.DB
}

func (p *PostgresDB) Open(dsn string, config gorm.Config) (*gorm.DB, error) {

	params := extractParamsFromDSN(dsn)
	createDBDsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		params["host"], params["user"], params["password"], params["port"])

	// Open connection for creating the database
	tempDB, err := gorm.Open(postgres.Open(createDBDsn), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		DisableAutomaticPing: true,
	})
	if err != nil {
		return nil, err
	}

	_ = tempDB.Exec("CREATE DATABASE " + params["dbname"] + ";")
	sqlDB, _ := tempDB.DB()
	_ = sqlDB.Close()

	// Open connection to the actual database
	p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		DisableAutomaticPing: true,
	})

	if err != nil {
		return nil, err
	}

	return p.db, nil
}

func (p *PostgresDB) PingContext() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}

	for i := 1; i <= 5; i++ {
		fmt.Printf("Attempt %d to ping the postgres database...\n", i)
		if err := sqlDB.Ping(); err != nil {
			fmt.Printf("Ping failed: %v\n", err)
			time.Sleep(2 * time.Second)
			continue
		}
		return nil
	}
	return fmt.Errorf("failed to ping the postgres database after retries")
}

func (p *PostgresDB) DB() *gorm.DB {
	return p.db
}

func extractParamsFromDSN(dsn string) map[string]string {
	params := make(map[string]string)
	parts := strings.Split(dsn, " ")

	for _, part := range parts {
		keyValue := strings.SplitN(part, "=", 2)
		if len(keyValue) == 2 {
			params[keyValue[0]] = keyValue[1]
		}
	}
	return params
}
