package postgres

import (
	"context"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectPool(dbURL string, logLevel LogLevel) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
	config.ConnConfig.LogLevel = pgx.LogLevel(logLevel)

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	return db, err
}
