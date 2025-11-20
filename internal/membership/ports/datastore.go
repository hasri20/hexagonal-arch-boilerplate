package ports

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/models"
)

// // CacheRepositoryAdapter - port secondary
// type CacheRepositoryAdapter interface {
// 	GetUserSessionFromCache(ctx context.Context) error
// 	UpdateUserSessionIntoCache(ctx context.Context) error
// }

// // DatabaseRepositoryAdapter - port secondary
// type DatabaseRepositoryAdapter interface {
// 	InsertUserInfoIntoDB(ctx context.Context) error
// 	GetUserInfoFromDB(ctx context.Context) error
// 	GetUserByUsername(ctx context.Context) error
// }

type DatastoreRepositoryAdapter interface {
	//redis
	GetUserSessionFromCache(ctx context.Context)
	UpdateUserSessionIntoCache(ctx context.Context)

	//postgre
	InsertUserInfoIntoDB(ctx context.Context)
	GetUserInfoFromDB(ctx context.Context, accountNumber string) (models.UserProfileInfo, error)
	GetUserByUsername(ctx context.Context, fullname string) (models.UserProfileInfo, error)
}
