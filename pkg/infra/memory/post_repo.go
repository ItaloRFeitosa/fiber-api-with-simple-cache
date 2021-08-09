package inmemmory

import "github.com/italorfeitosa/fiber-api-with-simple-cache/pkg/domain"

type PostRepo struct {
    mu sync.RWMutex
	Posts []domain.Post
}

func NewPostRepo() *PostRepo {
	var posts []domain.Post
	return &PostRepo{
		Posts: posts,
	}
}

func (r *PostRepo) Insert(post domain.Post) error{
	r.mu.Lock()
	r.Posts = append(r.Posts, post)
	r.mu.Unlock()
	return nil
}
