package controllers

import (
	"quiz-app/internal/services"

	"github.com/gofiber/fiber/v2"
)

type RoleController interface {
	Index(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type roleController struct {
	roleService services.RoleService
}

func NewRoleController(
	roleService services.RoleService,
) RoleController {
	return &roleController{
		roleService: roleService,
	}
}

func (handler *roleController) Index(ctx *fiber.Ctx) error {
	return nil
}

func (handler *roleController) Store(ctx *fiber.Ctx) error {
	return nil
}

func (handler *roleController) Show(ctx *fiber.Ctx) error {
	return nil
}

func (handler *roleController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (handler *roleController) Delete(ctx *fiber.Ctx) error {
	return nil
}
