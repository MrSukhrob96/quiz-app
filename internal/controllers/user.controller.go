package controllers

import (
	"quiz-app/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Index(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (handler *userController) Index(ctx *fiber.Ctx) error {
	ctx.Status(201)
	return ctx.JSON(fiber.Map{"name": "djkfhfjds", "status": 200})
}

func (handler *userController) Store(ctx *fiber.Ctx) error {
	return nil
}

func (handler *userController) Show(ctx *fiber.Ctx) error {
	return nil
}

func (handler *userController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (handler *userController) Delete(ctx *fiber.Ctx) error {
	return nil
}
