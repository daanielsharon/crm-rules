package app

import (
	"context"
	"database/sql"
	"fmt"
	"log-worker/config"

	"github.com/redis/go-redis/v9"
)

type App struct {
	Db    *sql.DB
	Redis *redis.Client
}

func New(cfg *config.Config) (*App, error) {
	db, err := initPostgres(cfg.PostgreSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize postgres: %v", err)
	}

	rdb, err := initRedis(cfg.Redis)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to initialize redis: %v", err)
	}

	return &App{
		Db:    db,
		Redis: rdb,
	}, nil
}

func initPostgres(cfg config.PostgresConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	return sql.Open("postgres", connStr)
}

func initRedis(cfg config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return rdb, nil
}

func (a *App) Cleanup() {
	if a.Db != nil {
		a.Db.Close()
	}
	if a.Redis != nil {
		a.Redis.Close()
	}
}
