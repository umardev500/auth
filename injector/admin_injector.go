package injector

import (
	"auth/app/delivery"
	"auth/app/repository"
	"auth/app/usecase"
	"auth/pb"

	"github.com/gofiber/fiber/v2"
)

func NewAdminInjector(router fiber.Router, user pb.UserServiceClient) {
	ar := repository.NewAdminRepository(user)
	au := usecase.NewAdminUsecase(ar)
	strg := NewRedisInjector()
	delivery.NewAdminDelivery(router, strg, au)
}
