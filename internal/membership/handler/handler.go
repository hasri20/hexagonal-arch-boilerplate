package handler

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/models"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/protos/v1/membership"
)

// GetUserInfo get user info
func (handler *Handler) GetUserInfo(ctx context.Context, req *membership.UserInfoRequest) (*membership.UserInfoResponse, error) {
	user, err := handler.membershipService.GetUserInfo(ctx, req.GetAccountNumber())
	if err != nil {
		return nil, err
	}

	return &membership.UserInfoResponse{
		Email:         user.Email,
		Fullname:      user.Fullname,
		AccountNumber: user.AccountNumber,
	}, nil
}

// SubmitLogin login
func (handler *Handler) SubmitLogin(ctx context.Context, req *membership.LoginRequest) (*membership.LoginResponse, error) {
	res, err := handler.membershipService.SubmitLogin(ctx, models.LoginInfo{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return &membership.LoginResponse{
		Success:      res.Success,
		LoginMessage: res.Message,
		Uuid:         res.UUID,
	}, nil
}

// SubmitLogout logout
func (handler *Handler) SubmitLogout(ctx context.Context, req *membership.LogoutRequest) (*membership.LogoutResponse, error) {
	handler.membershipService.SubmitLogout(ctx)
	return nil, nil
}

// SubmitRegistration register
func (handler *Handler) SubmitRegistration(ctx context.Context, req *membership.RegistrationRequest) (*membership.RegistrationResponse, error) {
	return nil, nil
}
