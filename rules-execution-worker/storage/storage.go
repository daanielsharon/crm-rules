package storage

import (
	"database/sql"
)

type User struct {
	ID    string
	Name  string
	Email string
}

type Storage interface {
	GetMatchingUsers(condition string) (*sql.Rows, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) GetMatchingUsers(condition string) (*sql.Rows, error) {
	query := "SELECT id, name, email FROM users WHERE " + condition
	return s.db.Query(query)
}
