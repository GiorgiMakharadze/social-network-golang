package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/GiorgiMakharadze/social-network-golang/internal/models"
)

var (
	ErrNotFound          = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		GetByID(context.Context, int64) (*models.Post, error)
		Create(context.Context, *models.Post) error
		Delete(context.Context, int64) error
		Update(context.Context, *models.Post) error
	}
	Users interface {
		Create(context.Context, *models.User) error
		GetByID(context.Context, int64) (*models.User, error)
	}
	Comments interface {
		GetByPostID(context.Context, int64) ([]models.Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UsersStore{db},
		Comments: &CommentsStore{db},
	}
}
