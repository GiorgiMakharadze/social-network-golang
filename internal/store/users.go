package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/GiorgiMakharadze/social-network-golang/internal/models"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id,
		created at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersStore) GetByID(ctx context.Context, userID int64) (*models.User, error) {
	query := `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE id =$1
	`

	user := &models.User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}
