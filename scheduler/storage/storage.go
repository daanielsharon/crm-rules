package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
        SELECT 
            r.id, 
            r.name, 
            r.condition, 
            r.schedule,
            to_char(r.created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') as created_at,
            to_char(r.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') as updated_at,
            COALESCE(
                (
                    SELECT json_agg(
                        json_build_object(
                            'id', ra.id,
                            'action', ra.action,
                            'created_at', to_char(ra.created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),
                            'updated_at', to_char(ra.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"')
                        )
                    )
                    FROM rule_actions ra
                    WHERE ra.rule_id = r.id
                ),
                '[]'
            ) as actions
        FROM rules r;
    `
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []models.Rule
	for rows.Next() {
		var rule models.Rule
		var actionsJSON []byte
		var createdAtStr, updatedAtStr string
		if err := rows.Scan(
			&rule.ID,
			&rule.Name,
			&rule.Condition,
			&rule.Schedule,
			&createdAtStr,
			&updatedAtStr,
			&actionsJSON,
		); err != nil {
			return nil, err
		}

		createdAt, err := time.Parse(time.RFC3339, createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse created_at timestamp: %w", err)
		}

		updatedAt, err := time.Parse(time.RFC3339, updatedAtStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse updated_at timestamp: %w", err)
		}

		rule.CreatedAt = createdAt
		rule.UpdatedAt = updatedAt

		if err := json.Unmarshal(actionsJSON, &rule.Actions); err != nil {
			return nil, fmt.Errorf("failed to unmarshal actions: %w", err)
		}

		rules = append(rules, rule)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rules, nil
}
