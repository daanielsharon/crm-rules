package models

type Log struct {
	ID         int    `json:"id"`
	RuleID     int    `json:"rule_id"`
	UserID     string `json:"user_id"`
	Action     string `json:"action"`
	Status     string `json:"status"`
	ExecutedAt string `json:"executed_at"`
}
