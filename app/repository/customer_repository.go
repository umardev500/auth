package repository

import "auth/domain"

type customerRepository struct{}

func NewCustomerRepository() domain.CustomerRepository {
	return &customerRepository{}
}
