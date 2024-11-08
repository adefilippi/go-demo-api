package database

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"example/web-service-gin/service/env"
)

var dbs map[string]*gorm.DB

// Database interface for common database functionality
type Database interface {
	Open(dsn string, config gorm.Config) (*gorm.DB, error)
	PingContext() error
	DB() *gorm.DB
}

func Setup() *gorm.DB {
	dbs = make(map[string]*gorm.DB)
	file, err := os.Open("config/database.yml")
	if err != nil {
		panic(fmt.Errorf("Error opening YAML file: %v", err))
	}
	defer file.Close()

	var config map[string]interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic(fmt.Errorf("Error decoding YAML file: %v", err))
	}

	if database, ok := config["database"].(map[string]interface{}); ok {
		for key, value := range database {
			if dbConfig, ok := value.(map[string]interface{}); ok {
				var dsn, adapter string

				if dsn, ok = dbConfig["dsn"].(string); !ok {
					panic(fmt.Sprintf("Invalid or missing dsn for database: %s", key))
				}
				dsn = env.GetEnvVariable(dsn)

				if adapter, ok = dbConfig["adapter"].(string); !ok {
					panic(fmt.Sprintf("Invalid or missing adapter for database: %s", key))
				}

				var db Database
				switch adapter {
				case "postgres":
					db = &PostgresDB{}
				case "sqlsrv":
					db = &SQLServerDB{}
				default:
					panic(fmt.Sprintf("Unsupported adapter: %s", adapter))
				}

				gormDB, err := db.Open(dsn, gorm.Config{})
				if err != nil {
					panic(fmt.Sprintf("Failed to setup %s database: %v", key, err))
				}

				if err := db.PingContext(); err != nil {
					panic(fmt.Sprintf("Failed to ping %s database: %v", key, err))
				}

				sqlDB, _ := gormDB.DB()
				sqlDB.SetMaxIdleConns(10)
				sqlDB.SetMaxOpenConns(100)
				sqlDB.SetConnMaxLifetime(time.Hour)

				dbs[key] = gormDB
			}
		}
	}
	return dbs["default"]
}

func GetDB(database string) *gorm.DB {
	// Ckeck if database key exist in map
	if _, ok := dbs[database]; !ok {
		panic(fmt.Sprintf("Invalid database: %s", database))
	}

	return dbs[database]
}
