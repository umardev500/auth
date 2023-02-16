package injector

import (
	"auth/app/delivery"
	"auth/pb"

	"github.com/gofiber/fiber/v2"
)

func NewCustomerInjector(router fiber.Router, customer pb.CustomerServiceClient) {
	delivery.NewCustomerDeliery(router)
}
