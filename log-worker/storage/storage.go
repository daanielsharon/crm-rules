package storage

import (
	"database/sql"
	"log-worker/models"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) LogStorageInterface {
	return &Storage{DB: db}
}

func (s *Storage) CreateLog(log models.Log) error {
	query := `
		INSERT INTO execution_logs (rule_id, user_id, action, status)
		VALUES ($1, $2, $3, $4)
	`
	_, err := s.DB.Exec(query, log.RuleID, log.UserID, log.Action, log.Status)
	return err
}
