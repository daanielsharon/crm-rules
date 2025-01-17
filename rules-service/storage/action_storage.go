package storage

import (
	"database/sql"
	"errors"
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

func (s *ActionStorage) GetActions() ([]models.Action, error) {
	query := "SELECT id, rule_id, action, created_at, updated_at FROM rule_actions"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actions []models.Action
	for rows.Next() {
		var action models.Action
		if err := rows.Scan(&action.ID, &action.RuleID, &action.Action, &action.CreatedAt, &action.UpdatedAt); err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}
	return actions, nil
}

func (s *ActionStorage) GetActionById(id string) (*models.Action, error) {
	query := "SELECT id, rule_id, action, created_at, updated_at FROM rule_actions WHERE id = $1"
	row := s.DB.QueryRow(query, id)

	var action models.Action
	err := row.Scan(&action.ID, &action.RuleID, &action.Action, &action.CreatedAt, &action.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("rule not found")
	} else if err != nil {
		return nil, err
	}

	return &action, nil
}

func (s *ActionStorage) UpdateAction(action models.Action) error {
	query := `
	UPDATE rule_actions
	SET action = $2, updated_at = NOW()
	WHERE id = $1
	`
	_, err := s.DB.Exec(query, action.ID, action.Action)
	return err
}

func (s *ActionStorage) DeleteAction(id string) error {
	query := "DELETE FROM rule_actions WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	return err
}
