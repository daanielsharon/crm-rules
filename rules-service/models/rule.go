package models

import "time"

type Rule struct {
	ID        string    `json:"id"`         // Unique identifier
	Name      string    `json:"name"`       // Rule name
	Condition string    `json:"condition"`  // Condition (e.g., "last_active > 90")
	Action    string    `json:"action"`     // Action to perform (e.g., "send_notification")
	CreatedAt time.Time `json:"created_at"` // Timestamp of creation
	UpdatedAt time.Time `json:"updated_at"` // Timestamp of last update
}
