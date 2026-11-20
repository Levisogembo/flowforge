package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env  string
	Addr string

	DB struct {
		Host        string
		Port        string
		User        string
		Password    string
		Name        string
		SSLMode     string
		MinConns    int32
		MaxConns    int32
		MaxIdleTime time.Duration
	}
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Env:  getEnv("APP_ENV", "development"),
		Addr: getEnv("APP_ADDR", ":8080"),
	}
	cfg.DB.Host = getEnv("DB_HOST", "localhost")
	cfg.DB.Port = getEnv("DB_PORT", "5432")
	cfg.DB.User = getEnv("DB_USER", "flowforge")
	cfg.DB.Password = getEnv("DB_PASSWORD", "")
	cfg.DB.Name = getEnv("DB_NAME", "flowforge")
	cfg.DB.SSLMode = getEnv("DB_SSLMODE", "disable")

	minConns, err := strconv.ParseInt(
		getEnv("DB_MIN_CONNS", "2"), 10, 32,
	)
	if err != nil || minConns < 0 {
		return nil, fmt.Errorf("invalid DB_MIN_CONNS")
	}

	maxConns, err := strconv.ParseInt(
		getEnv("DB_MAX_CONNS", "25"), 10, 32,
	)
	if err != nil || maxConns <= 0 || maxConns < minConns {
		return nil, fmt.Errorf("invalid DB_MAX_CONNS")
	}

	maxIdleTime, err := time.ParseDuration(
		getEnv("DB_MAX_IDLE_TIME", "15m"),
	)
	if err != nil || maxIdleTime <= 0 {
		return nil, fmt.Errorf("invalid DB_MAX_IDLE_TIME")
	}

	cfg.DB.MinConns = int32(minConns)
	cfg.DB.MaxConns = int32(maxConns)
	cfg.DB.MaxIdleTime = maxIdleTime

	if cfg.DB.Password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return fallback
	}

	return value
}
