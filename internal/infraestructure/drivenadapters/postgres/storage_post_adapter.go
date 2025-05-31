package postgres

import (
	"context"
	"database/sql"

	domain "github.com/felipeemora/social-hex-api/internal/domain/models"
	"github.com/lib/pq"
)

type StoragePostAdapter struct {
	db *sql.DB
}

func NewPostgresAdapter(db *sql.DB) *StoragePostAdapter {
	return &StoragePostAdapter{
		db,
	}
}

func (pa *StoragePostAdapter) Save(ctx context.Context, postModel *domain.PostModel) (*domain.PostModel, error) {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := pa.db.QueryRowContext(
		ctx, query, postModel.Content, postModel.Title, postModel.UserID, pq.Array(postModel.Tags),
	).Scan(&postModel.ID, &postModel.CreatedAt, &postModel.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return postModel, nil
}
