package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/fiber-api-with-simple-cache/api/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
