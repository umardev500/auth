package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type adminDelivery struct{}

func NewAdminDelivery(router fiber.Router) {
	handler := &adminDelivery{}
	router.Get("/admin", handler.Auth)
	router.Post("/admin", handler.Login)
}

func (a *adminDelivery) Login(ctx *fiber.Ctx) error {
	return ctx.JSON("login")
}

func (a *adminDelivery) Auth(ctx *fiber.Ctx) error {
	// token := ctx.GetReqHeaders()
	// fmt.Println(token)

	return ctx.Status(http.StatusUnauthorized).JSON(http.StatusText(http.StatusUnauthorized))
}
