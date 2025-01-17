package models

import "time"

type Action struct {
	ID        string    `json:"id,omitempty"`
	RuleID    string    `json:"rule_id,omitempty"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
