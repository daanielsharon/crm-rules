package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"shared/db"
	"syscall"
	"task-execution-service/consumer"
	"task-execution-service/publisher"
	"task-execution-service/storage"
)

func main() {
	// Setup context for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	redis := db.InitRedis()
	postgres := db.InitPostgres()
	store := storage.New(postgres)
	publisher := publisher.NewPublisher(redis)
	processor := consumer.NewTaskProcessor(store, *publisher)
	consumer := consumer.New(redis, processor)

	if err := runService(ctx, consumer); err != nil {
		log.Fatalf("Service error: %v", err)
	}
}

func runService(ctx context.Context, consumer *consumer.Consumer) error {
	errChan := make(chan error, 1)

	go func() {
		if err := consumer.Start(ctx); err != nil {
			errChan <- fmt.Errorf("consumer error: %v", err)
		}
	}()

	log.Println("Task Execution Service is running...")

	select {
	case <-ctx.Done():
		log.Println("Shutting down Task Execution Service...")
		return nil
	case err := <-errChan:
		return err
	}
}
