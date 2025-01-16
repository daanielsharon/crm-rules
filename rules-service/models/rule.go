package models

import "time"

type Rule struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Condition string    `json:"condition"`
	Schedule  string    `json:"schedule"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
