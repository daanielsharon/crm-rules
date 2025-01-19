package storage

import (
	"database/sql"
	"errors"
	"time"
	"user-service/models"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) StorageInterface {
	return &Storage{DB: db}
}

func (s *Storage) CreateUser(user models.User) error {
	query := "INSERT INTO users (name, email, plan) VALUES ($1, $2, $3)"
	_, err := s.DB.Exec(query, user.Name, user.Email, user.Plan)
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
	query := "SELECT id, name, email, last_active, plan, failed_logins, email_verified, created_at, updated_at FROM users"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.LastActive, &user.Plan, &user.FailedLogins, &user.EmailVerified, &user.CreatedAt, &user.UpdatedAt); err != nil {
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
		    email = COALESCE($3, email),
			plan = COALESCE($4, plan),
			last_active = COALESCE($5, last_active),
			failed_logins = COALESCE($6, failed_logins),
			email_verified = COALESCE($7, email_verified),
			updated_at = COALESCE($8, updated_at)
		WHERE id = $1
		RETURNING id
	`
	var updatedID string
	err := s.DB.QueryRow(query, user.ID, user.Name, user.Email, user.Plan, user.LastActive, user.FailedLogins, user.EmailVerified, time.Now()).Scan(&updatedID)
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
