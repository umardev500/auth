package usecase

import (
	"auth/domain"
	"auth/pb"
	"context"
)

type customerUsecase struct {
	respository domain.CustomerRepository
}

func NewCustomerUsecase(respository domain.CustomerRepository) domain.CustomerUsecase {
	return &customerUsecase{
		respository: respository,
	}
}

func (c *customerUsecase) Login(ctx context.Context, req domain.LoginRequest) (res domain.LoginResponse, err error) {
	res, err = c.respository.Login(ctx, &pb.CustomerLoginRequest{User: req.Username, Pass: req.Password})

	return
}
