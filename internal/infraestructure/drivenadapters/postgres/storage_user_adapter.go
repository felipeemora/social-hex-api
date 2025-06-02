package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	domain "github.com/felipeemora/social-hex-api/internal/domain/models"
)

var (
	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicateEmail    = errors.New("email already exists")
)

type StorageUserAdapter struct {
	db *sql.DB
}

func NewStorageUserAdapter(db *sql.DB) *StorageUserAdapter {
	return &StorageUserAdapter{db: db}
}

func (ua *StorageUserAdapter) CreateUserAndInvite(ctx context.Context, userModel *domain.UserModel, invitationId string, duration time.Duration) error {
	return withTransaction(ua.db, ctx, func(tx *sql.Tx) error {
		if err := ua.create(ctx, tx, userModel); err != nil {
			return err
		}

		if err := ua.createUserInvitation(ctx, tx, userModel.ID, invitationId, duration); err != nil {
			return err
		}

		return nil
	})
}

func (ua *StorageUserAdapter) create(ctx context.Context, tx *sql.Tx, user *domain.UserModel) error {
	query := `
		INSERT INTO users (username, email, password, role_id)
		VALUES ($1, $2, $3, (SELECT id FROM roles WHERE name = $4))
		RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := tx.QueryRowContext(
		ctx, query, user.Username, user.Email, user.PasswordHash, user.Role.Name,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		switch {
		case err.Error() == "pq: duplicate key value violates unique constraint \"idx_users_username\"":
			return ErrDuplicateUsername
		case err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"":
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (ua *StorageUserAdapter) createUserInvitation(ctx context.Context, tx *sql.Tx, userId int64, invitationID string, expiration time.Duration) error {
	query := `
		INSERT INTO user_invitations (user_id, token, expiration)
		VALUES ($1, $2, $3)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, userId, invitationID, time.Now().Add(expiration))
	if err != nil {
		return err
	}

	return nil
}


func (ua *StorageUserAdapter) Delete(ctx context.Context, id int64) error {
	return withTransaction(ua.db, ctx, func(tx *sql.Tx) error {
		if err := ua.delete(ctx, tx, id); err != nil {
			return err
		}

		if err := ua.deleteUserInvitations(ctx, tx, id); err != nil {
			return err
		}

		return nil
	})
}

func (ua *StorageUserAdapter) delete(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (ua *StorageUserAdapter) deleteUserInvitations(ctx context.Context, tx *sql.Tx, userId int64) error {
	query := `DELETE FROM user_invitations WHERE user_id = $1`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}
