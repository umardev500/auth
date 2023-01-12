package repository

import (
	"auth/domain"
	"auth/pb"
	"context"
)

type adminRepository struct {
	user pb.UserServiceClient
}

func NewAdminRepository(user pb.UserServiceClient) domain.AdminRepository {
	return &adminRepository{
		user: user,
	}
}

func (a *adminRepository) Login(ctx context.Context, req *pb.UserLoginRequest) (resp *pb.UserLoginResponse, err error) {
	resp, err = a.user.Login(ctx, req)

	return
}
