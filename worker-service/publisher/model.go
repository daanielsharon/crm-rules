package publisher

import (
	"time"
)

type Task struct {
	RuleID    int    `json:"rule_id"`
	Name      string `json:"name"`
	Action    string `json:"action"`
	Condition string `json:"condition"`
	Timestamp string `json:"timestamp"`
}

func NewTask(ruleID int, name, condition, action string) Task {
	return Task{
		RuleID:    ruleID,
		Name:      name,
		Action:    action,
		Condition: condition,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
