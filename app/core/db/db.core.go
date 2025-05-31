package core

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type SDatabase struct {
	Instance *gorm.DB
}

var logInstance = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,        // Don't include params in the SQL log
		Colorful:                  false,       // Disable color
	},
)

func NewDatabaseConnection(dsn string) *SDatabase {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logInstance,
	})
	if err != nil {
		log.Fatal(err)
	}
	database := &SDatabase{
		Instance: conn,
	}
	return database
}

func Migrate(dsn string, prefix string, dst ...interface{}) {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logInstance,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix + ".",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	conn.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", prefix))
	conn.AutoMigrate(dst...)
}

func Seed[K any](dsn string, prefix string, models []K) {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logInstance,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix + ".",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	conn.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", prefix))
	for _, model := range models {
		err := conn.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(&model).Error
		if err != nil {
			fmt.Printf("Error when seeding data: %s\n", model, err)
		} else {
			fmt.Printf("Success seed data: %s\n", model)
		}
	}
}

func Setup(dsn string) {
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logInstance,
	})
	if err != nil {
		log.Fatal(err)
	}
}
