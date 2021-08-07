package usecases

import "github.com/italorfeitosa/fiber-api-with-simple-cache/pkg/domain"

type PostRepo interface {
	Insert(post domain.Post)
}

type CacheService interface {
	Set(key, value string) error
}
