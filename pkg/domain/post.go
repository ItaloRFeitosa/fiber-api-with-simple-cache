package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"created_at"`
	OwnerID  string    `json:"owner_id"`
}

func NewPost(title, content, ownerId string) Post {
	return Post{
		ID:       uuid.NewString(),
		Title:    title,
		Content:  content,
		OwnerID:  ownerId,
		CreateAt: time.Now(),
	}
}
