package models

import "time"

type Action struct {
	ID        string    `json:"id,omitempty"`
	RuleID    string    `json:"rule_id"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
