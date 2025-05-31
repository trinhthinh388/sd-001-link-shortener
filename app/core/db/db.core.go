package core

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func Migrate(dsn string, dst ...interface{}) {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logInstance,
	})
	if err != nil {
		log.Fatal(err)
	}

	conn.AutoMigrate(dst...)
}

func Setup(dsn string) {
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logInstance,
	})
	if err != nil {
		log.Fatal(err)
	}
}
