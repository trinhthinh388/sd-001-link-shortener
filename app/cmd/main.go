package main

import (
	auth "app/core/auth"
	core "app/core/db"
	"app/services/middleware"
	"fmt"
)

const DSN = "postgres://postgres:postgres@localhost:5432/pocketlink"

func main() {
	core.Setup(DSN)
	auth.CreateProject()
	fmt.Println("Starting server...")
	middleware.Start()
}
