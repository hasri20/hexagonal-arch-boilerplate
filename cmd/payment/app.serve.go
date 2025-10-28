package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	http "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/config"
	payment "github.com/titikterang/hexagonal-arch-boilerplate/lib/protos/v1/payment"
)

func startService(cfg *config.Config) {
	handler, err := initHandler(cfg)
	if err != nil {
		log.Fatal("failed initiate NewHandler: %v", err)
	}

	grpcOpts := []grpc.ServerOption{
		grpc.Timeout(cfg.Http.Timeout),
		grpc.Address(cfg.App.GRPCAddress),
		grpc.Middleware(
			metadata.Server(),
			logging.Server(log.GetLogger()),
		),
		grpc.Logger(log.GetLogger()),
	}

	grpcServer := grpc.NewServer(
		grpcOpts...,
	)

	payment.RegisterPaymentServiceServer(grpcServer, handler)

	httpOpts := []http.ServerOption{
		http.Timeout(cfg.Http.Timeout),
		http.Address(cfg.App.HTTPAddress),
		http.Filter(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}),
			),
		),
		http.Middleware(
			metadata.Server(),
			logging.Server(log.GetLogger()),
		),
		http.Logger(log.GetLogger()),
	}

	httpServer := http.NewServer(
		httpOpts...,
	)

	payment.RegisterPaymentServiceHTTPServer(httpServer, handler)

	server := kratos.New(
		kratos.Name(cfg.App.Label),
		kratos.Server(
			httpServer,
			grpcServer,
		),
	)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
