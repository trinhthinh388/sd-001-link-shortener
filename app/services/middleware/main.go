package middleware

import (
	core "app/core/db"
	"app/services/middleware/constants"
	"app/services/middleware/domains/models"
	"fmt"
)

const DSN = "postgres://postgres:postgres@localhost:5432/pocketlink"

func Start() {
	fmt.Println("Starting Middleware Service...")
	core.Migrate(DSN, constants.MIDDLEWARE_SCHEMA, &models.URL{})
}
