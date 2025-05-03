package config

import (
	core "app/core/db"
	"fmt"
)

type SAuthDatabase struct {
	Database *core.SDatabase
}

var AuthDatabase *SAuthDatabase

func init() {
	fmt.Println("Initializing auth database connection...")
	AuthDatabase = &SAuthDatabase{
	Database: core.NewDatabaseConnection("user:password@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"),
	}
}
