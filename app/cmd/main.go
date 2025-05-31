package main

import (
	authCore "app/core/auth"
	dbCore "app/core/db"
	"app/services/auth"
	"app/services/middleware"
)

const DSN = "postgres://postgres:postgres@localhost:5432/pocketlink"

func main() {
	dbCore.Setup(DSN)
	authCore.CreateProject()

	middleware.Start()
	auth.Start()
}
