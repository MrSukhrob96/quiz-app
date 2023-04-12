package main

import (
	"log"
	"quiz-app/router"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"

	_ "github.com/swaggo/fiber-swagger/example/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	app := fiber.New()

	app.Get("/swagger/*", swagger.WrapHandler)

	router.START_API(app)

	log.Fatal(app.Listen(":5555"))
}
