package consumer

import (
	"context"
	"encoding/json"
	"log"
	"log-worker/models"
	"log-worker/storage"
	"shared/helpers"

	"github.com/redis/go-redis/v9"
)

func StartConsumer(redisClient *redis.Client, store storage.LogStorageInterface) {
	ctx := context.Background()
	subscriber := redisClient.Subscribe(ctx, "logs")

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}

		var logEntry models.Log
		if err := json.Unmarshal([]byte(msg.Payload), &logEntry); err != nil {
			log.Printf("Error unmarshaling log entry: %v", err)
			continue
		}

		log.Printf("Processing log: %+v", logEntry)

		err = store.CreateLog(logEntry)
		helpers.PanicIfError(err)
	}
}
