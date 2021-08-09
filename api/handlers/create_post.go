package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/fiber-api-with-simple-cache/pkg/infra/cache"
	"github.com/italorfeitosa/fiber-api-with-simple-cache/pkg/infra/memory"
	"github.com/italorfeitosa/fiber-api-with-simple-cache/pkg/usecases"
)

type CreatePostHandler struct {
	CreatePostUseCase *usecases.CreatePostUseCase
}

func NewCreatePostHandler() CreatePostHandler {
	postRepo := memory.NewMemoryPostRepo()
	redisClient := cache.NewRedisClient()
	useCase := usecases.NewCreatePostUseCase(postRepo, redisClient)

	return CreatePostHandler{
		CreatePostUseCase: useCase,
	}
}

func (h *CreatePostHandler) Handle(c *fiber.Ctx) error {
	var dto usecases.CreatePostDTO
	err := c.BodyParser(&dto)

	if err != nil {
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	h.CreatePostUseCase.Perform(dto)

	return c.JSON(&fiber.Map{
		"message": "Create posts route",
	})
}
