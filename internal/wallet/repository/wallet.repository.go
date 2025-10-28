package repository

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/wallet/models"
)

func (repository *WalletRepository) ReadBalanceInfoFromDatastore(ctx context.Context, userID string) (models.DatastoreBalanceResponse, error) {
	// hgetall from redis
	var respData = models.DatastoreBalanceResponse{
		List: make([]string, 0),
	}
	rdsKey := "user:balance:" + userID
	log.Infof("redis key : %s", rdsKey)
	resp, err := repository.client.HGetAll(ctx, rdsKey).Result()
	if err != nil {
		return models.DatastoreBalanceResponse{}, err
	}
	for _, v := range resp {
		respData.List = append(respData.List, v)
	}

	return respData, nil
}

func (repository *WalletRepository) AppendBalanceInfoIntoDatastore(ctx context.Context, userID string, amount float64) error {
	rdsKey := "user:balance:" + userID
	rdsHash := time.Now().UnixMilli()
	_, err := repository.client.HSet(ctx, rdsKey, rdsHash, amount).Result()
	return err
}
