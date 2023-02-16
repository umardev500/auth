package repository

import (
	"auth/domain"
	"auth/pb"
	"context"
)

type customerRepository struct {
	customer pb.CustomerServiceClient
}

func NewCustomerRepository(customer pb.CustomerServiceClient) domain.CustomerRepository {
	return &customerRepository{
		customer: customer,
	}
}

func (c *customerRepository) Login(ctx context.Context, req *pb.CustomerLoginRequest) (res domain.LoginResponse, err error) {
	resp, err := c.customer.Login(ctx, req)
	if err != nil {
		return
	}

	isEmpty := resp.IsEmpty
	if isEmpty {
		res = domain.LoginResponse{IsEmpty: true}
		return
	}

	res = domain.LoginResponse{
		Payload: domain.LoginResponseData{
			UserId: resp.Payload.CustomerId,
			User:   resp.Payload.User,
		},
	}

	return
}
