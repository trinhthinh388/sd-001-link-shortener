package main

import (
	core "app/core/db"
	"app/services/middleware"
	"fmt"
)

const DSN = "postgres://postgres:postgres@localhost:5432/pocketlink"

func main() {
	fmt.Println("Setting up database...")
	core.Setup(DSN)

	fmt.Println("Starting server...")
	middleware.Start()
}
