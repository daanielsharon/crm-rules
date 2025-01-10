package storage

import (
	"database/sql"
)

type User struct {
	ID    string
	Name  string
	Email string
}

type Store interface {
	GetMatchingUsers(condition string) (*sql.Rows, error)
	LogExecution(ruleID int, userID string, action string, status string) error
}

type PostgresStore struct {
	db *sql.DB
}

func New(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) GetMatchingUsers(condition string) (*sql.Rows, error) {
	query := "SELECT id, name, email FROM users WHERE " + condition
	return s.db.Query(query)
}

func (s *PostgresStore) LogExecution(ruleID int, userID string, action string, status string) error {
	query := `
		INSERT INTO rule_executions (rule_id, user_id, action, status, executed_at)
		VALUES ($1, $2, $3, $4, NOW())
	`
	_, err := s.db.Exec(query, ruleID, userID, action, status)
	return err
}
