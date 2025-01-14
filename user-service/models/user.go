package models

import (
	"time"
)

type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	LastActive    time.Time `json:"last_active"`
	Plan          string    `json:"plan"`
	FailedLogins  int       `json:"failed_logins"`
	EmailVerified bool      `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
