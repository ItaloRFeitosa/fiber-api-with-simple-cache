package cache_test

import (
	"testing"

	"github.com/italorfeitosa/fiber-api-with-simple-cache/pkg/infra/cache"
)

func TestRedisClient(t *testing.T) {
	t.Run("Should create a record on redis", func(t *testing.T) {
		redisClient := cache.NewRedisClient()

		key := "key"
		wantedValue := "value"

		err := redisClient.Set(key, wantedValue)

		if err != nil {
			t.Fatalf("Error on set record")
		}

		gotValue, err := redisClient.Get(key)

		if err != nil {
			t.Fatalf("Error on get record")
		}

		if gotValue != wantedValue {
			t.Fatalf("Record Values mismatched")
		}

		_, err = redisClient.Del(key)

		if err != nil {
			t.Fatalf("Error on delete record")
		}
	})

}
