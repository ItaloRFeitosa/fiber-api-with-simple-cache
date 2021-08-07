package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	setupPostRoutes(app)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"ok":      true,
			"message": "It works!",
		})
	})
}
