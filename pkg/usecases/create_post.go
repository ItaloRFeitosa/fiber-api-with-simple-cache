package usecases

import (
	"encoding/json"

	"github.com/italorfeitosa/fiber-api-with-simple-cache/pkg/domain"
)

type CreatePostUseCase struct {
	postRepo     PostRepo
	cacheService CacheService
}

type CreatePostDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	OwnerId string `json:"owner_id"`
}

func NewCreatePostUseCase(postRepo PostRepo, cacheService CacheService) *CreatePostUseCase {

	return &CreatePostUseCase{
		postRepo:     postRepo,
		cacheService: cacheService,
	}
}

func (u *CreatePostUseCase) Perform(dto CreatePostDTO) (domain.Post, error) {

	newPost := domain.NewPost(dto.Title, dto.Content, dto.OwnerId)

	u.postRepo.Insert(newPost)

	cacheKey := "post@" + newPost.ID
	postJson, _ := json.Marshal(newPost)
	cacheValue := string(postJson)
	err := u.cacheService.Set(cacheKey, cacheValue)

	if err != nil {
		return domain.Post{}, err
	}

	return newPost, nil
}
