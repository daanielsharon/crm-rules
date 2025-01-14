package storage

import (
	"database/sql"
	"errors"
	"user-service/models"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) StorageInterface {
	return &Storage{DB: db}
}

func (s *Storage) CreateUser(user models.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2)"
	_, err := s.DB.Exec(query, user.Name, user.Email)
	return err
}

func (s *Storage) GetUserById(id string) (*models.User, error) {
	query := "SELECT id, name, email, created_at FROM users WHERE id = $1"
	row := s.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *Storage) GetAllUsers() ([]models.User, error) {
	query := "SELECT id, name, email, created_at FROM users"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *Storage) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, email, created_at FROM users WHERE email = $1"
	row := s.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *Storage) UpdateUser(user models.User) error {
	query := `
		UPDATE users
		SET name = COALESCE($2, name),
		    email = COALESCE($3, email)
		WHERE id = $1
		RETURNING id
	`
	var updatedID string
	err := s.DB.QueryRow(query, user.ID, user.Name, user.Email).Scan(&updatedID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}

func (s *Storage) DeleteUser(id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	return err
}
