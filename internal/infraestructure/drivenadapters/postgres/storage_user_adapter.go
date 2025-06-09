package postgres

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"time"

	domain "github.com/felipeemora/social-hex-api/internal/domain"
	domainModels "github.com/felipeemora/social-hex-api/internal/domain/models"
)

type StorageUserAdapter struct {
	db *sql.DB
}

func NewStorageUserAdapter(db *sql.DB) *StorageUserAdapter {
	return &StorageUserAdapter{db: db}
}

func (ua *StorageUserAdapter) CreateUserAndInvite(ctx context.Context, userModel *domainModels.UserModel, invitationId string, duration time.Duration) error {
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

func (ua *StorageUserAdapter) create(ctx context.Context, tx *sql.Tx, user *domainModels.UserModel) error {
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
			return domain.ErrDuplicateUsername
		case err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"":
			return domain.ErrDuplicateEmail
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

// ************************ Activate User Functions Start ************************
func (ua *StorageUserAdapter) Activate(ctx context.Context, invitationID *string) error {
	return withTransaction(ua.db, ctx, func(tx *sql.Tx) error {
		user, err := ua.getUserFromInvitation(ctx, tx, invitationID)
		if err != nil {
			return err
		}

		user.IsActive = true

		if err := ua.update(ctx, tx, user); err != nil {
			return err
		}

		if err := ua.deleteUserInvitations(ctx, tx, user.ID); err != nil {
			return err
		}

		return nil
	})
}

func (ua *StorageUserAdapter) getUserFromInvitation(ctx context.Context, tx *sql.Tx, invitationID *string) (*domainModels.UserModel, error) {
	query := `
		SELECT u.id, u.username, u.email, u.created_at, u.is_active
		FROM users u
		JOIN user_invitations ui ON u.id = ui.user_id
		WHERE ui.token = $1 AND ui.expiration > $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	hash := sha256.Sum256([]byte(*invitationID))
	hashInvitationID := hex.EncodeToString(hash[:])

	user := &domainModels.UserModel{}
	err := tx.QueryRowContext(ctx, query, hashInvitationID, time.Now()).Scan(
		&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.IsActive)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, domain.ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (ua *StorageUserAdapter) update(ctx context.Context, tx *sql.Tx, user *domainModels.UserModel) error {
	query := `UPDATE users SET username = $1, email = $2, is_active = $3 WHERE id = $4`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.IsActive, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// ************************ Activate User Functions End ************************

func (ua *StorageUserAdapter) GetByEmail(ctx context.Context, email *string) (*domainModels.UserModel, error) {
	query := `
		SELECT id, username, email, created_at, password
		FROM users
		WHERE email = $1 AND is_active = true
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var user domainModels.UserModel
	err := ua.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.PasswordHash,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, domain.ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
