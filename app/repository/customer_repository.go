package repository

import (
	"auth/domain"
	"auth/pb"
)

type customerRepository struct {
	customer pb.CustomerServiceClient
}

func NewCustomerRepository(customer pb.CustomerServiceClient) domain.CustomerRepository {
	return &customerRepository{
		customer: customer,
	}
}
