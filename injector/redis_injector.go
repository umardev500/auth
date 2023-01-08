package injector

import (
	"auth/storage"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

func NewRedisInjector() fiber.Storage {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return storage.NewRedisStorage(client)
}
