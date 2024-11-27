package store

import (
	"context"
	"database/sql"

	"github.com/GiorgiMakharadze/social-network-golang/internal/models"
)

type CommentsStore struct {
	db *sql.DB
}

func (s *CommentsStore) GetByPostID(ctx context.Context, postID int64) ([]models.Comment, error) {
	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, users.username, users.id FROM comments c
		id FROM comments c
		JOIN users on users.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []models.Comment{}
	for rows.Next() {
		var c models.Comment
		c.User = models.User{}
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}
