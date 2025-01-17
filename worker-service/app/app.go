package app

import (
	"database/sql"

	"worker-service/publisher"
	"worker-service/scheduler"
	"worker-service/storage"

	"shared/db"

	"github.com/redis/go-redis/v9"
)

type App struct {
	postgres  *sql.DB
	redis     *redis.Client
	scheduler *scheduler.Scheduler
}

func New() (*App, error) {
	postgres := db.InitPostgres()
	redis := db.InitRedis()

	taskPublisher := publisher.NewPublisher(redis)
	store := storage.NewStorage(postgres)
	scheduler := scheduler.NewScheduler(store, taskPublisher)

	return &App{
		postgres:  postgres,
		redis:     redis,
		scheduler: scheduler,
	}, nil
}

func (a *App) Start() {
	a.scheduler.Start()
}
