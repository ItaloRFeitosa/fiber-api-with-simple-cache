package handlers

import "github.com/gofiber/fiber/v2"

type AllPostsHandler struct {
	CreatePostUseCase string
}

func NewAllPostsHandler() AllPostsHandler {

	return AllPostsHandler{
		CreatePostUseCase: "Test",
	}
}

func (h *AllPostsHandler) Handle(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"message": "Create posts route",
	})
}
