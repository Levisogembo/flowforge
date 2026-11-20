package database

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Levisogembo/flowforge/internal/config"
)

func New(cfg *config.Config) (*pgxpool.Pool, error) {
	// Build the PostgreSQL connection string.
	connURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DB.User, cfg.DB.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.DB.Host, cfg.DB.Port),
		Path:   cfg.DB.Name,
	}

	query := connURL.Query()
	query.Set("sslmode", cfg.DB.SSLMode)
	connURL.RawQuery = query.Encode()

	// Parse the connection string into a pgxpool configuration.
	poolConfig, err := pgxpool.ParseConfig(connURL.String())
	if err != nil {
		return nil, fmt.Errorf("parse database config: %w", err)
	}

	// Apply connection pool settings.
	poolConfig.MinConns = cfg.DB.MinConns
	poolConfig.MaxConns = cfg.DB.MaxConns
	poolConfig.MaxConnIdleTime = cfg.DB.MaxIdleTime

	// Create the connection pool.
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}

	// Verify that PostgreSQL is reachable.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	return pool, nil
}
