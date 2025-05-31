package auth

import (
	core "app/core/db"
	"app/services/auth/constants"
	"app/services/auth/domains/models"
	"fmt"
)

const DSN = "postgres://postgres:postgres@localhost:5432/pocketlink"

func Start() {
	fmt.Println("Starting Auth Service...")
	core.Migrate(DSN, constants.AUTH_SCHEMA,
		&models.Permission{},
		&models.Role{},
	)

	permissions := []models.Permission{
		{Permission: constants.URL_CREATE},
		{Permission: constants.URL_LIST},
		{Permission: constants.URL_UPDATE},
		{Permission: constants.URL_DELETE},
		{Permission: constants.URL_VIEW},
	}
	core.Seed(DSN, constants.AUTH_SCHEMA, permissions)

	roles := []models.Role{
		{Role: constants.ADMIN},
		{Role: constants.SUPER_ADMIN},
		{Role: constants.USER},
	}
	core.Seed(DSN, constants.AUTH_SCHEMA, roles)

	rolesPermissions := []models.RolesPermissions{
		{Role: constants.ADMIN, Permission: constants.URL_LIST},
		{Role: constants.ADMIN, Permission: constants.URL_VIEW},
		{Role: constants.ADMIN, Permission: constants.URL_UPDATE},
		{Role: constants.ADMIN, Permission: constants.URL_CREATE},
		{Role: constants.ADMIN, Permission: constants.URL_DELETE},

		{Role: constants.SUPER_ADMIN, Permission: constants.URL_LIST},
		{Role: constants.SUPER_ADMIN, Permission: constants.URL_VIEW},
		{Role: constants.SUPER_ADMIN, Permission: constants.URL_UPDATE},
		{Role: constants.SUPER_ADMIN, Permission: constants.URL_CREATE},
		{Role: constants.SUPER_ADMIN, Permission: constants.URL_DELETE},

		{Role: constants.USER, Permission: constants.URL_LIST},
		{Role: constants.USER, Permission: constants.URL_VIEW},
	}
	core.Seed(DSN, constants.AUTH_SCHEMA, rolesPermissions)
}
