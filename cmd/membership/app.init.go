package main

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/handler"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/repository"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/schema/ent"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/services"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/datastore/postgre"
)

func initHandler(cfg *config.Config) (*handler.Handler, error) {

	redisClient := InitRedis(cfg)
	dbClient, err := InitDB(cfg)

	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(cfg, redisClient, dbClient)
	service := services.NewService(cfg, repo)
	handler, err := handler.NewHandler(cfg, service)
	if err != nil {
		return nil, err
	}

	return handler, nil
}

func InitRedis(cfg *config.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		PoolSize: cfg.Redis.PoolSize,
	})

	return client
}

func InitDB(cfg *config.Config) (*ent.Client, error) {

	db := postgre.InitConfig(cfg)
	client, err := db.InitConnection()

	if err != nil {
		log.Fatalf("failed opening connection to postgre: %v", err)
	}

	return client, nil
}
