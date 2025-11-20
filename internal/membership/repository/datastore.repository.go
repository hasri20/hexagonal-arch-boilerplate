package repository

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/models"
	user "github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/schema/ent/user_profile"
)

func (repository DatastoreRepository) GetUserSessionFromCache(ctx context.Context) {

}

func (repository DatastoreRepository) UpdateUserSessionIntoCache(ctx context.Context) {

}

func (repository DatastoreRepository) InsertUserInfoIntoDB(ctx context.Context) {

}

func (repository DatastoreRepository) GetUserInfoFromDB(ctx context.Context, accountNumber string) (models.UserProfileInfo, error) {
	user, err := repository.db.User_Profile.Query().Where(user.AccountNumber(accountNumber)).Only(ctx)

	if err != nil {
		return models.UserProfileInfo{}, err
	}

	return models.UserProfileInfo{
		AccountNumber: user.AccountNumber,
		Email:         user.Email,
		Fullname:      user.Fullname,
		Status:        user.Status,
		CreatedAt:     user.CreatedAt,
	}, nil

}

func (repository DatastoreRepository) GetUserByUsername(ctx context.Context, fullname string) (models.UserProfileInfo, error) {
	user, err := repository.db.User_Profile.Query().Where(user.Fullname(fullname)).Only(ctx)

	if err != nil {
		return models.UserProfileInfo{}, err
	}

	return models.UserProfileInfo{
		AccountNumber: user.AccountNumber,
		Email:         user.Email,
		Fullname:      user.Fullname,
		Status:        user.Status,
		CreatedAt:     user.CreatedAt,
	}, nil

}
