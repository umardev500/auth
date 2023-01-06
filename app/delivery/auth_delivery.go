package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type authDelivery struct{}

func NewAuthDelivery(router fiber.Router) {
	handler := &authDelivery{}
	router.Get("/unauthorized", handler.Unauthorized)
}

func (a *authDelivery) Unauthorized(ctx *fiber.Ctx) error {
	status := http.StatusUnauthorized
	return ctx.Status(status).JSON(http.StatusText(status))
}
