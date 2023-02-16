package usecase

import "auth/domain"

type customerUsecase struct {
	respository domain.CustomerRepository
}

func NewCustomerUsecase(respository domain.CustomerRepository) domain.CustomerUsecase {
	return &customerUsecase{
		respository: respository,
	}
}
