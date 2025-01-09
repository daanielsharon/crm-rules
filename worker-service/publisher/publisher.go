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
	Channel     string
}

func NewPublisher(redisClient *redis.Client, channel string) *Publisher {
	return &Publisher{
		RedisClient: redisClient,
		Channel:     channel,
	}
}

func (p *Publisher) PublishTask(task Task) error {
	taskJSON, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = p.RedisClient.Publish(ctx, p.Channel, taskJSON).Err()
	if err != nil {
		return err
	}

	log.Printf("Task published to channel %s: %+v\n", p.Channel, task)
	return nil
}
