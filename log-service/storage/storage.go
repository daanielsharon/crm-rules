package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log-service/models"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) StorageInterface {
	return &Storage{DB: db}
}

func (s *Storage) GetLogs(ruleID, userID string) ([]models.Log, error) {
	query := "SELECT id, rule_id, user_id, action, status, executed_at FROM execution_logs WHERE 1=1"
	var args []interface{}
	argCount := 1

	if ruleID != "" {
		query += fmt.Sprintf(" AND rule_id = $%d", argCount)
		args = append(args, ruleID)
		argCount++
	}

	if userID != "" {
		query += fmt.Sprintf(" AND user_id = $%d", argCount)
		args = append(args, userID)
		argCount++
	}

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.Log
	for rows.Next() {
		var log models.Log
		if err := rows.Scan(&log.ID, &log.RuleID, &log.UserID, &log.Action, &log.Status, &log.ExecutedAt); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func (s *Storage) GetLogById(id string) (*models.Log, error) {
	query := "SELECT id, rule_id, user_id, action, status, executed_at FROM execution_logs WHERE id = $1"
	row := s.DB.QueryRow(query, id)

	var log models.Log
	err := row.Scan(&log.ID, &log.RuleID, &log.UserID, &log.Action, &log.Status, &log.ExecutedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &log, nil
}
