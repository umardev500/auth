package injector

import (
	"auth/app/delivery"

	"github.com/gofiber/fiber/v2"
)

func NewAdminInjector(router fiber.Router) {
	delivery.NewAdminDelivery(router)
}
