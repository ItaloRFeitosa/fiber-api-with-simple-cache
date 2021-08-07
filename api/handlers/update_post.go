package handlers

import "github.com/gofiber/fiber/v2"

type UpdatePostHandler struct {
	CreatePostUseCase string
}

func NewUpdatePostHandler() UpdatePostHandler {

	return UpdatePostHandler{
		CreatePostUseCase: "Test",
	}
}

func (h *UpdatePostHandler) Handle(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"message": "Create posts route",
	})
}
