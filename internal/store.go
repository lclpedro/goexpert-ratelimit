package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Store interface {
	Increment(key string, expiration time.Duration) (int64, error)
}

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr string) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao Redis: %w", err)
	}

	return &RedisStore{client: client}, nil
}

func (s *RedisStore) Increment(key string, expiration time.Duration) (int64, error) {
	ctx := context.Background()
	return s.client.Incr(ctx, key).Result()
}
