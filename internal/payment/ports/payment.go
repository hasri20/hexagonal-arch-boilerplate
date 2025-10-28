package ports

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/payment/models"
)

// PaymentServiceAdapter - port primary
type PaymentServiceAdapter interface {
	TransferUserBalance(ctx context.Context, payload models.TransferBalancePayload) (float64, error)
}

// WalletRepositoryAdapter - port secondary
type WalletRepositoryAdapter interface {
	ReadBalanceInfoFromWallet(ctx context.Context, userID string) (float64, error)
	AppendBalanceInfoIntoWallet(ctx context.Context, userID string, amount float64) error
}
