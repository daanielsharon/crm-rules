package storage

import (
	"database/sql"
	"time"

	"worker-service/models"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{DB: db}
}

func (s *Storage) GetRules() ([]models.Rule, error) {
	query := `
        SELECT id, name, condition, action, schedule, created_at, updated_at
        FROM rules;
    `
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []models.Rule
	for rows.Next() {
		var rule models.Rule
		if err := rows.Scan(&rule.ID, &rule.Name, &rule.Condition, &rule.Action, &rule.Schedule, &rule.CreatedAt, &rule.UpdatedAt); err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, nil
}

func (s *Storage) LogExecution(ruleID string, status string) error {
	query := `
        INSERT INTO execution_logs (rule_id, executed_at, status)
        VALUES ($1, $2, $3)
    `
	_, err := s.DB.Exec(query, ruleID, time.Now(), status)
	return err
}
