package services

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/payment/models"
)

func (service *PaymentService) TransferUserBalance(ctx context.Context, payload models.TransferBalancePayload) (float64, error) {
	// deduct source user balance
	service.repository.AppendBalanceInfoIntoWallet(ctx, payload.SourceUserID, payload.Amount*-1)
	// add destination user balance
	service.repository.AppendBalanceInfoIntoWallet(ctx, payload.TargetUserID, payload.Amount)
	// get final destination user balance
	resp, _ := service.repository.ReadBalanceInfoFromWallet(ctx, payload.TargetUserID)

	return resp, nil
}
