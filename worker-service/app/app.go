package app

import (
	"context"
	"database/sql"
	"fmt"

	"worker-service/config"
	"worker-service/publisher"
	"worker-service/scheduler"
	"worker-service/storage"

	"github.com/redis/go-redis/v9"
)

type App struct {
	db        *sql.DB
	redis     *redis.Client
	scheduler *scheduler.Scheduler
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

	taskPublisher := publisher.NewPublisher(rdb, cfg.Redis.Channel)
	store := storage.NewStorage(db)
	scheduler := scheduler.NewScheduler(store, taskPublisher)

	return &App{
		db:        db,
		redis:     rdb,
		scheduler: scheduler,
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

func (a *App) Start() {
	a.scheduler.Start()
}

func (a *App) Cleanup() {
	if a.db != nil {
		a.db.Close()
	}
	if a.redis != nil {
		a.redis.Close()
	}
}
