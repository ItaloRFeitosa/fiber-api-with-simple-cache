package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/fiber-api-with-simple-cache/api/handlers"
)

func setupPostRoutes(app *fiber.App) {
	createPostHandler := handlers.NewCreatePostHandler()
	app.Post("/posts", createPostHandler.Handle)

	app.Get("/posts", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Create posts route",
		})
	})

	app.Patch("/posts/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Create posts route",
		})
	})

	app.Get("/posts/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Create posts route",
		})
	})
}
