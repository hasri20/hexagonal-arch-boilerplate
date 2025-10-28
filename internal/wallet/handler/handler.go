package handler

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/wallet/models"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/protos/v1/wallet"
)

func (handler *Handler) GetUserBalance(ctx context.Context, in *wallet.GetBalanceRequest) (*wallet.GetBalanceResponse, error) {
	resp, err := handler.walletService.GetUserBalance(ctx, in.GetUserId())
	if err != nil {
		/// log err
		return nil, err
	}
	data := &wallet.GetBalanceResponse{
		UserId:  in.UserId,
		Balance: resp.AvailableBalance,
	}

	return data, nil
}

func (handler *Handler) UpdateUserBalance(ctx context.Context, in *wallet.UpdateBalanceRequest) (*wallet.UpdateBalanceResponse, error) {
	var message = "success"
	amount, err := handler.walletService.UpdateUserBalance(ctx, models.UpdateBalancePayload{
		UserID: in.GetUserId(),
		Amount: in.GetAmount(),
	})
	if err != nil {
		// log err
		message = err.Error()
	}

	return &wallet.UpdateBalanceResponse{
		Message:      message,
		Success:      err == nil,
		FinalBalance: amount,
	}, nil
}
