package services

import (
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/ports"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
)

type MembershipService struct {
	config     *config.Config
	repository ports.DatastoreRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.DatastoreRepositoryAdapter) ports.MembershipServiceAdapter {
	return &MembershipService{
		config:     cfg,
		repository: repository,
	}
}
