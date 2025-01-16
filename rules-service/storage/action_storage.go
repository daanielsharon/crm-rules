package storage

import (
	"database/sql"
	"rules-service/models"
)

type ActionStorage struct {
	DB *sql.DB
}

func NewActionStorage(db *sql.DB) ActionStorageInterface {
	return &ActionStorage{DB: db}
}

func (s *ActionStorage) CreateAction(action models.Action) error {
	query := `
	INSERT INTO rule_actions (rule_id, action)
	VALUES ($1, $2)
	RETURNING id, created_at, updated_at
	`
	return s.DB.QueryRow(query, action.RuleID, action.Action).Scan(&action.ID, &action.CreatedAt, &action.UpdatedAt)
}
