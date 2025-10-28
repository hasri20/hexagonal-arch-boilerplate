package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/wallet/handler"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/wallet/repository"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/wallet/services"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
	"github.com/spf13/viper"
)

func initConfig(configName string) (*config.Config, error) {
	cfg := &config.Config{}
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	fmt.Printf("%#v\n", cfg)
	return cfg, err
}

func initHandler(cfg *config.Config, redisClient *redis.Client) (*handler.Handler, error) {
	repo := repository.NewRepository(cfg, redisClient)
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
