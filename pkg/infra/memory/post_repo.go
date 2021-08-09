package memory

import "github.com/italorfeitosa/fiber-api-with-simple-cache/pkg/domain"

type MemoryPostRepo struct {
    mu sync.RWMutex
	posts []domain.Post
}

func NewMemoryPostRepo() *MemoryPostRepo {
	var posts []domain.Post
	return &MemoryPostRepo{
		posts: posts,
	}
}

func (r *MemoryPostRepo) Insert(post domain.Post) error{
	r.mu.Lock()
	defer r.mu.Unlock()
	r.posts = append(r.posts, post)
	return nil

func (r *MemoryPostRepo) FindAll() ([]domain.Post, error) {
    r.mu.Lock()
	defer r.mu.Unlock()
    return r.posts, nil
}
