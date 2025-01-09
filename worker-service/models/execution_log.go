package models

import "time"

type ExecutionLog struct {
	ID         string    `json:"id"`
	RuleID     string    `json:"rule_id"`
	ExecutedAt time.Time `json:"executed_at"`
	Status     string    `json:"status"`
}
