package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"reflect"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/adefilippi/go-demo-api/service/env"
	"github.com/adefilippi/go-demo-api/entity"
)

var dbs map[string]*gorm.DB

// Database interface for common database functionality
type Database interface {
	Open(dsn string, config map[string]interface{}) (*gorm.DB, error)
	PingContext() error
	DB() *gorm.DB
}

var typeRegistry = make(map[string]reflect.Type)

// RegisterType registers a type in the global registry.
func RegisterType(name string, t reflect.Type) {
	typeRegistry[name] = t
}

func Setup(path string) *gorm.DB {
	dbs = make(map[string]*gorm.DB)
	file, err := os.Open(path)
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
				var pool, timeout int
				var migrate bool

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

				var config map[string]interface{}
				if config, ok = dbConfig["config"].(map[string]interface{}); !ok {
					config = make(map[string]interface{})
				}

				gormDB, err := db.Open(dsn, config)
				if err != nil {
					panic(fmt.Sprintf("Failed to setup %s database: %v", key, err))
				}

				if err := db.PingContext(); err != nil {
					panic(fmt.Sprintf("Failed to ping %s database: %v", key, err))
				}

				if pool, ok = dbConfig["pool"].(int); !ok {
					pool = 100
				}

				if timeout, ok = dbConfig["timeout"].(int); !ok {
					timeout = 10
				}

				if migrate, ok = dbConfig["migrate"].(bool); !ok {
					migrate = false
				}

				if migrate {
					if entityNames, ok := dbConfig["entities"].([]interface{}); ok {
						for _, typeName := range entityNames {
							instance, err := entity.CreateStructFromString(typeName.(string))
							if err != nil {
								fmt.Printf("Error creating instance for %s: %v\n", typeName, err)
								continue
							}

							gormDB.AutoMigrate(instance)
						}
					}

				}

				sqlDB, _ := gormDB.DB()
				sqlDB.SetMaxIdleConns(timeout)
				sqlDB.SetMaxOpenConns(pool)
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
func GetDBs() map[string]*gorm.DB {
	return dbs
}

func GetLogLevel(level string) logger.LogLevel {
	switch level {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info // Default to Info if invalid
	}
}

func GetConfig(config map[string]interface{}) gorm.Config {
	var dbConfig gorm.Config

	var logLevel string
	var ok, skipDefaultTransaction, disableNestedTransaction, fullSaveAssociations, allowGlobalUpdate bool
	var createBatchSize int

	if logLevel, ok = config["log_level"].(string); !ok {
		logLevel = "info"
	}
	dbConfig.Logger = logger.Default.LogMode(GetLogLevel(logLevel))

	if skipDefaultTransaction, ok = config["skip_default_transaction"].(bool); !ok {
		skipDefaultTransaction = false
	}
	dbConfig.SkipDefaultTransaction = skipDefaultTransaction

	if disableNestedTransaction, ok = config["disable_nested_transaction"].(bool); !ok {
		disableNestedTransaction = false
	}
	dbConfig.DisableNestedTransaction = disableNestedTransaction

	if createBatchSize, ok = config["create_batch_size"].(int); !ok {
		createBatchSize = 100
	}
	dbConfig.CreateBatchSize = createBatchSize

	if fullSaveAssociations, ok = config["full_save_associations"].(bool); !ok {
		fullSaveAssociations = false
	}
	dbConfig.FullSaveAssociations = fullSaveAssociations

	if allowGlobalUpdate, ok = config["allow_global_update"].(bool); !ok {
		allowGlobalUpdate = false
	}
	dbConfig.AllowGlobalUpdate = allowGlobalUpdate

	return dbConfig
}
