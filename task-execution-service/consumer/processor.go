package consumer

import (
	"fmt"
	"log"
	"task-execution-service/storage"
	"task-execution-service/types"
)

type TaskProcessor struct {
	storage storage.Store
}

func NewTaskProcessor(storage storage.Store) *TaskProcessor {
	return &TaskProcessor{storage: storage}
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
	for _, action := range task.Actions {
		status := tp.executeAction(user, action)

		if err := tp.storage.LogExecution(task.RuleID, user.ID, action, status); err != nil {
			return fmt.Errorf("log execution: %w", err)
		}
	}
	return nil
}

func (tp *TaskProcessor) executeAction(user storage.User, action string) string {
	log.Printf("Executing action '%s' for user %s (%s)", action, user.Name, user.Email)
	return "success"
}
