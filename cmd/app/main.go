package main

import (
	"log"
	"quiz-app/router"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/docs/", swagger.New(swagger.Config{
		URL:         "http://localhost:5555/api/",
		DeepLinking: false,
	}))

	router.START_API(app)

	log.Fatal(app.Listen(":5555"))
}
