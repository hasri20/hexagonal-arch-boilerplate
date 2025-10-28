package main

import (
	"fmt"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/payment/handler"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/payment/repository"
	"github.com/hasri20/hexagonal-arch-boilerplate/internal/payment/services"
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

func initHandler(cfg *config.Config) (*handler.Handler, error) {
	repo := repository.NewRepository(cfg)
	service := services.NewService(cfg, repo)
	handler, err := handler.NewHandler(cfg, service)
	if err != nil {
		return nil, err
	}

	return handler, nil
}
