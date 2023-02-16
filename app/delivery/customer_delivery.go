package delivery

import "github.com/gofiber/fiber/v2"

type customerDelivery struct{}

func NewCustomerDeliery(router fiber.Router) {
	handler := &customerDelivery{}
	router.Post("/login", handler.Login)
}

func (c *customerDelivery) Login(ctx *fiber.Ctx) error {
	return ctx.JSON("customer login")
}
