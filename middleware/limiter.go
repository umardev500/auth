package middleware

import (
	"auth/domain"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(header, skipIp string, storage fiber.Storage, max, expires int64) (f func(ctx *fiber.Ctx) error) {
	f = func(ctx *fiber.Ctx) error {

		userId := ctx.Get(header)
		if userId == "" {
			return ctx.Status(http.StatusBadRequest).JSON(http.StatusText(http.StatusBadRequest))
		}

		return limiter.New(limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.IP() == "500.0.0.1"
			},
			Max:        int(max),
			Expiration: time.Duration(expires) * time.Second,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.Get(header)
			},
			LimitReached: func(c *fiber.Ctx) error {
				response := domain.BasicResponse{
					StatusCode: http.StatusTooManyRequests,
					Message:    http.StatusText(http.StatusTooManyRequests),
				}
				return c.JSON(response)

			},
			Storage: storage,
		})(ctx)
	}

	return
}
