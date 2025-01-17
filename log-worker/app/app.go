package app

import (
	"log-worker/consumer"
	"log-worker/storage"
	"shared/db"
)

func New() {
	postgres := db.InitPostgres()
	redis := db.InitRedis()

	store := storage.NewStorage(postgres)
	consumer.StartConsumer(redis, store)
}
