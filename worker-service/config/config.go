package config

import "os"

type Config struct {
	PostgreSQL PostgresConfig
	Redis      RedisConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Addr    string
	Channel string
}

func Load() (*Config, error) {
	return &Config{
		PostgreSQL: PostgresConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
		},
		Redis: RedisConfig{
			Addr:    os.Getenv("REDIS_ADDR"),
			Channel: os.Getenv("REDIS_CHANNEL"),
		},
	}, nil
}
