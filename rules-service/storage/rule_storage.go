package storage

import (
	"database/sql"
	"errors"
	"rules-service/models"
	"time"
)

type RuleStorage struct {
	DB *sql.DB
}

func NewRuleStorage(db *sql.DB) RuleStorageInterface {
	return &RuleStorage{DB: db}
}

func (s *RuleStorage) CreateRule(rule models.Rule) error {
	query := `
	INSERT INTO rules (name, condition, schedule)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at
	`
	return s.DB.QueryRow(query, rule.Name, rule.Condition, rule.Schedule).Scan(&rule.ID, &rule.CreatedAt, &rule.UpdatedAt)
}

func (s *RuleStorage) GetAllRules() ([]models.Rule, error) {
	query := `SELECT id, name, condition, schedule, created_at, updated_at FROM rules`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []models.Rule
	for rows.Next() {
		var rule models.Rule
		err := rows.Scan(&rule.ID, &rule.Name, &rule.Condition, &rule.Schedule, &rule.CreatedAt, &rule.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, nil
}

func (s *RuleStorage) GetRuleById(id string) (*models.Rule, error) {
	query := `SELECT id, name, condition, schedule, created_at, updated_at FROM rules WHERE id = $1`
	row := s.DB.QueryRow(query, id)

	var rule models.Rule
	err := row.Scan(&rule.ID, &rule.Name, &rule.Condition, &rule.Schedule, &rule.CreatedAt, &rule.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("rule not found")
	} else if err != nil {
		return nil, err
	}
	return &rule, nil
}

func (s *RuleStorage) UpdateRule(rule models.Rule) error {
	query := `
	UPDATE rules
	SET name = $2, condition = $3, schedule = $4, updated_at = $5
	WHERE id = $1
	`
	_, err := s.DB.Exec(query, rule.ID, rule.Name, rule.Condition, rule.Schedule, time.Now())
	return err
}

func (s *RuleStorage) DeleteRule(id string) error {
	query := `DELETE FROM rules WHERE id = $1`
	_, err := s.DB.Exec(query, id)
	return err
}
