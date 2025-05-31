package middleware

import (
	core "app/core/db"
	"app/services/middleware/domains/models"
	"fmt"
)

const DSN = "postgres://postgres:postgres@localhost:5432/pocketlink"

func Start() {
	fmt.Println("Starting middleware...")
	core.Migrate(DSN, &models.URL{})
}
