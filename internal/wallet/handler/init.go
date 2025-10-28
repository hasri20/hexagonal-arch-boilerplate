package handler

import (
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/wallet/ports"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
	pbHandler "github.com/hasri20/hexagonal-arch-boilerplate/lib/protos/v1/wallet"
)

type Handler struct {
	pbHandler.UnimplementedWalletServiceServer
	config        *config.Config
	walletService ports.WalletServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.WalletServiceAdapter) (*Handler, error) {
	return &Handler{
		config:        cfg,
		walletService: adapter,
	}, nil
}
