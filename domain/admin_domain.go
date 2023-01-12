package domain

import (
	"auth/pb"
	"context"
)

type AdminRepository interface {
	Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error)
}

type AdminUsecase interface {
	Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error)
}
