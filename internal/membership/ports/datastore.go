package ports

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/models"
)

type DatastoreRepositoryAdapter interface {
	//redis
	GetUserSessionFromCache(ctx context.Context)
	UpdateUserSessionIntoCache(ctx context.Context)

	//postgre
	InsertUserInfoIntoDB(ctx context.Context)
	GetUserInfoFromDB(ctx context.Context, accountNumber string) (models.UserProfileInfo, error)
	GetUserByUsername(ctx context.Context, fullname string) (models.UserProfileInfo, error)
}
