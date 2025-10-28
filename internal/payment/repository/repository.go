package repository

import (
	"context"
	"log"

	"github.com/titikterang/hexagonal-arch-boilerplate/lib/protos/v1/wallet"
)

func (repository *PaymentRepository) ReadBalanceInfoFromWallet(ctx context.Context, userID string) (float64, error) {

	payload := wallet.GetBalanceRequest{
		UserId: userID,
	}

	reply, err := repository.Client.GetUserBalance(ctx, &payload)
	if err != nil {
		log.Fatal(err)
	}

	return float64(reply.Balance), err
}

func (repository *PaymentRepository) AppendBalanceInfoIntoWallet(ctx context.Context, userID string, amount float64) error {
	payload := wallet.UpdateBalanceRequest{
		UserId: userID,
		Amount: amount,
	}

	_, err := repository.Client.UpdateUserBalance(ctx, &payload)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
