package ports

import (
	"context"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/payment/models"
)

// PaymentServiceAdapter - port primary
type PaymentServiceAdapter interface {
	TransferUserBalance(ctx context.Context, payload models.TransferBalancePayload) (float64, error)
}
