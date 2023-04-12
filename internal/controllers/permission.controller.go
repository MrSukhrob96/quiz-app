package controllers

import (
	"quiz-app/internal/services"

	"github.com/gofiber/fiber/v2"
)

type PermissionController interface {
	Index(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type permissionController struct {
	permissionService services.PermissionService
}

func NewPermissionController(
	permissionService services.PermissionService,
) PermissionController {
	return &permissionController{
		permissionService: permissionService,
	}
}

func (handler *permissionController) Index(ctx *fiber.Ctx) error {
	return nil
}

func (handler *permissionController) Store(ctx *fiber.Ctx) error {
	return nil
}

func (handler *permissionController) Show(ctx *fiber.Ctx) error {
	return nil
}

func (handler *permissionController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (handler *permissionController) Delete(ctx *fiber.Ctx) error {
	return nil
}
