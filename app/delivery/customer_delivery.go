package delivery

import (
	"auth/domain"

	"github.com/gofiber/fiber/v2"
)

type customerDelivery struct {
	usecase domain.CustomerUsecase
}

func NewCustomerDeliery(router fiber.Router, usecase domain.CustomerUsecase) {
	handler := &customerDelivery{
		usecase: usecase,
	}
	router.Post("/login", handler.Login)
}

func (c *customerDelivery) Login(ctx *fiber.Ctx) error {
	return ctx.JSON("customer login")
}
