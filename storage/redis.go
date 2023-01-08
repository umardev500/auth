package storage

import (
	"context"
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

func (r *RedisStorage) withTimeout(dur int, f func(context.Context)) {
	if dur == 0 {
		dur = 10 // set default timeout duration to 10 seconds
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(dur)*time.Second)
	defer cancel()
	f(ctx)
}

func (r *RedisStorage) Set(key string, val []byte, exp time.Duration) (err error) {
	r.withTimeout(0, func(ctx context.Context) {
		if len(key) == 0 || len(val) == 0 {
			return
		}

		if exp > 0 {
			err = r.client.Set(ctx, key, val, exp).Err()
			return
		}
	})

	return
}

func (r *RedisStorage) Get(key string) (val []byte, err error) {

	r.withTimeout(0, func(ctx context.Context) {
		val, err = r.client.Get(ctx, key).Bytes()
		if err == redis.Nil {
			val = nil
			err = nil
		}

		if err != nil {
			val = nil
		}
	})

	return
}

func (r *RedisStorage) Delete(key string) (err error) {
	r.withTimeout(0, func(ctx context.Context) {
		err = r.client.Del(ctx, key).Err()
	})

	return
}

func (r *RedisStorage) Reset() (err error) {
	r.withTimeout(0, func(ctx context.Context) {
		keys, err := r.client.Keys(r.ctx, "*").Result()
		if err != nil {
			return
		}

		if len(keys) > 0 {
			_, err = r.client.Del(r.ctx, keys...).Result()
			if err != nil {
				return
			}
		}

	})

	return nil
}

func (r *RedisStorage) Close() (err error) {
	r.withTimeout(0, func(ctx context.Context) {
		err = r.client.Close()
	})

	return
}
