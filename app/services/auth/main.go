package auth

import (
	core "app/core/db"
	"app/services/auth/config"
	"app/services/auth/entities"
)


func StartAuthService() {
	core.Migrate(
		"root:password@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local",
		entities.Session{},
		entities.User{},
	)
	app := config.NewAuthContext().CreateApp()
	app.Listen(":8081")
}