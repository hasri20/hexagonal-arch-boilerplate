package services

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/models"
)

func (service *MembershipService) SubmitRegisterUser(ctx context.Context) error {
	return nil
}

func (service *MembershipService) GetUserInfo(ctx context.Context, accountNumber string) (models.UserProfileInfo, error) {
	return service.repository.GetUserInfoFromDB(ctx, accountNumber)
}

func (service *MembershipService) SubmitLogin(ctx context.Context, loginInfo models.LoginInfo) (models.LoginResponse, error) {
	return models.LoginResponse{}, nil
}

func (service *MembershipService) SubmitLogout(ctx context.Context) error {
	return nil
}
