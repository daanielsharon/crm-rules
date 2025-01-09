package publisher

import "time"

type Task struct {
	RuleID    string   `json:"rule_id"`
	Name      string   `json:"name"`
	Action    []string `json:"action"`
	Condition string   `json:"condition"`
	Timestamp string   `json:"timestamp"`
}

func NewTask(ruleID, name, condition string, action []string) Task {
	return Task{
		RuleID:    ruleID,
		Name:      name,
		Action:    action,
		Condition: condition,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
