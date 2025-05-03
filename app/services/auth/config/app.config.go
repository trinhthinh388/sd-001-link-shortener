package config

import (
	config "app/services/auth/config/db"
	"app/services/auth/controllers"
	delivery "app/services/auth/delivery/http/in"
	"app/services/auth/usecases"

	"github.com/gofiber/fiber/v2"
)

type SAuthAppContext struct {
	app *fiber.App
	routes *delivery.SAuthInRoute
	usecases *usecases.SAuthUsecase
	controllers *controllers.SAuthController
}

func NewAuthContext() *SAuthAppContext {
	return &SAuthAppContext{}
}

func (c *SAuthAppContext) CreateApp() *fiber.App {
	c.app = fiber.New(fiber.Config{})
	c.createRouter()
	return c.app
}

func (c *SAuthAppContext) createUseCases() *usecases.SAuthUsecase {
	c.usecases = usecases.NewAuthUsecase(config.AuthDatabase)
	return c.usecases
}

func (c *SAuthAppContext) createControllers() *controllers.SAuthController {
	c.controllers = &controllers.SAuthController{
		Usecase: c.createUseCases(),
	}
	return c.controllers
}

func (c *SAuthAppContext) createRouter() {
	c.routes = &delivery.SAuthInRoute{
		App: c.app,
		AuthController: c.createControllers(),
	}
	c.routes.Setup()
}