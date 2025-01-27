package consumer

import (
	"fmt"
	"log"
	"task-execution-service/publisher"
	"task-execution-service/storage"
	"task-execution-service/types"
)

type TaskProcessor struct {
	storage   storage.Storage
	publisher publisher.LogPublisher
}

func NewTaskProcessor(storage storage.Storage, publisher publisher.LogPublisher) *TaskProcessor {
	return &TaskProcessor{
		storage:   storage,
		publisher: publisher,
	}
}

func (tp *TaskProcessor) ProcessTask(task types.Task) error {
	rows, err := tp.storage.GetMatchingUsers(task.Condition)
	if err != nil {
		return fmt.Errorf("get matching users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user storage.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Printf("Failed to scan user row: %v", err)
			continue
		}

		if err := tp.processUserActions(task, user); err != nil {
			log.Printf("Failed to process actions for user %s: %v", user.ID, err)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration: %w", err)
	}

	return nil
}

func (tp *TaskProcessor) processUserActions(task types.Task, user storage.User) error {
	status := tp.executeAction(user, task.Action)

	message := types.NewLog(task.RuleID, user.ID, task.Action, status)
	err := tp.publisher.PublishLogs(message)

	if err != nil {
		return fmt.Errorf("log execution: %w", err)
	}

	log.Printf("Log published to channel %s: %+v\n", "logs", message)

	return nil
}

func (tp *TaskProcessor) executeAction(user storage.User, action string) string {
	log.Printf("Executing action '%s' for user %s (%s)", action, user.Name, user.Email)
	return "success"
}
