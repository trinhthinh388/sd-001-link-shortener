package interfaces

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Register(ctx *fiber.Ctx) error
} 