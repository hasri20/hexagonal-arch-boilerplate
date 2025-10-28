package services

import (
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/payment/ports"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
)

type PaymentService struct {
	config     *config.Config
	repository ports.WalletRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.WalletRepositoryAdapter) ports.PaymentServiceAdapter {
	return &PaymentService{
		config:     cfg,
		repository: repository,
	}
}
