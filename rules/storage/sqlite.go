package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"rules/models"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(databasePath string) RuleStorageInterface {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS rules (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		condition TEXT NOT NULL,
		action TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create rules table: %v", err)
	}

	return &Storage{DB: db}
}

func (s *Storage) CreateRule(rule models.Rule) error {
	query := `
	INSERT INTO rules (id, name, condition, action, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?);
	`
	_, err := s.DB.Exec(query, rule.ID, rule.Name, rule.Condition, rule.Action, rule.CreatedAt, rule.UpdatedAt)
	return err
}

func (s *Storage) GetAllRules() ([]models.Rule, error) {
	query := `SELECT id, name, condition, action, created_at, updated_at FROM rules;`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []models.Rule
	for rows.Next() {
		var rule models.Rule
		err := rows.Scan(&rule.ID, &rule.Name, &rule.Condition, &rule.Action, &rule.CreatedAt, &rule.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, nil
}

func (s *Storage) GetRule(id string) (*models.Rule, error) {
	query := `SELECT id, name, condition, action, created_at, updated_at FROM rules WHERE id = ?;`
	var rule models.Rule

	err := s.DB.QueryRow(query, id).Scan(
		&rule.ID,
		&rule.Name,
		&rule.Condition,
		&rule.Action,
		&rule.CreatedAt,
		&rule.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("rule with ID %s not found", id)
		}
		return nil, err
	}

	return &rule, nil
}

func (s *Storage) UpdateRule(rule models.Rule) error {
	query := `
	UPDATE rules 
	SET name = ?, condition = ?, action = ?, updated_at = ? 
	WHERE id = ?;`

	rule.UpdatedAt = time.Now()

	result, err := s.DB.Exec(
		query,
		rule.Name,
		rule.Condition,
		rule.Action,
		rule.UpdatedAt,
		rule.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rule with ID %s not found", rule.ID)
	}

	return nil
}

func (s *Storage) DeleteRule(id string) error {
	query := `DELETE FROM rules WHERE id = ?;`

	result, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rule with ID %s not found", id)
	}

	return nil
}
