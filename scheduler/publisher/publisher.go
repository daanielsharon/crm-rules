package publisher

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Publisher struct {
	RedisClient *redis.Client
}

func NewPublisher(redisClient *redis.Client) *Publisher {
	return &Publisher{
		RedisClient: redisClient,
	}
}

func (p *Publisher) PublishTask(task Task) error {
	taskJSON, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = p.RedisClient.Publish(ctx, "tasks", taskJSON).Err()
	if err != nil {
		return err
	}

	log.Printf("Task published to channel %s: %+v\n", "tasks", task)
	return nil
}
