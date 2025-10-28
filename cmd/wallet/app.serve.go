package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	http "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
	wallet "github.com/hasri20/hexagonal-arch-boilerplate/lib/protos/v1/wallet"
)

func startService(cfg *config.Config) {
	redisClient := InitRedis(cfg)
	defer func() {
		err := redisClient.Close()
		if err != nil {
			log.Fatal("failed to close redis client: %v", err)
		}
	}()

	handler, err := initHandler(cfg, redisClient)
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

	wallet.RegisterWalletServiceServer(grpcServer, handler)

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

	wallet.RegisterWalletServiceHTTPServer(httpServer, handler)

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
