package models

import (
	"time"
)

type User struct {
	ID            string     `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	Email         string     `json:"email,omitempty"`
	LastActive    *time.Time `json:"last_active,omitempty"`
	Plan          string     `json:"plan,omitempty"`
	FailedLogins  int        `json:"failed_logins,omitempty"`
	EmailVerified bool       `json:"email_verified,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
