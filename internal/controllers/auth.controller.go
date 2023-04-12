package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
	Me(ctx *fiber.Ctx) error
	VerifyCode(ctx *fiber.Ctx) error
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (controller *authController) Login(ctx *fiber.Ctx) error {
	return nil
}

func (controller *authController) Register(ctx *fiber.Ctx) error {
	return nil
}

func (controller *authController) Me(ctx *fiber.Ctx) error {
	return nil
}

func (controller *authController) RefreshToken(ctx *fiber.Ctx) error {
	return nil
}

func (controller *authController) VerifyCode(ctx *fiber.Ctx) error {
	return nil
}
