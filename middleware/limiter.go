package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(storage fiber.Storage) (f func(ctx *fiber.Ctx) error) {
	f = func(ctx *fiber.Ctx) error {

		userId := ctx.Get("userid")
		if userId == "" {
			return ctx.Status(http.StatusBadRequest).JSON(http.StatusText(http.StatusBadRequest))
		}

		expiredTime, _ := strconv.Atoi(os.Getenv("EXPIRED_HOURS"))

		return limiter.New(limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.IP() == "500.0.0.1"
			},
			Max:        1,
			Expiration: time.Duration(expiredTime) * time.Minute,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.Get("userid")
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.JSON(http.StatusText(http.StatusTooManyRequests))
			},
			Storage: storage,
		})(ctx)
	}

	return
}
