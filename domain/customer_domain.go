package domain

import (
	"auth/pb"
	"context"
)

type CustomerUsecase interface {
	Login(ctx context.Context, req LoginRequest) (res LoginResponse, err error)
}

type CustomerRepository interface {
	Login(ctx context.Context, req *pb.CustomerLoginRequest) (res LoginResponse, err error)
}
