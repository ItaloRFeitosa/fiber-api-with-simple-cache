package inmemmory

import "github.com/italorfeitosa/fiber-api-with-simple-cache/pkg/domain"

type PostRepo struct {
	Posts []domain.Post
}

func NewPostRepo() *PostRepo {
	var posts []domain.Post
	return &PostRepo{
		Posts: posts,
	}
}

func (repo *PostRepo) Insert(post domain.Post) {
	repo.Posts = append(repo.Posts, post)
}
