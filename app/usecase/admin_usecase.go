package usecase

import (
	"auth/domain"
	"auth/pb"
	"context"
)

type adminUsecase struct {
	repository domain.AdminRepository
}

func NewAdminUsecase(repository domain.AdminRepository) domain.AdminUsecase {
	return &adminUsecase{
		repository: repository,
	}
}

func (a *adminUsecase) Login(ctx context.Context, req *pb.UserLoginRequest) (resp *pb.UserLoginResponse, err error) {
	resp, err = a.repository.Login(ctx, req)

	return
}
