package core

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SDatabase struct {
	Instance *gorm.DB
}

func NewDatabaseConnection(dsn string) *SDatabase {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database := &SDatabase{
		 Instance: conn,
	}
	return database
}

func Migrate(dsn string, dst ...interface{}) {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	conn.AutoMigrate(dst...)
}

func InitDatabase(dsn string, dst ...interface{}) {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	conn.Exec("CREATE DATABASE IF NOT EXISTS auth")
	conn.Exec("GRANT ALL PRIVILEGES ON auth.* TO 'user'@'%'")
}