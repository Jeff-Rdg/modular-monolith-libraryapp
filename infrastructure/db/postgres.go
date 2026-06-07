package db

import (
	"context"
	"modular-monolith-libraryApp/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPool(cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
