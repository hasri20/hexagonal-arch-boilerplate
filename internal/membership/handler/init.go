package handler

import (
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/ports"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
	pbHandler "github.com/hasri20/hexagonal-arch-boilerplate/lib/protos/v1/membership"
)

type Handler struct {
	pbHandler.UnimplementedMembershipServiceServer
	config            *config.Config
	membershipService ports.MembershipServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.MembershipServiceAdapter) (*Handler, error) {
	return &Handler{
		config:            cfg,
		membershipService: adapter,
	}, nil
}
