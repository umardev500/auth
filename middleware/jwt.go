package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JwtMiddleware() func(*fiber.Ctx) error {
	secret := os.Getenv("SECRET")

	f := jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})

	return f
}
