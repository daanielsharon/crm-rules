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
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	return &Config{
		Redis: RedisConfig{
			Addr: redisAddr,
		},
		Postgres: PostgresConfig{
			URL: dbURL,
		},
	}, nil
}
