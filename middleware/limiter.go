package middleware

import (
	"auth/domain"
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

		maxReq, _ := strconv.Atoi(os.Getenv("LOGIN_MAX_REQ"))
		expiredTime, _ := strconv.Atoi(os.Getenv("LOGIN_EXPIRED_TIME"))

		return limiter.New(limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.IP() == "500.0.0.1"
			},
			Max:        maxReq,
			Expiration: time.Duration(expiredTime) * time.Second,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.Get("userid")
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
