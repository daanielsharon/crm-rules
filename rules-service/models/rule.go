package models

import "time"

type Rule struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Condition string     `json:"condition,omitempty"`
	Schedule  string     `json:"schedule,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
