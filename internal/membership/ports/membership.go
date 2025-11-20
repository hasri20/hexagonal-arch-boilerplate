package ports

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/models"
)

// MembershipServiceAdapter - port primary
type MembershipServiceAdapter interface {
	SubmitRegisterUser(ctx context.Context) error
	GetUserInfo(ctx context.Context, accountNumber string) (models.UserProfileInfo, error)
	SubmitLogin(ctx context.Context, loginInfo models.LoginInfo) (models.LoginResponse, error)
	SubmitLogout(ctx context.Context) error
}
