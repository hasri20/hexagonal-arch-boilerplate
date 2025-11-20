package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/ports"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/schema/ent"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
)

type DatastoreRepository struct {
	config *config.Config
	client *redis.Client
	db     *ent.Client
}

func NewRepository(cfg *config.Config, client *redis.Client, db *ent.Client) ports.DatastoreRepositoryAdapter {

	return &DatastoreRepository{
		config: cfg,
		client: client,
		db:     db,
	}
}
