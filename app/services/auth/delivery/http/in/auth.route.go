package delivery

import (
	"app/core/interfaces"
	"app/services/auth/constants"
	"app/services/auth/controllers"

	"github.com/gofiber/fiber/v2"
)

type SAuthInRoute struct {
	interfaces.Route
	App *fiber.App
	AuthController *controllers.SAuthController
}

func (r *SAuthInRoute) Setup() {
	r.App.Post(constants.AuthRouteLogIn, r.AuthController.LogIn)
	r.App.Post(constants.AuthRouteLogOut, r.AuthController.LogOut)
}