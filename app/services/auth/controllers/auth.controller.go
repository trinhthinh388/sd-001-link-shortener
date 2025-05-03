package controllers

import (
	"app/services/auth/models"
	"app/services/auth/usecases"

	"github.com/gofiber/fiber/v2"
)

type SAuthController struct {
	Usecase *usecases.SAuthUsecase
}

func (c *SAuthController) LogIn(ctx *fiber.Ctx) error {
	request := new(models.AuthLoginRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := c.Usecase.LogIn(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (*SAuthController) LogOut(app *fiber.Ctx) error {
	return nil
}