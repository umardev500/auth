package injector

import (
	"auth/storage"
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

func NewRedisInjector() fiber.Storage {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return storage.NewRedisStorage(ctx, client)
}
