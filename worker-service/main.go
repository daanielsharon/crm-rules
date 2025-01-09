package main

import (
	"log"

	"worker-service/app"
	"worker-service/config"

	_ "github.com/lib/pq" // PostgreSQL driver
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

	log.Println("Starting worker service...")
	application.Start()

	// Block indefinitely to keep the process running
	select {}
}
