package models

import "time"

type Action struct {
	ID        int       `json:"id"`
	RuleID    string    `json:"rule_id"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
