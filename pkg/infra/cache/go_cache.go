package cache

import (
	"log"

	"github.com/patrickmn/go-cache"
)

type GoCache struct {
	c *cache.Cache
}

func logGoCacheError(err error) {
	log.Println("[GO CACHE ERROR]:", err)
}

func NewGoCache() *GoCache {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)

	return &GoCache{
		c,
	}
}

func (gc *GoCache) Set(key, value string) error {
	gc.c.Set(key, value, cache.NoExpiration)

	return nil
}

func (gc *GoCache) Get(key string) (string, error) {
	val, err := gc.Get(key)

	if err != nil {
		logGoCacheError(err)
		return "", err
	}

	return val, nil
}

func (gc *GoCache) Del(key string) (int64, error) {
	val, err := gc.Del(key)

	if err != nil {
		logGoCacheError(err)
		return val, err
	}

	return val, nil
}
