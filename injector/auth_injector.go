package injector

import (
	"auth/app/delivery"

	"github.com/gofiber/fiber/v2"
)

func NewAuthInjector(router fiber.Router) {
	delivery.NewAuthDelivery(router)
}
