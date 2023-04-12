package router

import (
	"quiz-app/internal/controllers"
	"quiz-app/internal/repositories"
	"quiz-app/internal/services"
	"quiz-app/pkg/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	database       *gorm.DB                    = db.ConnectDatabase()

	userRepository repositories.UserRepository = repositories.NewUserRepository(database)
	userService    services.UserService        = services.NewUserService(userRepository)
	userController controllers.UserController  = controllers.NewUserController(userService)
)

func START_API(router *fiber.App) {
	api := router.Group("api")
	user := api.Group("user")

	user.Get("/", userController.Index)
	user.Post("/", userController.Store)
	user.Get("/:id", userController.Show)
	user.Put("/", userController.Update)
	user.Delete("/", userController.Delete)

}
