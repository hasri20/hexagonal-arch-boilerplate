package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/wallet/ports"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
)

type WalletRepository struct {
	config *config.Config
	client *redis.Client
}

func NewRepository(cfg *config.Config, client *redis.Client) ports.WalletRepositoryAdapter {
	return &WalletRepository{
		config: cfg,
		client: client,
	}
}
