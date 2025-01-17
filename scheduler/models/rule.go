package models

import "time"

type Rule struct {
	ID        int          `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	Condition string       `json:"condition" db:"condition"`
	Actions   []RuleAction `json:"actions" db:"actions"`
	Schedule  string       `json:"schedule" db:"schedule"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
}

type RuleAction struct {
	ID        int       `json:"id" db:"id"`
	Action    string    `json:"action" db:"action"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
