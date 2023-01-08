package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

type RedisStorage struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedisStorage(client *redis.Client) fiber.Storage {
	cont := context.Background()

	return &RedisStorage{
		ctx:    cont,
		client: client,
	}
}

func (r *RedisStorage) Set(key string, val []byte, exp time.Duration) error {
	if len(key) == 0 || len(val) == 0 {
		return nil
	}

	if exp > 0 {
		return r.client.Set(r.ctx, key, val, exp).Err()
	}

	return nil
}

func (r *RedisStorage) Get(key string) ([]byte, error) {
	fmt.Println("getter")
	val, err := r.client.Get(r.ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		fmt.Println("disini", err)
		return nil, err
	}

	return val, nil
}

func (r *RedisStorage) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r *RedisStorage) Reset() error {
	keys, err := r.client.Keys(r.ctx, "*").Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		_, err := r.client.Del(r.ctx, keys...).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RedisStorage) Close() error {
	return r.client.Close()
}
