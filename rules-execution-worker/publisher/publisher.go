package publisher

import (
	"context"
	"encoding/json"
	"log"
	"task-execution-service/types"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type LogPublisher interface {
	PublishLogs(message types.Log) error
}

type Publisher struct {
	RedisClient *redis.Client
}

func NewPublisher(redisClient *redis.Client) LogPublisher {
	return &Publisher{
		RedisClient: redisClient,
	}
}

func (p *Publisher) PublishLogs(message types.Log) error {
	taskJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.RedisClient.Publish(ctx, "logs", taskJSON).Err()
	if err != nil {
		return err
	}

	log.Printf("Task published to channel %s: %+v\n", "logs", message)
	return nil
}
