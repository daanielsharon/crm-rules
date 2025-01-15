package main

import (
	"log"
	"log-worker/app"
	"log-worker/config"
	"log-worker/consumer"
	"log-worker/storage"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer application.Cleanup()

	store := &storage.Storage{DB: application.Db}

	log.Println("Log worker running...")
	consumer.StartConsumer(application.Redis, store)
}
