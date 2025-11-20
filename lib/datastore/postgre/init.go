package postgre

import (
	"context"
	"fmt"
	"log"

	"github.com/hasri20/hexagonal-arch-boilerplate/internal/membership/schema/ent"
	"github.com/hasri20/hexagonal-arch-boilerplate/lib/config"
	_ "github.com/lib/pq"
)

type Database struct {
	config *config.Config
	client *ent.Client
}

func InitConfig(config *config.Config) *Database {
	return &Database{
		config: config,
	}
}

func (db *Database) InitConnection() (*ent.Client, error) {
	config := db.config.PostgreConfig

	conn, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Address,
		config.Port,
		config.Username,
		config.Password,
		config.DBName,
		config.SSLMode,
	))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	if err := conn.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return conn, nil
}
