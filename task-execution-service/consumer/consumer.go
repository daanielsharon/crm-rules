package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"task-execution-service/types"

	"github.com/redis/go-redis/v9"
)

type Processor interface {
	ProcessTask(task types.Task) error
}

type Consumer struct {
	client    *redis.Client
	processor Processor
}

func New(client *redis.Client, processor Processor) *Consumer {
	return &Consumer{
		client:    client,
		processor: processor,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	sub := c.client.Subscribe(ctx, "tasks")
	defer sub.Close()

	log.Printf("Subscribed to Redis channel: %s", "tasks")

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			msg, err := sub.ReceiveMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return nil
				}
				return fmt.Errorf("receive message: %w", err)
			}

			if err := c.handleMessage(msg); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}
}

func (c *Consumer) handleMessage(msg *redis.Message) error {
	var task types.Task
	if err := json.Unmarshal([]byte(msg.Payload), &task); err != nil {
		return fmt.Errorf("unmarshal task: %w", err)
	}

	log.Printf("Processing task: %+v", task)

	if err := c.processor.ProcessTask(task); err != nil {
		return fmt.Errorf("process task: %w", err)
	}

	return nil
}
