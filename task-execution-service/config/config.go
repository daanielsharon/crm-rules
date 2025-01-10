package config

import (
	"fmt"
	"os"
)

type Config struct {
	Redis    RedisConfig
	Postgres PostgresConfig
}

type RedisConfig struct {
	Addr    string
	Channel string
}

type PostgresConfig struct {
	URL string
}

func Load() (*Config, error) {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		return nil, fmt.Errorf("REDIS_ADDR environment variable is required")
	}

	redisChannel := os.Getenv("REDIS_CHANNEL")
	if redisChannel == "" {
		return nil, fmt.Errorf("REDIS_CHANNEL environment variable is required")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required")
	}

	return &Config{
		Redis: RedisConfig{
			Addr:    redisAddr,
			Channel: redisChannel,
		},
		Postgres: PostgresConfig{
			URL: dbURL,
		},
	}, nil
}
