package repository

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	grpc "github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/titikterang/hexagonal-arch-boilerplate/internal/payment/ports"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/config"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/protos/v1/wallet"
)

type PaymentRepository struct {
	Client wallet.WalletServiceClient
	cfg    *config.Config
}

func NewRepository(cfg *config.Config) ports.WalletRepositoryAdapter {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:3002"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return &PaymentRepository{
		cfg:    cfg,
		Client: wallet.NewWalletServiceClient(conn),
	}
}
