package services

import (
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/wallet/ports"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
)

type WalletService struct {
	config     *config.Config
	repository ports.WalletRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.WalletRepositoryAdapter) ports.WalletServiceAdapter {
	return &WalletService{
		config:     cfg,
		repository: repository,
	}
}
