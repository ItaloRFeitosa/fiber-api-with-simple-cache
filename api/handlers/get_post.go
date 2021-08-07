package handlers

import "github.com/gofiber/fiber/v2"

type GetPostHandler struct {
	CreatePostUseCase string
}

func NewGetPostHandler() GetPostHandler {

	return GetPostHandler{
		CreatePostUseCase: "Test",
	}
}

func (h *GetPostHandler) Handle(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"message": "Create posts route",
	})
}
