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

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
