package handler

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/payment/models"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/protos/v1/payment"
)

func (handler *Handler) TransferBalanceService(ctx context.Context, in *payment.TransferBalanceRequest) (*payment.TransferBalanceResponse, error) {
	amount, err := handler.paymentService.TransferUserBalance(ctx, models.TransferBalancePayload{
		SourceUserID: in.GetSourceUserId(),
		TargetUserID: in.GetDestination(),
		Amount:       in.GetAmount(),
	})

	return &payment.TransferBalanceResponse{
		Success:           err == nil,
		DestinationAmount: amount,
	}, nil
}
