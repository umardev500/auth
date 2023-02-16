package usecase

import "auth/domain"

type customerUsecase struct{}

func NewCustomerUsecase() domain.CustomerUsecase {
	return &customerUsecase{}
}
