package main

import (
	"os"

	"github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
	zlog "github.com/rs/zerolog"
)

func main() {
	zlogT := zlog.New(os.Stdout).With().CallerWithSkipFrameCount(4).Timestamp().Logger()
	logger := zerolog.NewLogger(&zlogT)
	log.SetLogger(logger)

	log.Info("Starting backend service")

	cfg, err := config.InitConfig("config.membership", "./files/config")
	if err != nil {
		log.Fatalf("failed to parse config, %#v", err)
	}

	startService(cfg)
}
