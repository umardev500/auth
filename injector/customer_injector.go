package injector

import (
	"auth/app/delivery"
	"auth/app/repository"
	"auth/app/usecase"
	"auth/pb"

	"github.com/gofiber/fiber/v2"
)

func NewCustomerInjector(router fiber.Router, customer pb.CustomerServiceClient) {
	repo := repository.NewCustomerRepository(customer)
	uc := usecase.NewCustomerUsecase(repo)
	delivery.NewCustomerDeliery(router, uc)
}
